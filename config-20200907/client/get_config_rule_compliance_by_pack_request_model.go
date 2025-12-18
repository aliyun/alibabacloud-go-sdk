// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleComplianceByPackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *GetConfigRuleComplianceByPackRequest
	GetCompliancePackId() *string
}

type GetConfigRuleComplianceByPackRequest struct {
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetConfigRuleComplianceByPackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetConfigRuleComplianceByPackRequest) SetCompliancePackId(v string) *GetConfigRuleComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetConfigRuleComplianceByPackRequest) Validate() error {
	return dara.Validate(s)
}
