// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivacyRuleSharedMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrivacyRuleSharedMemberResponseBody
	GetCode() *string
	SetData(v []*ListPrivacyRuleSharedMemberResponseBodyData) *ListPrivacyRuleSharedMemberResponseBody
	GetData() []*ListPrivacyRuleSharedMemberResponseBodyData
	SetHttpStatusCode(v int32) *ListPrivacyRuleSharedMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPrivacyRuleSharedMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrivacyRuleSharedMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPrivacyRuleSharedMemberResponseBody
	GetSuccess() *bool
}

type ListPrivacyRuleSharedMemberResponseBody struct {
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListPrivacyRuleSharedMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPrivacyRuleSharedMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleSharedMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleSharedMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrivacyRuleSharedMemberResponseBody) GetData() []*ListPrivacyRuleSharedMemberResponseBodyData {
	return s.Data
}

func (s *ListPrivacyRuleSharedMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPrivacyRuleSharedMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrivacyRuleSharedMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivacyRuleSharedMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPrivacyRuleSharedMemberResponseBody) SetCode(v string) *ListPrivacyRuleSharedMemberResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBody) SetData(v []*ListPrivacyRuleSharedMemberResponseBodyData) *ListPrivacyRuleSharedMemberResponseBody {
	s.Data = v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBody) SetHttpStatusCode(v int32) *ListPrivacyRuleSharedMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBody) SetMessage(v string) *ListPrivacyRuleSharedMemberResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBody) SetRequestId(v string) *ListPrivacyRuleSharedMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBody) SetSuccess(v bool) *ListPrivacyRuleSharedMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivacyRuleSharedMemberResponseBodyData struct {
	BizChainId   *string                                                  `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	BizChainName *string                                                  `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	MemberList   []*ListPrivacyRuleSharedMemberResponseBodyDataMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s ListPrivacyRuleSharedMemberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleSharedMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) GetBizChainName() *string {
	return s.BizChainName
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) GetMemberList() []*ListPrivacyRuleSharedMemberResponseBodyDataMemberList {
	return s.MemberList
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) SetBizChainId(v string) *ListPrivacyRuleSharedMemberResponseBodyData {
	s.BizChainId = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) SetBizChainName(v string) *ListPrivacyRuleSharedMemberResponseBodyData {
	s.BizChainName = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) SetMemberList(v []*ListPrivacyRuleSharedMemberResponseBodyDataMemberList) *ListPrivacyRuleSharedMemberResponseBodyData {
	s.MemberList = v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBodyData) Validate() error {
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

type ListPrivacyRuleSharedMemberResponseBodyDataMemberList struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Shared     *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
}

func (s ListPrivacyRuleSharedMemberResponseBodyDataMemberList) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleSharedMemberResponseBodyDataMemberList) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) GetMemberId() *string {
	return s.MemberId
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) GetMemberName() *string {
	return s.MemberName
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) GetShared() *string {
	return s.Shared
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) SetMemberId(v string) *ListPrivacyRuleSharedMemberResponseBodyDataMemberList {
	s.MemberId = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) SetMemberName(v string) *ListPrivacyRuleSharedMemberResponseBodyDataMemberList {
	s.MemberName = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) SetShared(v string) *ListPrivacyRuleSharedMemberResponseBodyDataMemberList {
	s.Shared = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponseBodyDataMemberList) Validate() error {
	return dara.Validate(s)
}
