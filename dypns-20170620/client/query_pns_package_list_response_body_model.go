// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPnsPackageListResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryPnsPackageListResponseBody
	GetCode() *string
	SetData(v string) *QueryPnsPackageListResponseBody
	GetData() *string
}

type QueryPnsPackageListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPnsPackageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPnsPackageListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPnsPackageListResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryPnsPackageListResponseBody) SetRequestId(v string) *QueryPnsPackageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPnsPackageListResponseBody) SetCode(v string) *QueryPnsPackageListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPnsPackageListResponseBody) SetData(v string) *QueryPnsPackageListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryPnsPackageListResponseBody) Validate() error {
	return dara.Validate(s)
}
