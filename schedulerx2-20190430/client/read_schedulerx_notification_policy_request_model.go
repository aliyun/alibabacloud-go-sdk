// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxNotificationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ReadSchedulerxNotificationPolicyRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadSchedulerxNotificationPolicyRequest
	GetNextToken() *string
	SetPolicyName(v string) *ReadSchedulerxNotificationPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *ReadSchedulerxNotificationPolicyRequest
	GetRegionId() *string
}

type ReadSchedulerxNotificationPolicyRequest struct {
	// The maximum number of entries returned. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The cursor for pagination. Leave this parameter empty for the first request. When the returned value is empty, all data has been retrieved.
	//
	// example:
	//
	// O39nXKu5XafATl3/cJjSJw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the notification policy. Supports fuzzy matching.
	//
	// example:
	//
	// test-weekdays
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadSchedulerxNotificationPolicyRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadSchedulerxNotificationPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ReadSchedulerxNotificationPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReadSchedulerxNotificationPolicyRequest) SetMaxResults(v int32) *ReadSchedulerxNotificationPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyRequest) SetNextToken(v string) *ReadSchedulerxNotificationPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyRequest) SetPolicyName(v string) *ReadSchedulerxNotificationPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyRequest) SetRegionId(v string) *ReadSchedulerxNotificationPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
