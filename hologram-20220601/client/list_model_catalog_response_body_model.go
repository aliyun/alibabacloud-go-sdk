// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListModelCatalogResponseBody
	GetMaxResults() *int32
	SetModelList(v []*ListModelCatalogResponseBodyModelList) *ListModelCatalogResponseBody
	GetModelList() []*ListModelCatalogResponseBodyModelList
	SetNextToken(v string) *ListModelCatalogResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelCatalogResponseBody
	GetRequestId() *string
}

type ListModelCatalogResponseBody struct {
	// maxResults
	//
	// example:
	//
	// 50
	MaxResults *int32                                   `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	ModelList  []*ListModelCatalogResponseBodyModelList `json:"modelList,omitempty" xml:"modelList,omitempty" type:"Repeated"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelCatalogResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelCatalogResponseBody) GetModelList() []*ListModelCatalogResponseBodyModelList {
	return s.ModelList
}

func (s *ListModelCatalogResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelCatalogResponseBody) SetMaxResults(v int32) *ListModelCatalogResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelCatalogResponseBody) SetModelList(v []*ListModelCatalogResponseBodyModelList) *ListModelCatalogResponseBody {
	s.ModelList = v
	return s
}

func (s *ListModelCatalogResponseBody) SetNextToken(v string) *ListModelCatalogResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelCatalogResponseBody) SetRequestId(v string) *ListModelCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelCatalogResponseBody) Validate() error {
	if s.ModelList != nil {
		for _, item := range s.ModelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelCatalogResponseBodyModelList struct {
	// example:
	//
	// {"timeout": 600, "max_retries": 10, "max_retry_delay": 8, "initial_retry_delay": 0.5}
	DefaultParams *string `json:"defaultParams,omitempty" xml:"defaultParams,omitempty"`
	// example:
	//
	// {\\"deepThink\\":false,\\"onlineSearch\\":true}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// qwen3.5-plus
	ModelType     *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	ParamsExample *string `json:"paramsExample,omitempty" xml:"paramsExample,omitempty"`
	// example:
	//
	// bailian
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// cn-beijing
	ServiceDeployRegion *string `json:"serviceDeployRegion,omitempty" xml:"serviceDeployRegion,omitempty"`
	// example:
	//
	// chat/completions
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s ListModelCatalogResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s ListModelCatalogResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *ListModelCatalogResponseBodyModelList) GetDefaultParams() *string {
	return s.DefaultParams
}

func (s *ListModelCatalogResponseBodyModelList) GetExtra() *string {
	return s.Extra
}

func (s *ListModelCatalogResponseBodyModelList) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelCatalogResponseBodyModelList) GetParamsExample() *string {
	return s.ParamsExample
}

func (s *ListModelCatalogResponseBodyModelList) GetProvider() *string {
	return s.Provider
}

func (s *ListModelCatalogResponseBodyModelList) GetServiceDeployRegion() *string {
	return s.ServiceDeployRegion
}

func (s *ListModelCatalogResponseBodyModelList) GetTaskType() *string {
	return s.TaskType
}

func (s *ListModelCatalogResponseBodyModelList) SetDefaultParams(v string) *ListModelCatalogResponseBodyModelList {
	s.DefaultParams = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) SetExtra(v string) *ListModelCatalogResponseBodyModelList {
	s.Extra = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) SetModelType(v string) *ListModelCatalogResponseBodyModelList {
	s.ModelType = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) SetParamsExample(v string) *ListModelCatalogResponseBodyModelList {
	s.ParamsExample = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) SetProvider(v string) *ListModelCatalogResponseBodyModelList {
	s.Provider = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) SetServiceDeployRegion(v string) *ListModelCatalogResponseBodyModelList {
	s.ServiceDeployRegion = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) SetTaskType(v string) *ListModelCatalogResponseBodyModelList {
	s.TaskType = &v
	return s
}

func (s *ListModelCatalogResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
