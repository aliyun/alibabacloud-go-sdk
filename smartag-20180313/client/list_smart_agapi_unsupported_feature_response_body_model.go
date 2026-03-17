// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartAGApiUnsupportedFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v []*ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) *ListSmartAGApiUnsupportedFeatureResponseBody
	GetFeatures() []*ListSmartAGApiUnsupportedFeatureResponseBodyFeatures
	SetRequestId(v string) *ListSmartAGApiUnsupportedFeatureResponseBody
	GetRequestId() *string
}

type ListSmartAGApiUnsupportedFeatureResponseBody struct {
	// A list of unsupported features.
	Features []*ListSmartAGApiUnsupportedFeatureResponseBodyFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7459545D-2F0D-43E6-9957-CB7E0223332B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSmartAGApiUnsupportedFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGApiUnsupportedFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBody) GetFeatures() []*ListSmartAGApiUnsupportedFeatureResponseBodyFeatures {
	return s.Features
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBody) SetFeatures(v []*ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) *ListSmartAGApiUnsupportedFeatureResponseBody {
	s.Features = v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBody) SetRequestId(v string) *ListSmartAGApiUnsupportedFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBody) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmartAGApiUnsupportedFeatureResponseBodyFeatures struct {
	// The unsupported feature.
	//
	// For more information about the description of each feature, see the related API reference.
	//
	// example:
	//
	// ISP
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
}

func (s ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) GoString() string {
	return s.String()
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) GetFeature() *string {
	return s.Feature
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) SetFeature(v string) *ListSmartAGApiUnsupportedFeatureResponseBodyFeatures {
	s.Feature = &v
	return s
}

func (s *ListSmartAGApiUnsupportedFeatureResponseBodyFeatures) Validate() error {
	return dara.Validate(s)
}
