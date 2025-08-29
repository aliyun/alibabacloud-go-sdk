// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySampleConsistencyJobDifferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureName(v string) *QuerySampleConsistencyJobDifferenceRequest
	GetFeatureName() *string
	SetFeatureType(v string) *QuerySampleConsistencyJobDifferenceRequest
	GetFeatureType() *string
	SetInstanceId(v string) *QuerySampleConsistencyJobDifferenceRequest
	GetInstanceId() *string
}

type QuerySampleConsistencyJobDifferenceRequest struct {
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QuerySampleConsistencyJobDifferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySampleConsistencyJobDifferenceRequest) GoString() string {
	return s.String()
}

func (s *QuerySampleConsistencyJobDifferenceRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *QuerySampleConsistencyJobDifferenceRequest) GetFeatureType() *string {
	return s.FeatureType
}

func (s *QuerySampleConsistencyJobDifferenceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySampleConsistencyJobDifferenceRequest) SetFeatureName(v string) *QuerySampleConsistencyJobDifferenceRequest {
	s.FeatureName = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceRequest) SetFeatureType(v string) *QuerySampleConsistencyJobDifferenceRequest {
	s.FeatureType = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceRequest) SetInstanceId(v string) *QuerySampleConsistencyJobDifferenceRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySampleConsistencyJobDifferenceRequest) Validate() error {
	return dara.Validate(s)
}
