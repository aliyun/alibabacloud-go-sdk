// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregatorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorAccountsShrink(v string) *UpdateAggregatorShrinkRequest
	GetAggregatorAccountsShrink() *string
	SetAggregatorId(v string) *UpdateAggregatorShrinkRequest
	GetAggregatorId() *string
	SetAggregatorName(v string) *UpdateAggregatorShrinkRequest
	GetAggregatorName() *string
	SetClientToken(v string) *UpdateAggregatorShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAggregatorShrinkRequest
	GetDescription() *string
	SetFolderId(v string) *UpdateAggregatorShrinkRequest
	GetFolderId() *string
	SetTagShrink(v string) *UpdateAggregatorShrinkRequest
	GetTagShrink() *string
}

type UpdateAggregatorShrinkRequest struct {
	// The members in the account group.
	//
	// >  When you modify the configurations of an account group, this parameter can be left empty. In this case, the member list is not updated. If you want to update the member list, you must configure both the `AccountId` and `AccountType` parameters.
	//
	// if can be null:
	// false
	AggregatorAccountsShrink *string `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty"`
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-dacf86d8314e00eb****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The name of the account group.
	//
	// For more information about how to obtain the name of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// example:
	//
	// Test_Group
	AggregatorName *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the account group.
	//
	// For more information about how to obtain the description of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// example:
	//
	// Test_Aggregator_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The folder ID. Separate multiple folder IDs with commas (,).
	//
	// example:
	//
	// fd-brHdgv****,fd-brHdgk****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s UpdateAggregatorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregatorShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorShrinkRequest) GetAggregatorAccountsShrink() *string {
	return s.AggregatorAccountsShrink
}

func (s *UpdateAggregatorShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregatorShrinkRequest) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *UpdateAggregatorShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregatorShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregatorShrinkRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *UpdateAggregatorShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *UpdateAggregatorShrinkRequest) SetAggregatorAccountsShrink(v string) *UpdateAggregatorShrinkRequest {
	s.AggregatorAccountsShrink = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetAggregatorId(v string) *UpdateAggregatorShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetAggregatorName(v string) *UpdateAggregatorShrinkRequest {
	s.AggregatorName = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetClientToken(v string) *UpdateAggregatorShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetDescription(v string) *UpdateAggregatorShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetFolderId(v string) *UpdateAggregatorShrinkRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetTagShrink(v string) *UpdateAggregatorShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
