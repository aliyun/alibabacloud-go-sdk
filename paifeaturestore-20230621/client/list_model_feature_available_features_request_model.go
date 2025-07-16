// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeatureAvailableFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureName(v string) *ListModelFeatureAvailableFeaturesRequest
	GetFeatureName() *string
}

type ListModelFeatureAvailableFeaturesRequest struct {
	// example:
	//
	// f1
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListModelFeatureAvailableFeaturesRequest) SetFeatureName(v string) *ListModelFeatureAvailableFeaturesRequest {
	s.FeatureName = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
