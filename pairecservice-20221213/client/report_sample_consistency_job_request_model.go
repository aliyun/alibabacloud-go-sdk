// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportSampleConsistencyJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReportSampleConsistencyJobRequest
	GetInstanceId() *string
}

type ReportSampleConsistencyJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReportSampleConsistencyJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportSampleConsistencyJobRequest) GoString() string {
	return s.String()
}

func (s *ReportSampleConsistencyJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReportSampleConsistencyJobRequest) SetInstanceId(v string) *ReportSampleConsistencyJobRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportSampleConsistencyJobRequest) Validate() error {
	return dara.Validate(s)
}
