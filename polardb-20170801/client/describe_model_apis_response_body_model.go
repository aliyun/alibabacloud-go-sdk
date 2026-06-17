// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeModelApisResponseBodyItems) *DescribeModelApisResponseBody
	GetItems() []*DescribeModelApisResponseBodyItems
	SetPageNumber(v int32) *DescribeModelApisResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeModelApisResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeModelApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeModelApisResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeModelApisResponseBody
	GetTotalRecordCount() *int32
}

type DescribeModelApisResponseBody struct {
	// A list of Model API objects.
	Items []*DescribeModelApisResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records returned on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of records to return on each page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the query.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeModelApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelApisResponseBody) GetItems() []*DescribeModelApisResponseBodyItems {
	return s.Items
}

func (s *DescribeModelApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModelApisResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeModelApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelApisResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeModelApisResponseBody) SetItems(v []*DescribeModelApisResponseBodyItems) *DescribeModelApisResponseBody {
	s.Items = v
	return s
}

func (s *DescribeModelApisResponseBody) SetPageNumber(v int32) *DescribeModelApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelApisResponseBody) SetPageRecordCount(v int32) *DescribeModelApisResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeModelApisResponseBody) SetPageSize(v int32) *DescribeModelApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeModelApisResponseBody) SetRequestId(v string) *DescribeModelApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelApisResponseBody) SetTotalRecordCount(v int32) *DescribeModelApisResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeModelApisResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelApisResponseBodyItems struct {
	// The model category. Valid values:
	//
	// - **text**
	//
	// - **embedding**
	//
	// - **rerank**
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The Model API ID.
	//
	// example:
	//
	// mi-xxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// The Model API name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The API path prefix.
	//
	// example:
	//
	// tests/models/
	PathPrefix *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	// The protocol. Valid values:
	//
	// - **openai**
	//
	// - **anthropic**
	//
	// - **bailian**
	//
	// - **vllm**
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of input records.
	//
	// example:
	//
	// 10
	RecordInput *string `json:"RecordInput,omitempty" xml:"RecordInput,omitempty"`
	// The number of output records.
	//
	// example:
	//
	// 10
	RecordOutput *string `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty"`
	// A JSON array of routing rules, formatted as a string.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "RuleName": "string",
	//
	//     "FallbackMode": "failover",
	//
	//     "MatchModelListJson": "[]",
	//
	//     "providerBalancerAlgorithm": "round-robin",
	//
	//     "Providers": [
	//
	//       {
	//
	//         "ModelServiceName": "string",
	//
	//         "Weight": "0",
	//
	//         "model_protocol": "vllm"
	//
	//         "ModelList": "[]"
	//
	//       }
	//
	//     ],
	//
	//     "FallbackProviders": [
	//
	//       {
	//
	//         "ModelServiceName": "string",
	//
	//         "model_protocol": "anthropic",
	//
	//         "Weight": "10",
	//
	//         "ModelList": "[]"
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	// ]
	RouteRules *string `json:"RouteRules,omitempty" xml:"RouteRules,omitempty"`
	// The Model API status.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModelApisResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelApisResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeModelApisResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeModelApisResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeModelApisResponseBodyItems) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *DescribeModelApisResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeModelApisResponseBodyItems) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *DescribeModelApisResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeModelApisResponseBodyItems) GetRecordInput() *string {
	return s.RecordInput
}

func (s *DescribeModelApisResponseBodyItems) GetRecordOutput() *string {
	return s.RecordOutput
}

func (s *DescribeModelApisResponseBodyItems) GetRouteRules() *string {
	return s.RouteRules
}

func (s *DescribeModelApisResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelApisResponseBodyItems) SetCategory(v string) *DescribeModelApisResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetGmtCreated(v string) *DescribeModelApisResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetModelApiId(v string) *DescribeModelApisResponseBodyItems {
	s.ModelApiId = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetName(v string) *DescribeModelApisResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetPathPrefix(v string) *DescribeModelApisResponseBodyItems {
	s.PathPrefix = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetProtocol(v string) *DescribeModelApisResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetRecordInput(v string) *DescribeModelApisResponseBodyItems {
	s.RecordInput = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetRecordOutput(v string) *DescribeModelApisResponseBodyItems {
	s.RecordOutput = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetRouteRules(v string) *DescribeModelApisResponseBodyItems {
	s.RouteRules = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetStatus(v string) *DescribeModelApisResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
