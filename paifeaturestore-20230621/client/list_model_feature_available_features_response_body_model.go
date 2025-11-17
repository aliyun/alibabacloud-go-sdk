// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeatureAvailableFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvaliableFeatures(v []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) *ListModelFeatureAvailableFeaturesResponseBody
	GetAvaliableFeatures() []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures
	SetTotalCount(v int64) *ListModelFeatureAvailableFeaturesResponseBody
	GetTotalCount() *int64
	SetRequestId(v string) *ListModelFeatureAvailableFeaturesResponseBody
	GetRequestId() *string
}

type ListModelFeatureAvailableFeaturesResponseBody struct {
	AvaliableFeatures []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures `json:"AvaliableFeatures,omitempty" xml:"AvaliableFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// ED4DEA2F-F216-57F0-AE28-08D791233280
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) GetAvaliableFeatures() []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	return s.AvaliableFeatures
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) SetAvaliableFeatures(v []*ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) *ListModelFeatureAvailableFeaturesResponseBody {
	s.AvaliableFeatures = v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) SetTotalCount(v int64) *ListModelFeatureAvailableFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) SetRequestId(v string) *ListModelFeatureAvailableFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBody) Validate() error {
	if s.AvaliableFeatures != nil {
		for _, item := range s.AvaliableFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures struct {
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// user_fea
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// example:
	//
	// FeatureView
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) GetName() *string {
	return s.Name
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) GetSourceName() *string {
	return s.SourceName
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) GetSourceType() *string {
	return s.SourceType
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) GetType() *string {
	return s.Type
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetName(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.Name = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetSourceName(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.SourceName = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetSourceType(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.SourceType = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) SetType(v string) *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures {
	s.Type = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponseBodyAvaliableFeatures) Validate() error {
	return dara.Validate(s)
}
