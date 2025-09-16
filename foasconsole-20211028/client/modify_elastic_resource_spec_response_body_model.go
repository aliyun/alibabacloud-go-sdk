// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticResourceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyElasticResourceSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyElasticResourceSpecResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyElasticResourceSpecResponseBody
	GetSuccess() *bool
}

type ModifyElasticResourceSpecResponseBody struct {
	// example:
	//
	// 211473228320700
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// B21DC47E-8928-199A-9F32-36D45E4693B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyElasticResourceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticResourceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticResourceSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyElasticResourceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticResourceSpecResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyElasticResourceSpecResponseBody) SetOrderId(v int64) *ModifyElasticResourceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyElasticResourceSpecResponseBody) SetRequestId(v string) *ModifyElasticResourceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticResourceSpecResponseBody) SetSuccess(v bool) *ModifyElasticResourceSpecResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyElasticResourceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
