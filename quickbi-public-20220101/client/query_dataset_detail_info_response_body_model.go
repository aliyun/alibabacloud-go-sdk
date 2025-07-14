// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetDetailInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDatasetDetailInfoResponseBody
	GetRequestId() *string
	SetResult(v string) *QueryDatasetDetailInfoResponseBody
	GetResult() *string
	SetSuccess(v bool) *QueryDatasetDetailInfoResponseBody
	GetSuccess() *bool
}

type QueryDatasetDetailInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the dataset data in JSON format: `{ "cube": { "dimensions": [ { "caption": "customer name", "dataType": "string", "dimensionType": "standard_dimension", "factColumn": "customer_name", "uid": "N5820f5_customer_name" }, { "caption": "datastring", "" standard_dimension", "factColumn": "order_id", "uid": "N5820f5_order_id" }, ], "measures": [ { "caption": "order amount ", "dataType": "number", "factColumn": "order_amt", "measureType": "standard_measure ": " Nderamid " }, " { "customsql": false, "dsId": "261b252d-c3c3-498a-a0a7-5d1ec6cd****", "tableName": "company_sales_record_copy" } }, "datasetId": "5820f58c-c734-4d8a-baf1-7979af4f****", "datasetName": "company_sales_record_copy12", "datasource": { "dsId": "261b252d-c3c3-498a-a0a7-5d1ec6cd****", "dsName": "Self-use", "dsType": "mysql" }, "directory" { "id": "schemaad8aad00-9c55-4984-a767-b4e0ec60****", "name": "My dataset", "pathId": "schemaad8aad00-9c55-4984-a767-b4e0ec60****", "pathName": "My dataset" }, "ownerId": "13651626232****", "ownerName": "Zhang San", "rowLevel": false, "workspaceId": "95296e95-ca89-4c7d-8af9-dedf0ad0****", "workspaceName": "Test Workspace" }`
	//
	// example:
	//
	// A JSON dataset is returned. For more information, see the description on the left.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetDetailInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetDetailInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDatasetDetailInfoResponseBody) GetResult() *string {
	return s.Result
}

func (s *QueryDatasetDetailInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDatasetDetailInfoResponseBody) SetRequestId(v string) *QueryDatasetDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetDetailInfoResponseBody) SetResult(v string) *QueryDatasetDetailInfoResponseBody {
	s.Result = &v
	return s
}

func (s *QueryDatasetDetailInfoResponseBody) SetSuccess(v bool) *QueryDatasetDetailInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDatasetDetailInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
