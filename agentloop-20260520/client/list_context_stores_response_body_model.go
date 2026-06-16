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
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*ListContextStoresResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// 5
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
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// example:
	//
	// experience
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 我的上下文库
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// ["order-service","payment-service"]
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListContextStoresResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoresResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListContextStoresResponseBodyResults) GetAgentSpace() *string {
	return s.AgentSpace
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

func (s *ListContextStoresResponseBodyResults) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *ListContextStoresResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *ListContextStoresResponseBodyResults) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListContextStoresResponseBodyResults) SetAgentSpace(v string) *ListContextStoresResponseBodyResults {
	s.AgentSpace = &v
	return s
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

func (s *ListContextStoresResponseBodyResults) SetServiceNames(v []*string) *ListContextStoresResponseBodyResults {
	s.ServiceNames = v
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

func (s *ListContextStoresResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
