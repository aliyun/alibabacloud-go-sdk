// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPnsPackageDetailResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryPnsPackageDetailResponseBody
	GetCode() *string
	SetData(v string) *QueryPnsPackageDetailResponseBody
	GetData() *string
}

type QueryPnsPackageDetailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPnsPackageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPnsPackageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPnsPackageDetailResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryPnsPackageDetailResponseBody) SetRequestId(v string) *QueryPnsPackageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPnsPackageDetailResponseBody) SetCode(v string) *QueryPnsPackageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPnsPackageDetailResponseBody) SetData(v string) *QueryPnsPackageDetailResponseBody {
	s.Data = &v
	return s
}

func (s *QueryPnsPackageDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
