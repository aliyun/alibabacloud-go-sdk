// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclGroupCidrsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAclGroupCidrsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *UpdateAclGroupCidrsResponseBody
	GetData() *bool
	SetErrCode(v string) *UpdateAclGroupCidrsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateAclGroupCidrsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateAclGroupCidrsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateAclGroupCidrsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAclGroupCidrsResponseBody
	GetSuccess() *bool
}

type UpdateAclGroupCidrsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// [RDS%22]
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// CB356AEE-6B3F-5FC8-9C2C-7B2D881EA9E2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAclGroupCidrsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclGroupCidrsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclGroupCidrsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAclGroupCidrsResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAclGroupCidrsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateAclGroupCidrsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateAclGroupCidrsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAclGroupCidrsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAclGroupCidrsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAclGroupCidrsResponseBody) SetAccessDeniedDetail(v string) *UpdateAclGroupCidrsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) SetData(v bool) *UpdateAclGroupCidrsResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) SetErrCode(v string) *UpdateAclGroupCidrsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) SetErrMessage(v string) *UpdateAclGroupCidrsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) SetHttpStatusCode(v int32) *UpdateAclGroupCidrsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) SetRequestId(v string) *UpdateAclGroupCidrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) SetSuccess(v bool) *UpdateAclGroupCidrsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAclGroupCidrsResponseBody) Validate() error {
	return dara.Validate(s)
}
