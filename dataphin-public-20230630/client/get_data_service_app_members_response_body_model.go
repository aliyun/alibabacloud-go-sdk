// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceAppMembersResponseBody
	GetCode() *string
	SetData(v *GetDataServiceAppMembersResponseBodyData) *GetDataServiceAppMembersResponseBody
	GetData() *GetDataServiceAppMembersResponseBodyData
	SetHttpStatusCode(v int32) *GetDataServiceAppMembersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAppMembersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceAppMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAppMembersResponseBody
	GetSuccess() *bool
}

type GetDataServiceAppMembersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataServiceAppMembersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceAppMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAppMembersResponseBody) GetData() *GetDataServiceAppMembersResponseBodyData {
	return s.Data
}

func (s *GetDataServiceAppMembersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAppMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAppMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAppMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAppMembersResponseBody) SetCode(v string) *GetDataServiceAppMembersResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBody) SetData(v *GetDataServiceAppMembersResponseBodyData) *GetDataServiceAppMembersResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceAppMembersResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAppMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBody) SetMessage(v string) *GetDataServiceAppMembersResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBody) SetRequestId(v string) *GetDataServiceAppMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBody) SetSuccess(v bool) *GetDataServiceAppMembersResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataServiceAppMembersResponseBodyData struct {
	MemberList []*GetDataServiceAppMembersResponseBodyDataMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s GetDataServiceAppMembersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppMembersResponseBodyData) GetMemberList() []*GetDataServiceAppMembersResponseBodyDataMemberList {
	return s.MemberList
}

func (s *GetDataServiceAppMembersResponseBodyData) SetMemberList(v []*GetDataServiceAppMembersResponseBodyDataMemberList) *GetDataServiceAppMembersResponseBodyData {
	s.MemberList = v
	return s
}

func (s *GetDataServiceAppMembersResponseBodyData) Validate() error {
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataServiceAppMembersResponseBodyDataMemberList struct {
	// example:
	//
	// 2026-12-12
	EffectiveEnd *string `json:"EffectiveEnd,omitempty" xml:"EffectiveEnd,omitempty"`
	// example:
	//
	// general
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 200000245
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDataServiceAppMembersResponseBodyDataMemberList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppMembersResponseBodyDataMemberList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) GetRole() *string {
	return s.Role
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) GetUserId() *string {
	return s.UserId
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) SetEffectiveEnd(v string) *GetDataServiceAppMembersResponseBodyDataMemberList {
	s.EffectiveEnd = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) SetRole(v string) *GetDataServiceAppMembersResponseBodyDataMemberList {
	s.Role = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) SetUserId(v string) *GetDataServiceAppMembersResponseBodyDataMemberList {
	s.UserId = &v
	return s
}

func (s *GetDataServiceAppMembersResponseBodyDataMemberList) Validate() error {
	return dara.Validate(s)
}
