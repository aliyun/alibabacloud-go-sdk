// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCustomAgentsResponseBodyPagingInfo) *ListCustomAgentsResponseBody
	GetPagingInfo() *ListCustomAgentsResponseBodyPagingInfo
	SetRequestId(v string) *ListCustomAgentsResponseBody
	GetRequestId() *string
}

type ListCustomAgentsResponseBody struct {
	PagingInfo *ListCustomAgentsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCustomAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBody) GetPagingInfo() *ListCustomAgentsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCustomAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAgentsResponseBody) SetPagingInfo(v *ListCustomAgentsResponseBodyPagingInfo) *ListCustomAgentsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCustomAgentsResponseBody) SetRequestId(v string) *ListCustomAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAgentsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomAgentsResponseBodyPagingInfo struct {
	Agents []*ListCustomAgentsResponseBodyPagingInfoAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 5
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomAgentsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBodyPagingInfo) GetAgents() []*ListCustomAgentsResponseBodyPagingInfoAgents {
	return s.Agents
}

func (s *ListCustomAgentsResponseBodyPagingInfo) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomAgentsResponseBodyPagingInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomAgentsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomAgentsResponseBodyPagingInfo) SetAgents(v []*ListCustomAgentsResponseBodyPagingInfoAgents) *ListCustomAgentsResponseBodyPagingInfo {
	s.Agents = v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfo) SetMaxResults(v int32) *ListCustomAgentsResponseBodyPagingInfo {
	s.MaxResults = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfo) SetNextToken(v string) *ListCustomAgentsResponseBodyPagingInfo {
	s.NextToken = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfo) SetTotalCount(v int32) *ListCustomAgentsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfo) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomAgentsResponseBodyPagingInfoAgents struct {
	// example:
	//
	// 123456
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListCustomAgentsResponseBodyPagingInfoAgents) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBodyPagingInfoAgents) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetDescription() *string {
	return s.Description
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetName() *string {
	return s.Name
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) GetVisibility() *string {
	return s.Visibility
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetCreatorId(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.CreatorId = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetDescription(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.Description = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetDisplayName(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.DisplayName = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetGmtCreateTime(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.GmtCreateTime = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetGmtModifiedTime(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetModifierId(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.ModifierId = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetName(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.Name = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) SetVisibility(v string) *ListCustomAgentsResponseBodyPagingInfoAgents {
	s.Visibility = &v
	return s
}

func (s *ListCustomAgentsResponseBodyPagingInfoAgents) Validate() error {
	return dara.Validate(s)
}
