// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCompliancePackReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GenerateCompliancePackReportRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *GenerateCompliancePackReportRequest
	GetCompliancePackId() *string
}

type GenerateCompliancePackReportRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GenerateCompliancePackReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateCompliancePackReportRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GenerateCompliancePackReportRequest) SetClientToken(v string) *GenerateCompliancePackReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateCompliancePackReportRequest) SetCompliancePackId(v string) *GenerateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateCompliancePackReportRequest) Validate() error {
	return dara.Validate(s)
}
