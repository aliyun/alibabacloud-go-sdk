// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPnsConfigResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryPnsConfigResponseBody
	GetCode() *string
	SetData(v bool) *QueryPnsConfigResponseBody
	GetData() *bool
}

type QueryPnsConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPnsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPnsConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPnsConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPnsConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *QueryPnsConfigResponseBody) SetRequestId(v string) *QueryPnsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPnsConfigResponseBody) SetCode(v string) *QueryPnsConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPnsConfigResponseBody) SetData(v bool) *QueryPnsConfigResponseBody {
	s.Data = &v
	return s
}

func (s *QueryPnsConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
