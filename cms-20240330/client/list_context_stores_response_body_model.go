// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextStoresResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextStoresResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListContextStoresResponseBody
	GetRequestId() *string
	SetResults(v []*ListContextStoresResponseBodyResults) *ListContextStoresResponseBody
	GetResults() []*ListContextStoresResponseBodyResults
	SetTotal(v int32) *ListContextStoresResponseBody
	GetTotal() *int32
}

type ListContextStoresResponseBody struct {
	// The maximum number of results returned per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token to retrieve the next page of results. If this field is empty, all results have been returned.
	//
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// A list of context stores.
	Results []*ListContextStoresResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// The total number of context stores.
	//
	// example:
	//
	// 56
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListContextStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextStoresResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextStoresResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextStoresResponseBody) GetResults() []*ListContextStoresResponseBodyResults {
	return s.Results
}

func (s *ListContextStoresResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextStoresResponseBody) SetMaxResults(v int32) *ListContextStoresResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContextStoresResponseBody) SetNextToken(v string) *ListContextStoresResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContextStoresResponseBody) SetRequestId(v string) *ListContextStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextStoresResponseBody) SetResults(v []*ListContextStoresResponseBodyResults) *ListContextStoresResponseBody {
	s.Results = v
	return s
}

func (s *ListContextStoresResponseBody) SetTotal(v int32) *ListContextStoresResponseBody {
	s.Total = &v
	return s
}

func (s *ListContextStoresResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextStoresResponseBodyResults struct {
	// The name of the context store.
	//
	// example:
	//
	// test-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// The context type.
	//
	// example:
	//
	// memory
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// The creation time of the context store. The value is a Unix timestamp in milliseconds.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 173847364841938
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the context store.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The status of the context store.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last update time of the context store. The value is a Unix timestamp in milliseconds.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 173847364841938
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContextStoresResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoresResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListContextStoresResponseBodyResults) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *ListContextStoresResponseBodyResults) GetContextType() *string {
	return s.ContextType
}

func (s *ListContextStoresResponseBodyResults) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListContextStoresResponseBodyResults) GetDescription() *string {
	return s.Description
}

func (s *ListContextStoresResponseBodyResults) GetRegionId() *string {
	return s.RegionId
}

func (s *ListContextStoresResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *ListContextStoresResponseBodyResults) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListContextStoresResponseBodyResults) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContextStoresResponseBodyResults) SetContextStoreName(v string) *ListContextStoresResponseBodyResults {
	s.ContextStoreName = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetContextType(v string) *ListContextStoresResponseBodyResults {
	s.ContextType = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetCreateTime(v string) *ListContextStoresResponseBodyResults {
	s.CreateTime = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetDescription(v string) *ListContextStoresResponseBodyResults {
	s.Description = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetRegionId(v string) *ListContextStoresResponseBodyResults {
	s.RegionId = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetStatus(v string) *ListContextStoresResponseBodyResults {
	s.Status = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetUpdateTime(v string) *ListContextStoresResponseBodyResults {
	s.UpdateTime = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) SetWorkspace(v string) *ListContextStoresResponseBodyResults {
	s.Workspace = &v
	return s
}

func (s *ListContextStoresResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
