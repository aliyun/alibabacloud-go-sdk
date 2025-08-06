// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlipayUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GetAlipayUrlRequest
	GetBizType() *string
	SetOrderId(v int64) *GetAlipayUrlRequest
	GetOrderId() *int64
	SetReturnUrl(v string) *GetAlipayUrlRequest
	GetReturnUrl() *string
	SetType(v string) *GetAlipayUrlRequest
	GetType() *string
}

type GetAlipayUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.beian_assist
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 206129201170307
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// https://nfyt.lzsgtghchy.com:502/sigin/
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// web
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlipayUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlipayUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlRequest) GetBizType() *string {
	return s.BizType
}

func (s *GetAlipayUrlRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetAlipayUrlRequest) GetReturnUrl() *string {
	return s.ReturnUrl
}

func (s *GetAlipayUrlRequest) GetType() *string {
	return s.Type
}

func (s *GetAlipayUrlRequest) SetBizType(v string) *GetAlipayUrlRequest {
	s.BizType = &v
	return s
}

func (s *GetAlipayUrlRequest) SetOrderId(v int64) *GetAlipayUrlRequest {
	s.OrderId = &v
	return s
}

func (s *GetAlipayUrlRequest) SetReturnUrl(v string) *GetAlipayUrlRequest {
	s.ReturnUrl = &v
	return s
}

func (s *GetAlipayUrlRequest) SetType(v string) *GetAlipayUrlRequest {
	s.Type = &v
	return s
}

func (s *GetAlipayUrlRequest) Validate() error {
	return dara.Validate(s)
}
