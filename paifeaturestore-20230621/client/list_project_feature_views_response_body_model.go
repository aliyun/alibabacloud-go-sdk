// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectFeatureViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureViews(v []*ListProjectFeatureViewsResponseBodyFeatureViews) *ListProjectFeatureViewsResponseBody
	GetFeatureViews() []*ListProjectFeatureViewsResponseBodyFeatureViews
	SetRequestId(v string) *ListProjectFeatureViewsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListProjectFeatureViewsResponseBody
	GetTotalCount() *int64
}

type ListProjectFeatureViewsResponseBody struct {
	FeatureViews []*ListProjectFeatureViewsResponseBodyFeatureViews `json:"FeatureViews,omitempty" xml:"FeatureViews,omitempty" type:"Repeated"`
	// example:
	//
	// AE2AF33E-0C0D-51A8-B89B-C5F1DF261D92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectFeatureViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeatureViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponseBody) GetFeatureViews() []*ListProjectFeatureViewsResponseBodyFeatureViews {
	return s.FeatureViews
}

func (s *ListProjectFeatureViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectFeatureViewsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProjectFeatureViewsResponseBody) SetFeatureViews(v []*ListProjectFeatureViewsResponseBodyFeatureViews) *ListProjectFeatureViewsResponseBody {
	s.FeatureViews = v
	return s
}

func (s *ListProjectFeatureViewsResponseBody) SetRequestId(v string) *ListProjectFeatureViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBody) SetTotalCount(v int64) *ListProjectFeatureViewsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBody) Validate() error {
	if s.FeatureViews != nil {
		for _, item := range s.FeatureViews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectFeatureViewsResponseBodyFeatureViews struct {
	// example:
	//
	// 3
	FeatureViewId *string                                                    `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	Features      []*ListProjectFeatureViewsResponseBodyFeatureViewsFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// item_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// feature_view1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// item_id
	ParentJoinId *string `json:"ParentJoinId,omitempty" xml:"ParentJoinId,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectFeatureViewsResponseBodyFeatureViews) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeatureViewsResponseBodyFeatureViews) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) GetFeatures() []*ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	return s.Features
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) GetJoinId() *string {
	return s.JoinId
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) GetName() *string {
	return s.Name
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) GetParentJoinId() *string {
	return s.ParentJoinId
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) GetType() *string {
	return s.Type
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetFeatureViewId(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.FeatureViewId = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetFeatures(v []*ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.Features = v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetJoinId(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.JoinId = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetName(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.Name = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetParentJoinId(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.ParentJoinId = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) SetType(v string) *ListProjectFeatureViewsResponseBodyFeatureViews {
	s.Type = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViews) Validate() error {
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

type ListProjectFeatureViewsResponseBodyFeatureViewsFeatures struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) GetAttributes() []*string {
	return s.Attributes
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) GetName() *string {
	return s.Name
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) GetType() *string {
	return s.Type
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) SetAttributes(v []*string) *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	s.Attributes = v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) SetName(v string) *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	s.Name = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) SetType(v string) *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures {
	s.Type = &v
	return s
}

func (s *ListProjectFeatureViewsResponseBodyFeatureViewsFeatures) Validate() error {
	return dara.Validate(s)
}
