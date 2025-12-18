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
	// The information about the member accounts in the account group. Example:
	//
	//     [{
	//
	//     	"accountId": 171322098523****,
	//
	//     	"accountType":"ResourceDirectory",
	//
	//                     "accountName":"Alice"
	//
	//     }, {
	//
	//     	"accountId": 100532098349****,
	//
	//     	"accountType":"ResourceDirectory",
	//
	//                     "accountName":"Tom"
	//
	//     }]
	//
	// >  If `AggregatorType` is set to `RD` or `FOLDER`, this parameter can be left empty, which indicates that all accounts in the resource directory are added to the global account group.
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
	// Test_Group
	AggregatorName *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	// The type of the account group. Valid values:
	//
	// 	- RD: global account group.
	//
	// 	- FOLDER: account group of the folder.
	//
	// 	- CUSTOM (default): custom account group.
	//
	// example:
	//
	// CUSTOM
	AggregatorType *string `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the account group.
	//
	// example:
	//
	// Aggregator description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the folder to which the account group is attached. You must specify this parameter if `AggregatorType` is set to `FOLDER`. Multiple resource folder IDs should be separated by commas (,).
	//
	// example:
	//
	// fd-brHdgv****,fd-brHdgk****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
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
	// The member account ID. For more information about how to obtain the ID of a member account, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// example:
	//
	// 171322098523****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member account. For more information about how to obtain the name of a member account, see [ListAccounts](https://help.aliyun.com/document_detail/160016.html).
	//
	// example:
	//
	// Alice
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the member account. Set this parameter to ResourceDirectory.
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
