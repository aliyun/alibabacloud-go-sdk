// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCompliancePackReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *GenerateCompliancePackReportResponseBody
	GetCompliancePackId() *string
	SetRequestId(v string) *GenerateCompliancePackReportResponseBody
	GetRequestId() *string
}

type GenerateCompliancePackReportResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCompliancePackReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportResponseBody) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GenerateCompliancePackReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCompliancePackReportResponseBody) SetCompliancePackId(v string) *GenerateCompliancePackReportResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateCompliancePackReportResponseBody) SetRequestId(v string) *GenerateCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCompliancePackReportResponseBody) Validate() error {
	return dara.Validate(s)
}
