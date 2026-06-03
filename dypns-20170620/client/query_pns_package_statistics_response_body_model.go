// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPnsPackageStatisticsResponseBody
	GetRequestId() *string
	SetCode(v string) *QueryPnsPackageStatisticsResponseBody
	GetCode() *string
	SetData(v string) *QueryPnsPackageStatisticsResponseBody
	GetData() *string
}

type QueryPnsPackageStatisticsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPnsPackageStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPnsPackageStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPnsPackageStatisticsResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryPnsPackageStatisticsResponseBody) SetRequestId(v string) *QueryPnsPackageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPnsPackageStatisticsResponseBody) SetCode(v string) *QueryPnsPackageStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPnsPackageStatisticsResponseBody) SetData(v string) *QueryPnsPackageStatisticsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryPnsPackageStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
