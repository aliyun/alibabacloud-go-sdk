// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupExistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetGroupExistResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetGroupExistResponseBody
	GetCode() *string
	SetData(v string) *GetGroupExistResponseBody
	GetData() *string
	SetMessage(v string) *GetGroupExistResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGroupExistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGroupExistResponseBody
	GetSuccess() *bool
}

type GetGroupExistResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// False
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGroupExistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupExistResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupExistResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetGroupExistResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGroupExistResponseBody) GetData() *string {
	return s.Data
}

func (s *GetGroupExistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGroupExistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupExistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGroupExistResponseBody) SetAccessDeniedDetail(v string) *GetGroupExistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetGroupExistResponseBody) SetCode(v string) *GetGroupExistResponseBody {
	s.Code = &v
	return s
}

func (s *GetGroupExistResponseBody) SetData(v string) *GetGroupExistResponseBody {
	s.Data = &v
	return s
}

func (s *GetGroupExistResponseBody) SetMessage(v string) *GetGroupExistResponseBody {
	s.Message = &v
	return s
}

func (s *GetGroupExistResponseBody) SetRequestId(v string) *GetGroupExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupExistResponseBody) SetSuccess(v bool) *GetGroupExistResponseBody {
	s.Success = &v
	return s
}

func (s *GetGroupExistResponseBody) Validate() error {
	return dara.Validate(s)
}
