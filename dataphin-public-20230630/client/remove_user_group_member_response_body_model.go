// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveUserGroupMemberResponseBody
	GetCode() *string
	SetData(v bool) *RemoveUserGroupMemberResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *RemoveUserGroupMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveUserGroupMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveUserGroupMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveUserGroupMemberResponseBody
	GetSuccess() *bool
}

type RemoveUserGroupMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveUserGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveUserGroupMemberResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveUserGroupMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveUserGroupMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveUserGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveUserGroupMemberResponseBody) SetCode(v string) *RemoveUserGroupMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetData(v bool) *RemoveUserGroupMemberResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetHttpStatusCode(v int32) *RemoveUserGroupMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetMessage(v string) *RemoveUserGroupMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetRequestId(v string) *RemoveUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetSuccess(v bool) *RemoveUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
