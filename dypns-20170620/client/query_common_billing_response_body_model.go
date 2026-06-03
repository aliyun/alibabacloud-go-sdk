// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonBillingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCommonBillingResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryCommonBillingResponseBody
	GetCode() *string
	SetData(v string) *QueryCommonBillingResponseBody
	GetData() *string
}

type QueryCommonBillingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCommonBillingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonBillingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommonBillingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCommonBillingResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCommonBillingResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCommonBillingResponseBody) SetRequestId(v string) *QueryCommonBillingResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommonBillingResponseBody) SetCode(v string) *QueryCommonBillingResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCommonBillingResponseBody) SetData(v string) *QueryCommonBillingResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCommonBillingResponseBody) Validate() error {
	return dara.Validate(s)
}
