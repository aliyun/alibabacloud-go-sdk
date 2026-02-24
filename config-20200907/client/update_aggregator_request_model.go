// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorAccounts(v []*UpdateAggregatorRequestAggregatorAccounts) *UpdateAggregatorRequest
	GetAggregatorAccounts() []*UpdateAggregatorRequestAggregatorAccounts
	SetAggregatorId(v string) *UpdateAggregatorRequest
	GetAggregatorId() *string
	SetAggregatorName(v string) *UpdateAggregatorRequest
	GetAggregatorName() *string
	SetClientToken(v string) *UpdateAggregatorRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAggregatorRequest
	GetDescription() *string
	SetFolderId(v string) *UpdateAggregatorRequest
	GetFolderId() *string
	SetTag(v []*UpdateAggregatorRequestTag) *UpdateAggregatorRequest
	GetTag() []*UpdateAggregatorRequestTag
}

type UpdateAggregatorRequest struct {
	// The members in the account group.
	//
	// > You can leave this parameter empty to skip updating the member list. To update the member list, you must specify both `AccountId` and `AccountType`.
	//
	// if can be null:
	// false
	AggregatorAccounts []*UpdateAggregatorRequestAggregatorAccounts `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty" type:"Repeated"`
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
	// A client token that ensures the idempotence of the request. Generate a unique token for each request. The token can contain only ASCII characters and must be no more than 64 characters in length.
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
	// 测试组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the folder. You can enter multiple folder IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// fd-brHdgv****,fd-brHdgk****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated and no longer takes effect. Ignore this parameter.
	//
	// You can attach up to 20 tags.
	Tag []*UpdateAggregatorRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UpdateAggregatorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregatorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorRequest) GetAggregatorAccounts() []*UpdateAggregatorRequestAggregatorAccounts {
	return s.AggregatorAccounts
}

func (s *UpdateAggregatorRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregatorRequest) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *UpdateAggregatorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregatorRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregatorRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *UpdateAggregatorRequest) GetTag() []*UpdateAggregatorRequestTag {
	return s.Tag
}

func (s *UpdateAggregatorRequest) SetAggregatorAccounts(v []*UpdateAggregatorRequestAggregatorAccounts) *UpdateAggregatorRequest {
	s.AggregatorAccounts = v
	return s
}

func (s *UpdateAggregatorRequest) SetAggregatorId(v string) *UpdateAggregatorRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregatorRequest) SetAggregatorName(v string) *UpdateAggregatorRequest {
	s.AggregatorName = &v
	return s
}

func (s *UpdateAggregatorRequest) SetClientToken(v string) *UpdateAggregatorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregatorRequest) SetDescription(v string) *UpdateAggregatorRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregatorRequest) SetFolderId(v string) *UpdateAggregatorRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateAggregatorRequest) SetTag(v []*UpdateAggregatorRequestTag) *UpdateAggregatorRequest {
	s.Tag = v
	return s
}

func (s *UpdateAggregatorRequest) Validate() error {
	if s.AggregatorAccounts != nil {
		for _, item := range s.AggregatorAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAggregatorRequestAggregatorAccounts struct {
	// The ID of the member.
	//
	// For more information about how to obtain the ID of a member, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// > To update the member list, you must specify both `AccountId` and `AccountType`.
	//
	// example:
	//
	// 173808452267****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member.
	//
	// For more information about how to obtain the name of a member, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// > To update the member list, you must specify both `AccountId` and `AccountType`.
	//
	// example:
	//
	// Tony
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The affiliation of the member. Only ResourceDirectory is supported.
	//
	// > To update the member list, you must specify both `AccountId` and `AccountType`.
	//
	// example:
	//
	// ResourceDirectory
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s UpdateAggregatorRequestAggregatorAccounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregatorRequestAggregatorAccounts) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorRequestAggregatorAccounts) GetAccountId() *int64 {
	return s.AccountId
}

func (s *UpdateAggregatorRequestAggregatorAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *UpdateAggregatorRequestAggregatorAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *UpdateAggregatorRequestAggregatorAccounts) SetAccountId(v int64) *UpdateAggregatorRequestAggregatorAccounts {
	s.AccountId = &v
	return s
}

func (s *UpdateAggregatorRequestAggregatorAccounts) SetAccountName(v string) *UpdateAggregatorRequestAggregatorAccounts {
	s.AccountName = &v
	return s
}

func (s *UpdateAggregatorRequestAggregatorAccounts) SetAccountType(v string) *UpdateAggregatorRequestAggregatorAccounts {
	s.AccountType = &v
	return s
}

func (s *UpdateAggregatorRequestAggregatorAccounts) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregatorRequestTag struct {
	// The key of the tag. A tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateAggregatorRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregatorRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateAggregatorRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateAggregatorRequestTag) SetKey(v string) *UpdateAggregatorRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateAggregatorRequestTag) SetValue(v string) *UpdateAggregatorRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateAggregatorRequestTag) Validate() error {
	return dara.Validate(s)
}
