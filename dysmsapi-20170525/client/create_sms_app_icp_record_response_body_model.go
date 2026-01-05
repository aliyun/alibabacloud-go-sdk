// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAppIcpRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSmsAppIcpRecordResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateSmsAppIcpRecordResponseBody
	GetCode() *string
	SetData(v string) *CreateSmsAppIcpRecordResponseBody
	GetData() *string
	SetMessage(v string) *CreateSmsAppIcpRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSmsAppIcpRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSmsAppIcpRecordResponseBody
	GetSuccess() *bool
}

type CreateSmsAppIcpRecordResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10000*******
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSmsAppIcpRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAppIcpRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsAppIcpRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSmsAppIcpRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsAppIcpRecordResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSmsAppIcpRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsAppIcpRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmsAppIcpRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSmsAppIcpRecordResponseBody) SetAccessDeniedDetail(v string) *CreateSmsAppIcpRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponseBody) SetCode(v string) *CreateSmsAppIcpRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponseBody) SetData(v string) *CreateSmsAppIcpRecordResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponseBody) SetMessage(v string) *CreateSmsAppIcpRecordResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponseBody) SetRequestId(v string) *CreateSmsAppIcpRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponseBody) SetSuccess(v bool) *CreateSmsAppIcpRecordResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSmsAppIcpRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
