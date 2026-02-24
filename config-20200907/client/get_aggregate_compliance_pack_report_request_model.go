// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateCompliancePackReportRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *GetAggregateCompliancePackReportRequest
	GetCompliancePackId() *string
}

type GetAggregateCompliancePackReportRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateCompliancePackReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateCompliancePackReportRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateCompliancePackReportRequest) SetAggregatorId(v string) *GetAggregateCompliancePackReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackReportRequest) SetCompliancePackId(v string) *GetAggregateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackReportRequest) Validate() error {
	return dara.Validate(s)
}
