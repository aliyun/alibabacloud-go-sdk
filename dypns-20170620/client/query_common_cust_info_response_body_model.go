// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonCustInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCommonCustInfoResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryCommonCustInfoResponseBody
	GetCode() *string
	SetData(v string) *QueryCommonCustInfoResponseBody
	GetData() *string
}

type QueryCommonCustInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCommonCustInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonCustInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommonCustInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCommonCustInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCommonCustInfoResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCommonCustInfoResponseBody) SetRequestId(v string) *QueryCommonCustInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommonCustInfoResponseBody) SetCode(v string) *QueryCommonCustInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCommonCustInfoResponseBody) SetData(v string) *QueryCommonCustInfoResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCommonCustInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
