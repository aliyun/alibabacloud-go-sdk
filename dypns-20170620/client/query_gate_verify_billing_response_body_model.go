// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyBillingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryGateVerifyBillingResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryGateVerifyBillingResponseBody
	GetCode() *string
	SetData(v bool) *QueryGateVerifyBillingResponseBody
	GetData() *bool
}

type QueryGateVerifyBillingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryGateVerifyBillingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGateVerifyBillingResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGateVerifyBillingResponseBody) GetData() *bool {
	return s.Data
}

func (s *QueryGateVerifyBillingResponseBody) SetRequestId(v string) *QueryGateVerifyBillingResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGateVerifyBillingResponseBody) SetCode(v string) *QueryGateVerifyBillingResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyBillingResponseBody) SetData(v bool) *QueryGateVerifyBillingResponseBody {
	s.Data = &v
	return s
}

func (s *QueryGateVerifyBillingResponseBody) Validate() error {
	return dara.Validate(s)
}
