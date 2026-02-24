// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorAccounts(v []*CreateAggregatorRequestAggregatorAccounts) *CreateAggregatorRequest
	GetAggregatorAccounts() []*CreateAggregatorRequestAggregatorAccounts
	SetAggregatorName(v string) *CreateAggregatorRequest
	GetAggregatorName() *string
	SetAggregatorType(v string) *CreateAggregatorRequest
	GetAggregatorType() *string
	SetClientToken(v string) *CreateAggregatorRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAggregatorRequest
	GetDescription() *string
	SetFolderId(v string) *CreateAggregatorRequest
	GetFolderId() *string
	SetTag(v []*CreateAggregatorRequestTag) *CreateAggregatorRequest
	GetTag() []*CreateAggregatorRequestTag
}

type CreateAggregatorRequest struct {
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
	AggregatorAccounts []*CreateAggregatorRequestAggregatorAccounts `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty" type:"Repeated"`
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
	Tag []*CreateAggregatorRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateAggregatorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregatorRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregatorRequest) GetAggregatorAccounts() []*CreateAggregatorRequestAggregatorAccounts {
	return s.AggregatorAccounts
}

func (s *CreateAggregatorRequest) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *CreateAggregatorRequest) GetAggregatorType() *string {
	return s.AggregatorType
}

func (s *CreateAggregatorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregatorRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregatorRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateAggregatorRequest) GetTag() []*CreateAggregatorRequestTag {
	return s.Tag
}

func (s *CreateAggregatorRequest) SetAggregatorAccounts(v []*CreateAggregatorRequestAggregatorAccounts) *CreateAggregatorRequest {
	s.AggregatorAccounts = v
	return s
}

func (s *CreateAggregatorRequest) SetAggregatorName(v string) *CreateAggregatorRequest {
	s.AggregatorName = &v
	return s
}

func (s *CreateAggregatorRequest) SetAggregatorType(v string) *CreateAggregatorRequest {
	s.AggregatorType = &v
	return s
}

func (s *CreateAggregatorRequest) SetClientToken(v string) *CreateAggregatorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregatorRequest) SetDescription(v string) *CreateAggregatorRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregatorRequest) SetFolderId(v string) *CreateAggregatorRequest {
	s.FolderId = &v
	return s
}

func (s *CreateAggregatorRequest) SetTag(v []*CreateAggregatorRequestTag) *CreateAggregatorRequest {
	s.Tag = v
	return s
}

func (s *CreateAggregatorRequest) Validate() error {
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

type CreateAggregatorRequestAggregatorAccounts struct {
	// The member ID. For more information about how to obtain the member ID, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// example:
	//
	// 171322098523****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The member name. For more information about how to obtain the member name, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// example:
	//
	// Alice
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The affiliation of the member. Only `ResourceDirectory` is supported.
	//
	// example:
	//
	// ResourceDirectory
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s CreateAggregatorRequestAggregatorAccounts) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregatorRequestAggregatorAccounts) GoString() string {
	return s.String()
}

func (s *CreateAggregatorRequestAggregatorAccounts) GetAccountId() *int64 {
	return s.AccountId
}

func (s *CreateAggregatorRequestAggregatorAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAggregatorRequestAggregatorAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateAggregatorRequestAggregatorAccounts) SetAccountId(v int64) *CreateAggregatorRequestAggregatorAccounts {
	s.AccountId = &v
	return s
}

func (s *CreateAggregatorRequestAggregatorAccounts) SetAccountName(v string) *CreateAggregatorRequestAggregatorAccounts {
	s.AccountName = &v
	return s
}

func (s *CreateAggregatorRequestAggregatorAccounts) SetAccountType(v string) *CreateAggregatorRequestAggregatorAccounts {
	s.AccountType = &v
	return s
}

func (s *CreateAggregatorRequestAggregatorAccounts) Validate() error {
	return dara.Validate(s)
}

type CreateAggregatorRequestTag struct {
	// The tag key of the resource. You can specify a maximum of 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http\\:// or https\\://.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify a maximum of 20 tag values. The tag value can be an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAggregatorRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregatorRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAggregatorRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAggregatorRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAggregatorRequestTag) SetKey(v string) *CreateAggregatorRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAggregatorRequestTag) SetValue(v string) *CreateAggregatorRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAggregatorRequestTag) Validate() error {
	return dara.Validate(s)
}
