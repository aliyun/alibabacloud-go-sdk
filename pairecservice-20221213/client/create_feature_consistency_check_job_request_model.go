// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureConsistencyCheckJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *CreateFeatureConsistencyCheckJobRequest
	GetEnvironment() *string
	SetFeatureConsistencyCheckJobConfigId(v string) *CreateFeatureConsistencyCheckJobRequest
	GetFeatureConsistencyCheckJobConfigId() *string
	SetInstanceId(v string) *CreateFeatureConsistencyCheckJobRequest
	GetInstanceId() *string
	SetSamplingDuration(v int32) *CreateFeatureConsistencyCheckJobRequest
	GetSamplingDuration() *int32
}

type CreateFeatureConsistencyCheckJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	FeatureConsistencyCheckJobConfigId *string `json:"FeatureConsistencyCheckJobConfigId,omitempty" xml:"FeatureConsistencyCheckJobConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	SamplingDuration *int32 `json:"SamplingDuration,omitempty" xml:"SamplingDuration,omitempty"`
}

func (s CreateFeatureConsistencyCheckJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureConsistencyCheckJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureConsistencyCheckJobRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CreateFeatureConsistencyCheckJobRequest) GetFeatureConsistencyCheckJobConfigId() *string {
	return s.FeatureConsistencyCheckJobConfigId
}

func (s *CreateFeatureConsistencyCheckJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFeatureConsistencyCheckJobRequest) GetSamplingDuration() *int32 {
	return s.SamplingDuration
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetEnvironment(v string) *CreateFeatureConsistencyCheckJobRequest {
	s.Environment = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetFeatureConsistencyCheckJobConfigId(v string) *CreateFeatureConsistencyCheckJobRequest {
	s.FeatureConsistencyCheckJobConfigId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetInstanceId(v string) *CreateFeatureConsistencyCheckJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) SetSamplingDuration(v int32) *CreateFeatureConsistencyCheckJobRequest {
	s.SamplingDuration = &v
	return s
}

func (s *CreateFeatureConsistencyCheckJobRequest) Validate() error {
	return dara.Validate(s)
}
