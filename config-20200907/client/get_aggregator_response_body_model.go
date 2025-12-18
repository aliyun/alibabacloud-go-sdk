// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregator(v *GetAggregatorResponseBodyAggregator) *GetAggregatorResponseBody
	GetAggregator() *GetAggregatorResponseBodyAggregator
	SetRequestId(v string) *GetAggregatorResponseBody
	GetRequestId() *string
}

type GetAggregatorResponseBody struct {
	// The details of the account group.
	Aggregator *GetAggregatorResponseBodyAggregator `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 733DD93C-2277-4905-AE0C-0BA95C04B8BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBody) GetAggregator() *GetAggregatorResponseBodyAggregator {
	return s.Aggregator
}

func (s *GetAggregatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregatorResponseBody) SetAggregator(v *GetAggregatorResponseBodyAggregator) *GetAggregatorResponseBody {
	s.Aggregator = v
	return s
}

func (s *GetAggregatorResponseBody) SetRequestId(v string) *GetAggregatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregatorResponseBody) Validate() error {
	if s.Aggregator != nil {
		if err := s.Aggregator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregatorResponseBodyAggregator struct {
	// The ID of the management account that is used to create the account group.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The number of members in the account group.
	//
	// example:
	//
	// 2
	AggregatorAccountCount *int64 `json:"AggregatorAccountCount,omitempty" xml:"AggregatorAccountCount,omitempty"`
	// The information about the members in the account group.
	AggregatorAccounts []*GetAggregatorResponseBodyAggregatorAggregatorAccounts `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty" type:"Repeated"`
	// The timestamp generated when the account group was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1623034091000
	AggregatorCreateTimestamp *string `json:"AggregatorCreateTimestamp,omitempty" xml:"AggregatorCreateTimestamp,omitempty"`
	// The ID of the account group.
	//
	// example:
	//
	// ca-88ea626622af0055****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The name of the account group.
	//
	// example:
	//
	// Test-Group
	AggregatorName *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	// The status of the account group. Valid values:
	//
	// 	- 0: The account group is being created.
	//
	// 	- 1: The account group was created.
	//
	// example:
	//
	// 1
	AggregatorStatus *int32 `json:"AggregatorStatus,omitempty" xml:"AggregatorStatus,omitempty"`
	// The type of the account group. Valid values:
	//
	// 	- RD: a global account group.
	//
	// 	- FOLDER: an account group for a folder.
	//
	// 	- CUSTOM: a custom account group.
	//
	// example:
	//
	// CUSTOM
	AggregatorType *string `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	// The description of the account group.
	//
	// example:
	//
	// The description of the test account group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the attached folder of the account group.
	//
	// example:
	//
	// fd-brHdgv****
	FolderId   *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// tags
	Tags []*GetAggregatorResponseBodyAggregatorTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetAggregatorResponseBodyAggregator) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorResponseBodyAggregator) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBodyAggregator) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorAccountCount() *int64 {
	return s.AggregatorAccountCount
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorAccounts() []*GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	return s.AggregatorAccounts
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorCreateTimestamp() *string {
	return s.AggregatorCreateTimestamp
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorStatus() *int32 {
	return s.AggregatorStatus
}

func (s *GetAggregatorResponseBodyAggregator) GetAggregatorType() *string {
	return s.AggregatorType
}

func (s *GetAggregatorResponseBodyAggregator) GetDescription() *string {
	return s.Description
}

func (s *GetAggregatorResponseBodyAggregator) GetFolderId() *string {
	return s.FolderId
}

func (s *GetAggregatorResponseBodyAggregator) GetFolderName() *string {
	return s.FolderName
}

func (s *GetAggregatorResponseBodyAggregator) GetTags() []*GetAggregatorResponseBodyAggregatorTags {
	return s.Tags
}

func (s *GetAggregatorResponseBodyAggregator) SetAccountId(v int64) *GetAggregatorResponseBodyAggregator {
	s.AccountId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorAccountCount(v int64) *GetAggregatorResponseBodyAggregator {
	s.AggregatorAccountCount = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorAccounts(v []*GetAggregatorResponseBodyAggregatorAggregatorAccounts) *GetAggregatorResponseBodyAggregator {
	s.AggregatorAccounts = v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorCreateTimestamp(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorCreateTimestamp = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorId(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorName(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorName = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorStatus(v int32) *GetAggregatorResponseBodyAggregator {
	s.AggregatorStatus = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorType(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorType = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetDescription(v string) *GetAggregatorResponseBodyAggregator {
	s.Description = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetFolderId(v string) *GetAggregatorResponseBodyAggregator {
	s.FolderId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetFolderName(v string) *GetAggregatorResponseBodyAggregator {
	s.FolderName = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetTags(v []*GetAggregatorResponseBodyAggregatorTags) *GetAggregatorResponseBodyAggregator {
	s.Tags = v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) Validate() error {
	if s.AggregatorAccounts != nil {
		for _, item := range s.AggregatorAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregatorResponseBodyAggregatorAggregatorAccounts struct {
	// The ID of the member.
	//
	// example:
	//
	// 171322098523****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The display name of the member.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// Alice
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The resource directory to which the member belongs. Valid value: ResourceDirectory. ResourceDirectory indicates that the member belongs to a resource directory.
	//
	// example:
	//
	// ResourceDirectory
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The status of the configuration recorder for the member. Valid values:
	//
	// 	- REGISTRABLE: The configuration recorder is not registered.
	//
	// 	- BUILDING: The configuration recorder is being deployed.
	//
	// 	- REGISTERED: The configuration recorder is registered.
	//
	// 	- REBUILDING: The configuration recorder is being redeployed.
	//
	// example:
	//
	// REGISTERED
	RecorderStatus *string `json:"RecorderStatus,omitempty" xml:"RecorderStatus,omitempty"`
}

func (s GetAggregatorResponseBodyAggregatorAggregatorAccounts) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorResponseBodyAggregatorAggregatorAccounts) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) GetRecorderStatus() *string {
	return s.RecorderStatus
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetAccountId(v int64) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.AccountId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetAccountName(v string) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.AccountName = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetAccountType(v string) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.AccountType = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetRecorderStatus(v string) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.RecorderStatus = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) Validate() error {
	return dara.Validate(s)
}

type GetAggregatorResponseBodyAggregatorTags struct {
	// The tag key.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetAggregatorResponseBodyAggregatorTags) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorResponseBodyAggregatorTags) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBodyAggregatorTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregatorResponseBodyAggregatorTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregatorResponseBodyAggregatorTags) SetTagKey(v string) *GetAggregatorResponseBodyAggregatorTags {
	s.TagKey = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorTags) SetTagValue(v string) *GetAggregatorResponseBodyAggregatorTags {
	s.TagValue = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorTags) Validate() error {
	return dara.Validate(s)
}
