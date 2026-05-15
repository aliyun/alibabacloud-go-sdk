// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentResponseBody
	GetCode() *string
	SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody
	GetData() *GetAgentResponseBodyData
	SetHttpStatusCode(v int64) *GetAgentResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentResponseBody
	GetSuccess() *bool
}

type GetAgentResponseBody struct {
	// example:
	//
	// Success
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentResponseBody) GetData() *GetAgentResponseBodyData {
	return s.Data
}

func (s *GetAgentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetHttpStatusCode(v int64) *GetAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetSuccess(v bool) *GetAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentResponseBodyData struct {
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 222222
	AgentId     *int64                               `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	DisplayName *string                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GroupList   []*GetAgentResponseBodyDataGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAgentResponseBodyData) GetAgentId() *int64 {
	return s.AgentId
}

func (s *GetAgentResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAgentResponseBodyData) GetGroupList() []*GetAgentResponseBodyDataGroupList {
	return s.GroupList
}

func (s *GetAgentResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetAgentResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetAgentResponseBodyData) SetAccountName(v string) *GetAgentResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetAgentId(v int64) *GetAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentResponseBodyData) SetDisplayName(v string) *GetAgentResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetGroupList(v []*GetAgentResponseBodyDataGroupList) *GetAgentResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetAgentResponseBodyData) SetStatus(v int32) *GetAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentResponseBodyData) SetTenantId(v int64) *GetAgentResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetAgentResponseBodyData) Validate() error {
	if s.GroupList != nil {
		for _, item := range s.GroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentResponseBodyDataGroupList struct {
	// example:
	//
	// 1
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s GetAgentResponseBodyDataGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyDataGroupList) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyDataGroupList) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *GetAgentResponseBodyDataGroupList) GetDescription() *string {
	return s.Description
}

func (s *GetAgentResponseBodyDataGroupList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAgentResponseBodyDataGroupList) GetName() *string {
	return s.Name
}

func (s *GetAgentResponseBodyDataGroupList) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *GetAgentResponseBodyDataGroupList) SetChannelType(v int32) *GetAgentResponseBodyDataGroupList {
	s.ChannelType = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetDescription(v string) *GetAgentResponseBodyDataGroupList {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetDisplayName(v string) *GetAgentResponseBodyDataGroupList {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetName(v string) *GetAgentResponseBodyDataGroupList {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetSkillGroupId(v int64) *GetAgentResponseBodyDataGroupList {
	s.SkillGroupId = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) Validate() error {
	return dara.Validate(s)
}
