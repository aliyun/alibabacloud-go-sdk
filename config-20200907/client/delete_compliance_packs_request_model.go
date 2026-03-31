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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance package. Separate multiple compliance package IDs with commas (,).
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackIds *string `json:"CompliancePackIds,omitempty" xml:"CompliancePackIds,omitempty"`
	// Specifies whether to delete the rules in the compliance package. Valid values:
	//
	// 	- true: The rules are deleted.
	//
	// 	- false (default): The rules are not deleted.
	//
	// example:
	//
	// false
	DeleteRule *bool `json:"DeleteRule,omitempty" xml:"DeleteRule,omitempty"`
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
