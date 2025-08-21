// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateGovernanceReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateGovernanceReportResponseBody
	GetRequestId() *string
	SetState(v string) *GenerateGovernanceReportResponseBody
	GetState() *string
}

type GenerateGovernanceReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 492E130C-76D3-55D5-BE5C-C023E431369A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The report generation status. Valid values:
	//
	// 	- Started: The system starts to generate a governance report.
	//
	// 	- Progressing: The system is generating a governance report.
	//
	// 	- Completed: A governance report is generated.
	//
	// example:
	//
	// Started
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GenerateGovernanceReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateGovernanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateGovernanceReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateGovernanceReportResponseBody) GetState() *string {
	return s.State
}

func (s *GenerateGovernanceReportResponseBody) SetRequestId(v string) *GenerateGovernanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateGovernanceReportResponseBody) SetState(v string) *GenerateGovernanceReportResponseBody {
	s.State = &v
	return s
}

func (s *GenerateGovernanceReportResponseBody) Validate() error {
	return dara.Validate(s)
}
