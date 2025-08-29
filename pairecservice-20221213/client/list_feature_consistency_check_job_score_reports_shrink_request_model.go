// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobScoreReportsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeRequestIdsShrink(v string) *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest
	GetExcludeRequestIdsShrink() *string
	SetInstanceId(v string) *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest
	GetInstanceId() *string
}

type ListFeatureConsistencyCheckJobScoreReportsShrinkRequest struct {
	ExcludeRequestIdsShrink *string `json:"ExcludeRequestIds,omitempty" xml:"ExcludeRequestIds,omitempty"`
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) GetExcludeRequestIdsShrink() *string {
	return s.ExcludeRequestIdsShrink
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) SetExcludeRequestIdsShrink(v string) *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest {
	s.ExcludeRequestIdsShrink = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
