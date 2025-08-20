// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactSubscriptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListArtifactSubscriptionRuleRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListArtifactSubscriptionRuleRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListArtifactSubscriptionRuleRequest
	GetPageSize() *int32
}

type ListArtifactSubscriptionRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactSubscriptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactSubscriptionRuleRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListArtifactSubscriptionRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactSubscriptionRuleRequest) SetInstanceId(v string) *ListArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleRequest) SetPageNo(v int32) *ListArtifactSubscriptionRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionRuleRequest) SetPageSize(v int32) *ListArtifactSubscriptionRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
