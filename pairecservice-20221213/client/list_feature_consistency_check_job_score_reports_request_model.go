// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobScoreReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeRequestIds(v []*string) *ListFeatureConsistencyCheckJobScoreReportsRequest
	GetExcludeRequestIds() []*string
	SetInstanceId(v string) *ListFeatureConsistencyCheckJobScoreReportsRequest
	GetInstanceId() *string
}

type ListFeatureConsistencyCheckJobScoreReportsRequest struct {
	ExcludeRequestIds []*string `json:"ExcludeRequestIds,omitempty" xml:"ExcludeRequestIds,omitempty" type:"Repeated"`
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) GetExcludeRequestIds() []*string {
	return s.ExcludeRequestIds
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) SetExcludeRequestIds(v []*string) *ListFeatureConsistencyCheckJobScoreReportsRequest {
	s.ExcludeRequestIds = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobScoreReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsRequest) Validate() error {
	return dara.Validate(s)
}
