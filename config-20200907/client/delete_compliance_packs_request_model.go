// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompliancePacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCompliancePacksRequest
	GetClientToken() *string
	SetCompliancePackIds(v string) *DeleteCompliancePacksRequest
	GetCompliancePackIds() *string
	SetDeleteRule(v bool) *DeleteCompliancePacksRequest
	GetDeleteRule() *bool
}

type DeleteCompliancePacksRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CompliancePackIds *string `json:"CompliancePackIds,omitempty" xml:"CompliancePackIds,omitempty"`
	DeleteRule        *bool   `json:"DeleteRule,omitempty" xml:"DeleteRule,omitempty"`
}

func (s DeleteCompliancePacksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCompliancePacksRequest) GetCompliancePackIds() *string {
	return s.CompliancePackIds
}

func (s *DeleteCompliancePacksRequest) GetDeleteRule() *bool {
	return s.DeleteRule
}

func (s *DeleteCompliancePacksRequest) SetClientToken(v string) *DeleteCompliancePacksRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCompliancePacksRequest) SetCompliancePackIds(v string) *DeleteCompliancePacksRequest {
	s.CompliancePackIds = &v
	return s
}

func (s *DeleteCompliancePacksRequest) SetDeleteRule(v bool) *DeleteCompliancePacksRequest {
	s.DeleteRule = &v
	return s
}

func (s *DeleteCompliancePacksRequest) Validate() error {
	return dara.Validate(s)
}
