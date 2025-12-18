// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateCompliancePacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DeleteAggregateCompliancePacksRequest
	GetAggregatorId() *string
	SetClientToken(v string) *DeleteAggregateCompliancePacksRequest
	GetClientToken() *string
	SetCompliancePackIds(v string) *DeleteAggregateCompliancePacksRequest
	GetCompliancePackIds() *string
	SetDeleteRule(v bool) *DeleteAggregateCompliancePacksRequest
	GetDeleteRule() *bool
}

type DeleteAggregateCompliancePacksRequest struct {
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance package. Separate multiple compliance package IDs with commas (,).
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
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

func (s DeleteAggregateCompliancePacksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeleteAggregateCompliancePacksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAggregateCompliancePacksRequest) GetCompliancePackIds() *string {
	return s.CompliancePackIds
}

func (s *DeleteAggregateCompliancePacksRequest) GetDeleteRule() *bool {
	return s.DeleteRule
}

func (s *DeleteAggregateCompliancePacksRequest) SetAggregatorId(v string) *DeleteAggregateCompliancePacksRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetClientToken(v string) *DeleteAggregateCompliancePacksRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetCompliancePackIds(v string) *DeleteAggregateCompliancePacksRequest {
	s.CompliancePackIds = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetDeleteRule(v bool) *DeleteAggregateCompliancePacksRequest {
	s.DeleteRule = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) Validate() error {
	return dara.Validate(s)
}
