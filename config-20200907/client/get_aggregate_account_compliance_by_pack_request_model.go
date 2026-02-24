// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateAccountComplianceByPackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateAccountComplianceByPackRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *GetAggregateAccountComplianceByPackRequest
	GetCompliancePackId() *string
}

type GetAggregateAccountComplianceByPackRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateAccountComplianceByPackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateAccountComplianceByPackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateAccountComplianceByPackRequest) SetAggregatorId(v string) *GetAggregateAccountComplianceByPackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackRequest) SetCompliancePackId(v string) *GetAggregateAccountComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackRequest) Validate() error {
	return dara.Validate(s)
}
