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
	SetTotalPages(v int32) *DescribeModelApisResponseBody
	GetTotalPages() *int32
	SetTotalRecordCount(v int32) *DescribeModelApisResponseBody
	GetTotalRecordCount() *int32
}

type DescribeModelApisResponseBody struct {
	// The list of model APIs.
	Items []*DescribeModelApisResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 3
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The total number of records.
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

func (s *DescribeModelApisResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
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

func (s *DescribeModelApisResponseBody) SetTotalPages(v int32) *DescribeModelApisResponseBody {
	s.TotalPages = &v
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
	// The category. Valid values:
	//
	// 	- **text**
	//
	// 	- **embedding**
	//
	// 	- **rerank**
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The gateway retry configuration.
	//
	// example:
	//
	// {"failover":{"enabled":true,"max_provider_retries":2,"max_failover_providers":1,"retryable_status_codes":[429,500,502,503,504],"retry_delay":0.5}}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the gateway instance to which the model API belongs.
	//
	// example:
	//
	// pg-2ze3x6i1d7g5rto28
	CustinsName *string `json:"CustinsName,omitempty" xml:"CustinsName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time of the model API.
	//
	// example:
	//
	// 2026-11-09T06:43:33Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The model API IDs.
	//
	// example:
	//
	// mi-xxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// The model name mappings, which is a JSON array string. Each element is in the format of {"Source":"client model name","Target":"upstream model name"}. If not configured, this parameter is empty.
	//
	// example:
	//
	// [{"Source":"client-model-a","Target":"upstream-model-a"}]
	ModelMappings *string `json:"ModelMappings,omitempty" xml:"ModelMappings,omitempty"`
	// The intelligent routing candidate list, which is a JSON object string in the format of {"AllowedModels":["candidate model name"]}. If not configured, this parameter is empty.
	//
	// example:
	//
	// {"AllowedModels":["deepseek-v4-pro","glm-5.3","glm-5.3-flash"]}
	ModelRouting *string `json:"ModelRouting,omitempty" xml:"ModelRouting,omitempty"`
	// The model API name.
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
	// 	- **openai**
	//
	// 	- **anthropic**
	//
	// 	- **bailian**
	//
	// 	- **vllm**
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of input points.
	//
	// example:
	//
	// 10
	RecordInput *string `json:"RecordInput,omitempty" xml:"RecordInput,omitempty"`
	// The number of output points.
	//
	// example:
	//
	// 10
	RecordOutput *string `json:"RecordOutput,omitempty" xml:"RecordOutput,omitempty"`
	// The list of routing rules, which is a JSON array string.
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
	// The model API status.
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

func (s *DescribeModelApisResponseBodyItems) GetConfig() *string {
	return s.Config
}

func (s *DescribeModelApisResponseBodyItems) GetCustinsName() *string {
	return s.CustinsName
}

func (s *DescribeModelApisResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeModelApisResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeModelApisResponseBodyItems) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *DescribeModelApisResponseBodyItems) GetModelMappings() *string {
	return s.ModelMappings
}

func (s *DescribeModelApisResponseBodyItems) GetModelRouting() *string {
	return s.ModelRouting
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

func (s *DescribeModelApisResponseBodyItems) SetConfig(v string) *DescribeModelApisResponseBodyItems {
	s.Config = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetCustinsName(v string) *DescribeModelApisResponseBodyItems {
	s.CustinsName = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetGmtCreated(v string) *DescribeModelApisResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetGmtModified(v string) *DescribeModelApisResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetModelApiId(v string) *DescribeModelApisResponseBodyItems {
	s.ModelApiId = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetModelMappings(v string) *DescribeModelApisResponseBodyItems {
	s.ModelMappings = &v
	return s
}

func (s *DescribeModelApisResponseBodyItems) SetModelRouting(v string) *DescribeModelApisResponseBodyItems {
	s.ModelRouting = &v
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
