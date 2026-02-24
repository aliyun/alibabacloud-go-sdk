// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregatorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorAccountsShrink(v string) *CreateAggregatorShrinkRequest
	GetAggregatorAccountsShrink() *string
	SetAggregatorName(v string) *CreateAggregatorShrinkRequest
	GetAggregatorName() *string
	SetAggregatorType(v string) *CreateAggregatorShrinkRequest
	GetAggregatorType() *string
	SetClientToken(v string) *CreateAggregatorShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAggregatorShrinkRequest
	GetDescription() *string
	SetFolderId(v string) *CreateAggregatorShrinkRequest
	GetFolderId() *string
	SetTagShrink(v string) *CreateAggregatorShrinkRequest
	GetTagShrink() *string
}

type CreateAggregatorShrinkRequest struct {
	// The member accounts of the account group.
	//
	// > - If you set `AggregatorType` to \\`RD, you can leave this parameter empty. This indicates that all members in the resource directory are added to the global account group.
	//
	// >
	//
	// > - If you set `AggregatorType` to `FOLDER`, you can leave this parameter empty. This indicates that all members in a specific folder in the resource directory are added to the folder account group.
	//
	// if can be null:
	// false
	AggregatorAccountsShrink *string `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty"`
	// The name of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Example_Aggregator
	AggregatorName *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	// The type of the account group. Valid values:
	//
	// - RD: global account group.
	//
	// - FOLDER: folder account group. You must also set the `FolderId` parameter. For more information about how to obtain a folder ID, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// - CUSTOM (default): custom account group. You must also set the `AccountId` and `AccountType` parameters for `AggregatorAccounts`.
	//
	// example:
	//
	// CUSTOM
	AggregatorType *string `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You must make sure that the token is unique for different requests. The `ClientToken` parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the account group.
	//
	// example:
	//
	// Example aggregator used to demonstrate how to create an aggregator.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the attached folder. You can specify multiple folder IDs. Separate the IDs with commas (,).
	//
	// This parameter is required if you set `AggregatorType` to `FOLDER`.
	//
	// example:
	//
	// fd-brHdgv****,fd-brHdgk****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The tags of the resource.
	//
	// You can attach a maximum of 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateAggregatorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregatorShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregatorShrinkRequest) GetAggregatorAccountsShrink() *string {
	return s.AggregatorAccountsShrink
}

func (s *CreateAggregatorShrinkRequest) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *CreateAggregatorShrinkRequest) GetAggregatorType() *string {
	return s.AggregatorType
}

func (s *CreateAggregatorShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregatorShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregatorShrinkRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateAggregatorShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateAggregatorShrinkRequest) SetAggregatorAccountsShrink(v string) *CreateAggregatorShrinkRequest {
	s.AggregatorAccountsShrink = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetAggregatorName(v string) *CreateAggregatorShrinkRequest {
	s.AggregatorName = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetAggregatorType(v string) *CreateAggregatorShrinkRequest {
	s.AggregatorType = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetClientToken(v string) *CreateAggregatorShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetDescription(v string) *CreateAggregatorShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetFolderId(v string) *CreateAggregatorShrinkRequest {
	s.FolderId = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetTagShrink(v string) *CreateAggregatorShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
