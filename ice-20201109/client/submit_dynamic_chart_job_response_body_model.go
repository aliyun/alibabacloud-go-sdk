// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicChartJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitDynamicChartJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitDynamicChartJobResponseBody
	GetRequestId() *string
}

type SubmitDynamicChartJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDynamicChartJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicChartJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDynamicChartJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDynamicChartJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDynamicChartJobResponseBody) SetJobId(v string) *SubmitDynamicChartJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDynamicChartJobResponseBody) SetRequestId(v string) *SubmitDynamicChartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDynamicChartJobResponseBody) Validate() error {
	return dara.Validate(s)
}
