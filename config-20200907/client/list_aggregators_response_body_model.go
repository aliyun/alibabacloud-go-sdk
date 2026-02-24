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
	AggregatorsResult *ListAggregatorsResponseBodyAggregatorsResult `json:"AggregatorsResult,omitempty" xml:"AggregatorsResult,omitempty" type:"Struct"`
	RequestId         *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Aggregators []*ListAggregatorsResponseBodyAggregatorsResultAggregators `json:"Aggregators,omitempty" xml:"Aggregators,omitempty" type:"Repeated"`
	NextToken   *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	AccountId                 *int64                                                         `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorAccountCount    *int64                                                         `json:"AggregatorAccountCount,omitempty" xml:"AggregatorAccountCount,omitempty"`
	AggregatorCreateTimestamp *int64                                                         `json:"AggregatorCreateTimestamp,omitempty" xml:"AggregatorCreateTimestamp,omitempty"`
	AggregatorId              *string                                                        `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName            *string                                                        `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	AggregatorStatus          *int32                                                         `json:"AggregatorStatus,omitempty" xml:"AggregatorStatus,omitempty"`
	AggregatorType            *string                                                        `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	Description               *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	FolderId                  *string                                                        `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Tags                      []*ListAggregatorsResponseBodyAggregatorsResultAggregatorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
