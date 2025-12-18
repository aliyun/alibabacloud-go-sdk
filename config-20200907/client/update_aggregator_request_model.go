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
	// >  When you modify the configurations of an account group, this parameter can be left empty. In this case, the member list is not updated. If you want to update the member list, you must configure both the `AccountId` and `AccountType` parameters.
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
	// >  If you want to update the member list, you must configure both the `AccountId` and `AccountType` parameters.
	//
	// example:
	//
	// 173808452267****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The display name of the member.
	//
	// For more information about how to obtain the name of a member, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// >  If you want to update the member list, you must configure both the `AccountId` and `AccountType` parameters.
	//
	// example:
	//
	// Tony
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The resource directory to which the member belongs. Valid value: ResourceDirectory. ResourceDirectory indicates that the member belongs to a resource directory.
	//
	// >  If you want to update the member list, you must configure both the `AccountId` and `AccountType` parameters.
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
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs`:. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
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
