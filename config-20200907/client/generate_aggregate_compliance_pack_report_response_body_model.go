// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateCompliancePackReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *GenerateAggregateCompliancePackReportResponseBody
	GetCompliancePackId() *string
	SetRequestId(v string) *GenerateAggregateCompliancePackReportResponseBody
	GetRequestId() *string
}

type GenerateAggregateCompliancePackReportResponseBody struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAggregateCompliancePackReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAggregateCompliancePackReportResponseBody) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GenerateAggregateCompliancePackReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAggregateCompliancePackReportResponseBody) SetCompliancePackId(v string) *GenerateAggregateCompliancePackReportResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponseBody) SetRequestId(v string) *GenerateAggregateCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponseBody) Validate() error {
	return dara.Validate(s)
}
