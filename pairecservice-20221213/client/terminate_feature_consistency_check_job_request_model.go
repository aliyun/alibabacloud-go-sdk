// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateFeatureConsistencyCheckJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *TerminateFeatureConsistencyCheckJobRequest
	GetInstanceId() *string
}

type TerminateFeatureConsistencyCheckJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s TerminateFeatureConsistencyCheckJobRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateFeatureConsistencyCheckJobRequest) GoString() string {
	return s.String()
}

func (s *TerminateFeatureConsistencyCheckJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TerminateFeatureConsistencyCheckJobRequest) SetInstanceId(v string) *TerminateFeatureConsistencyCheckJobRequest {
	s.InstanceId = &v
	return s
}

func (s *TerminateFeatureConsistencyCheckJobRequest) Validate() error {
	return dara.Validate(s)
}
