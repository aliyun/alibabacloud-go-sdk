// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationCustomizeReportChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperationCustomizeReportChartResponseBody
	GetRequestId() *string
}

type OperationCustomizeReportChartResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 432B2D4E-C8D3-52E4-9F68-35E0C05F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperationCustomizeReportChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperationCustomizeReportChartResponseBody) GoString() string {
	return s.String()
}

func (s *OperationCustomizeReportChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperationCustomizeReportChartResponseBody) SetRequestId(v string) *OperationCustomizeReportChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperationCustomizeReportChartResponseBody) Validate() error {
	return dara.Validate(s)
}
