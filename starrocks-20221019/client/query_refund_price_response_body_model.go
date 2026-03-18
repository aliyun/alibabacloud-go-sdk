// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRefundPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRefundPriceResponseBody
	GetAccessDeniedDetail() *string
	SetData(v float64) *QueryRefundPriceResponseBody
	GetData() *float64
	SetErrCode(v string) *QueryRefundPriceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryRefundPriceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryRefundPriceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryRefundPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRefundPriceResponseBody
	GetSuccess() *bool
}

type QueryRefundPriceResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// example:
	//
	// 18837
	Data *float64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// None
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 8C69A6D0-49B7-54B9-BF21-9AF52172A5F7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryRefundPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRefundPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRefundPriceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRefundPriceResponseBody) GetData() *float64 {
	return s.Data
}

func (s *QueryRefundPriceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryRefundPriceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryRefundPriceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryRefundPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRefundPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRefundPriceResponseBody) SetAccessDeniedDetail(v string) *QueryRefundPriceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRefundPriceResponseBody) SetData(v float64) *QueryRefundPriceResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRefundPriceResponseBody) SetErrCode(v string) *QueryRefundPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryRefundPriceResponseBody) SetErrMessage(v string) *QueryRefundPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryRefundPriceResponseBody) SetHttpStatusCode(v int32) *QueryRefundPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRefundPriceResponseBody) SetRequestId(v string) *QueryRefundPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRefundPriceResponseBody) SetSuccess(v bool) *QueryRefundPriceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRefundPriceResponseBody) Validate() error {
	return dara.Validate(s)
}
