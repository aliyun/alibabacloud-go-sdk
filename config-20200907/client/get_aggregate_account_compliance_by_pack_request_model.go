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
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-04b3fd170e340007****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-541e626622af0087****
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
