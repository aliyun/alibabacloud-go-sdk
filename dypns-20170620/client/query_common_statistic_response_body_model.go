// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCommonStatisticResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryCommonStatisticResponseBody
	GetCode() *string
	SetData(v string) *QueryCommonStatisticResponseBody
	GetData() *string
}

type QueryCommonStatisticResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCommonStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommonStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCommonStatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCommonStatisticResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryCommonStatisticResponseBody) SetRequestId(v string) *QueryCommonStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommonStatisticResponseBody) SetCode(v string) *QueryCommonStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCommonStatisticResponseBody) SetData(v string) *QueryCommonStatisticResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCommonStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
