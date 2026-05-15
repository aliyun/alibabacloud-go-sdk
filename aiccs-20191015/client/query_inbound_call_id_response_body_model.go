// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInboundCallIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInboundCallIdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryInboundCallIdResponseBody
	GetCode() *string
	SetData(v string) *QueryInboundCallIdResponseBody
	GetData() *string
	SetMessage(v string) *QueryInboundCallIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInboundCallIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInboundCallIdResponseBody
	GetSuccess() *bool
}

type QueryInboundCallIdResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123^456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInboundCallIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInboundCallIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInboundCallIdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInboundCallIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInboundCallIdResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryInboundCallIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInboundCallIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInboundCallIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInboundCallIdResponseBody) SetAccessDeniedDetail(v string) *QueryInboundCallIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInboundCallIdResponseBody) SetCode(v string) *QueryInboundCallIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInboundCallIdResponseBody) SetData(v string) *QueryInboundCallIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryInboundCallIdResponseBody) SetMessage(v string) *QueryInboundCallIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInboundCallIdResponseBody) SetRequestId(v string) *QueryInboundCallIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInboundCallIdResponseBody) SetSuccess(v bool) *QueryInboundCallIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInboundCallIdResponseBody) Validate() error {
	return dara.Validate(s)
}
