package bq

import (
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func NewBQClient(c context.Context, projectId string) *bigquery.Client {
	client, _ := bigquery.NewClient(c, projectId)
	return client
}

func GetQueryResults(c context.Context, BQClient *bigquery.Client, queryString string) ([][]bigquery.Value, []byte) {
	var q = BQClient.Query(queryString)
	// q.DryRun = true
	// res, _ := q.Run(c)
	// fmt.Println(res.LastStatus().Statistics.TotalBytesProcessed)
	it, _ := q.Read(c)
	var results [][]bigquery.Value
	var resultsSchema []byte
	for {
		var values []bigquery.Value
		// var values map[string]bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			var sc, _ = it.Schema.ToJSONFields()
			resultsSchema = sc
			break
		}
		results = append(results, values)
		if err != nil {
			break
		}
	}
	return results, resultsSchema
}
