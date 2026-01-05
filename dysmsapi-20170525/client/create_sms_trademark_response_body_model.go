// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTrademarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSmsTrademarkResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateSmsTrademarkResponseBody
	GetCode() *string
	SetData(v string) *CreateSmsTrademarkResponseBody
	GetData() *string
	SetMessage(v string) *CreateSmsTrademarkResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSmsTrademarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSmsTrademarkResponseBody
	GetSuccess() *bool
}

type CreateSmsTrademarkResponseBody struct {
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

func (s CreateSmsTrademarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTrademarkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsTrademarkResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSmsTrademarkResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsTrademarkResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSmsTrademarkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsTrademarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmsTrademarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSmsTrademarkResponseBody) SetAccessDeniedDetail(v string) *CreateSmsTrademarkResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSmsTrademarkResponseBody) SetCode(v string) *CreateSmsTrademarkResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsTrademarkResponseBody) SetData(v string) *CreateSmsTrademarkResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSmsTrademarkResponseBody) SetMessage(v string) *CreateSmsTrademarkResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsTrademarkResponseBody) SetRequestId(v string) *CreateSmsTrademarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsTrademarkResponseBody) SetSuccess(v bool) *CreateSmsTrademarkResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSmsTrademarkResponseBody) Validate() error {
	return dara.Validate(s)
}
