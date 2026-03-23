// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *QueryJobOrderRequest
	GetAppCode() *string
	SetAppName(v string) *QueryJobOrderRequest
	GetAppName() *string
	SetOrderId(v int64) *QueryJobOrderRequest
	GetOrderId() *int64
	SetRequestId(v string) *QueryJobOrderRequest
	GetRequestId() *string
}

type QueryJobOrderRequest struct {
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryJobOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryJobOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryJobOrderRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryJobOrderRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryJobOrderRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *QueryJobOrderRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryJobOrderRequest) SetAppCode(v string) *QueryJobOrderRequest {
	s.AppCode = &v
	return s
}

func (s *QueryJobOrderRequest) SetAppName(v string) *QueryJobOrderRequest {
	s.AppName = &v
	return s
}

func (s *QueryJobOrderRequest) SetOrderId(v int64) *QueryJobOrderRequest {
	s.OrderId = &v
	return s
}

func (s *QueryJobOrderRequest) SetRequestId(v string) *QueryJobOrderRequest {
	s.RequestId = &v
	return s
}

func (s *QueryJobOrderRequest) Validate() error {
	return dara.Validate(s)
}
