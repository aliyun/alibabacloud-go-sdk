// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyRecordListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryGateVerifyRecordListResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryGateVerifyRecordListResponseBody
	GetCode() *string
	SetData(v bool) *QueryGateVerifyRecordListResponseBody
	GetData() *bool
}

type QueryGateVerifyRecordListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryGateVerifyRecordListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyRecordListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyRecordListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGateVerifyRecordListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGateVerifyRecordListResponseBody) GetData() *bool {
	return s.Data
}

func (s *QueryGateVerifyRecordListResponseBody) SetRequestId(v string) *QueryGateVerifyRecordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGateVerifyRecordListResponseBody) SetCode(v string) *QueryGateVerifyRecordListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyRecordListResponseBody) SetData(v bool) *QueryGateVerifyRecordListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryGateVerifyRecordListResponseBody) Validate() error {
	return dara.Validate(s)
}
