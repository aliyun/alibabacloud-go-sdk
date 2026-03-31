// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerReportSendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportTemplateId(v string) *TriggerReportSendRequest
	GetReportTemplateId() *string
}

type TriggerReportSendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// crt-xxx
	ReportTemplateId *string `json:"ReportTemplateId,omitempty" xml:"ReportTemplateId,omitempty"`
}

func (s TriggerReportSendRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerReportSendRequest) GoString() string {
	return s.String()
}

func (s *TriggerReportSendRequest) GetReportTemplateId() *string {
	return s.ReportTemplateId
}

func (s *TriggerReportSendRequest) SetReportTemplateId(v string) *TriggerReportSendRequest {
	s.ReportTemplateId = &v
	return s
}

func (s *TriggerReportSendRequest) Validate() error {
	return dara.Validate(s)
}
