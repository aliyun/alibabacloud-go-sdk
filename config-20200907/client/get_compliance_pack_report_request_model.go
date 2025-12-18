// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *GetCompliancePackReportRequest
	GetCompliancePackId() *string
}

type GetCompliancePackReportRequest struct {
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetCompliancePackReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetCompliancePackReportRequest) SetCompliancePackId(v string) *GetCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackReportRequest) Validate() error {
	return dara.Validate(s)
}
