// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorsResult(v *ListAggregatorsResponseBodyAggregatorsResult) *ListAggregatorsResponseBody
	GetAggregatorsResult() *ListAggregatorsResponseBodyAggregatorsResult
	SetRequestId(v string) *ListAggregatorsResponseBody
	GetRequestId() *string
}

type ListAggregatorsResponseBody struct {
	// The account groups.
	AggregatorsResult *ListAggregatorsResponseBodyAggregatorsResult `json:"AggregatorsResult,omitempty" xml:"AggregatorsResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20C8526D-12C5-4336-BC72-EBD5D1BA732F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBody) GetAggregatorsResult() *ListAggregatorsResponseBodyAggregatorsResult {
	return s.AggregatorsResult
}

func (s *ListAggregatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregatorsResponseBody) SetAggregatorsResult(v *ListAggregatorsResponseBodyAggregatorsResult) *ListAggregatorsResponseBody {
	s.AggregatorsResult = v
	return s
}

func (s *ListAggregatorsResponseBody) SetRequestId(v string) *ListAggregatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregatorsResponseBody) Validate() error {
	if s.AggregatorsResult != nil {
		if err := s.AggregatorsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregatorsResponseBodyAggregatorsResult struct {
	// The list of the account groups.
	Aggregators []*ListAggregatorsResponseBodyAggregatorsResultAggregators `json:"Aggregators,omitempty" xml:"Aggregators,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU2hhcmVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAggregatorsResponseBodyAggregatorsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsResponseBodyAggregatorsResult) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) GetAggregators() []*ListAggregatorsResponseBodyAggregatorsResultAggregators {
	return s.Aggregators
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) SetAggregators(v []*ListAggregatorsResponseBodyAggregatorsResultAggregators) *ListAggregatorsResponseBodyAggregatorsResult {
	s.Aggregators = v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) SetNextToken(v string) *ListAggregatorsResponseBodyAggregatorsResult {
	s.NextToken = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) Validate() error {
	if s.Aggregators != nil {
		for _, item := range s.Aggregators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregatorsResponseBodyAggregatorsResultAggregators struct {
	// The ID of the management account that is used to create the account group.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The number of member accounts in the account group.
	//
	// example:
	//
	// 2
	AggregatorAccountCount *int64 `json:"AggregatorAccountCount,omitempty" xml:"AggregatorAccountCount,omitempty"`
	// The timestamp generated when the account group was created.
	//
	// example:
	//
	// 1623036305000
	AggregatorCreateTimestamp *int64 `json:"AggregatorCreateTimestamp,omitempty" xml:"AggregatorCreateTimestamp,omitempty"`
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
	// Test_Group
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
	// 	- RD: global account group.
	//
	// 	- FOLDER: account group of the folder.
	//
	// 	- CUSTOM: custom account group.
	//
	// example:
	//
	// CUSTOM
	AggregatorType *string `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	// The description of the account group.
	//
	// example:
	//
	// Example-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// r-BU****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// tags
	Tags []*ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAggregatorsResponseBodyAggregatorsResultAggregators) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsResponseBodyAggregatorsResultAggregators) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAggregatorAccountCount() *int64 {
	return s.AggregatorAccountCount
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAggregatorCreateTimestamp() *int64 {
	return s.AggregatorCreateTimestamp
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAggregatorStatus() *int32 {
	return s.AggregatorStatus
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetAggregatorType() *string {
	return s.AggregatorType
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetDescription() *string {
	return s.Description
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) GetTags() []*ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags {
	return s.Tags
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAccountId(v int64) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AccountId = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorAccountCount(v int64) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorAccountCount = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorCreateTimestamp(v int64) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorCreateTimestamp = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorId(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorName(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorName = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorStatus(v int32) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorStatus = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorType(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorType = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetDescription(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.Description = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetFolderId(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.FolderId = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetTags(v []*ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.Tags = v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) Validate() error {
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

type ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags struct {
	// The tag keys of the resource.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag values of the resource.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) SetTagKey(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags {
	s.TagKey = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) SetTagValue(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags {
	s.TagValue = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags) Validate() error {
	return dara.Validate(s)
}
