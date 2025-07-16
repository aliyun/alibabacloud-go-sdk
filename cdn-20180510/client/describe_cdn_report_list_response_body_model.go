// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeCdnReportListResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeCdnReportListResponseBody
	GetRequestId() *string
}

type DescribeCdnReportListResponseBody struct {
	// The information about the report that is queried.
	//
	// example:
	//
	// "data":[{"reportId":1,"deliver":{"report":{"title":"DomainPvUv","format":"chart","shape":"line","xAxis":"ds","yAxis":"cnt","legend":"cnt_type","header":["ds","cnt_type","cnt"]}}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeCdnReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnReportListResponseBody) SetContent(v string) *DescribeCdnReportListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnReportListResponseBody) SetRequestId(v string) *DescribeCdnReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnReportListResponseBody) Validate() error {
	return dara.Validate(s)
}
