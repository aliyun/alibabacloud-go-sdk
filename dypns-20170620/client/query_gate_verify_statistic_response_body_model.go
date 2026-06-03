// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryGateVerifyStatisticResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryGateVerifyStatisticResponseBody
	GetCode() *string
	SetData(v bool) *QueryGateVerifyStatisticResponseBody
	GetData() *bool
}

type QueryGateVerifyStatisticResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryGateVerifyStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGateVerifyStatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGateVerifyStatisticResponseBody) GetData() *bool {
	return s.Data
}

func (s *QueryGateVerifyStatisticResponseBody) SetRequestId(v string) *QueryGateVerifyStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGateVerifyStatisticResponseBody) SetCode(v string) *QueryGateVerifyStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyStatisticResponseBody) SetData(v bool) *QueryGateVerifyStatisticResponseBody {
	s.Data = &v
	return s
}

func (s *QueryGateVerifyStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
