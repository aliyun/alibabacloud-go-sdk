// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTagStatusResponseBody
	GetCode() *string
	SetData(v bool) *QueryTagStatusResponseBody
	GetData() *bool
	SetRequestId(v string) *QueryTagStatusResponseBody
	GetRequestId() *string
}

type QueryTagStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTagStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTagStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTagStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *QueryTagStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTagStatusResponseBody) SetCode(v string) *QueryTagStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagStatusResponseBody) SetData(v bool) *QueryTagStatusResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTagStatusResponseBody) SetRequestId(v string) *QueryTagStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
