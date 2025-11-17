// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v []*ListProjectFeaturesResponseBodyFeatures) *ListProjectFeaturesResponseBody
	GetFeatures() []*ListProjectFeaturesResponseBodyFeatures
	SetTotalCount(v int32) *ListProjectFeaturesResponseBody
	GetTotalCount() *int32
	SetRequestId(v string) *ListProjectFeaturesResponseBody
	GetRequestId() *string
}

type ListProjectFeaturesResponseBody struct {
	Features []*ListProjectFeaturesResponseBodyFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7D497816-607C-5B67-97B1-61354B6ACB2B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectFeaturesResponseBody) GetFeatures() []*ListProjectFeaturesResponseBodyFeatures {
	return s.Features
}

func (s *ListProjectFeaturesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectFeaturesResponseBody) SetFeatures(v []*ListProjectFeaturesResponseBodyFeatures) *ListProjectFeaturesResponseBody {
	s.Features = v
	return s
}

func (s *ListProjectFeaturesResponseBody) SetTotalCount(v int32) *ListProjectFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectFeaturesResponseBody) SetRequestId(v string) *ListProjectFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectFeaturesResponseBody) Validate() error {
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

type ListProjectFeaturesResponseBodyFeatures struct {
	// example:
	//
	// age1,age2
	AliasNames *string `json:"AliasNames,omitempty" xml:"AliasNames,omitempty"`
	// example:
	//
	// 1
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// example:
	//
	// fv1
	FeatureViewName *string `json:"FeatureViewName,omitempty" xml:"FeatureViewName,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 10
	ModelFeatureCount *int32 `json:"ModelFeatureCount,omitempty" xml:"ModelFeatureCount,omitempty"`
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectFeaturesResponseBodyFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeaturesResponseBodyFeatures) GoString() string {
	return s.String()
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetAliasNames() *string {
	return s.AliasNames
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetFeatureViewName() *string {
	return s.FeatureViewName
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetModelFeatureCount() *int32 {
	return s.ModelFeatureCount
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetName() *string {
	return s.Name
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectFeaturesResponseBodyFeatures) GetType() *string {
	return s.Type
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetAliasNames(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.AliasNames = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetFeatureViewId(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetFeatureViewName(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.FeatureViewName = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetGmtCreateTime(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.GmtCreateTime = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetModelFeatureCount(v int32) *ListProjectFeaturesResponseBodyFeatures {
	s.ModelFeatureCount = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetName(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.Name = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetOwner(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.Owner = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) SetType(v string) *ListProjectFeaturesResponseBodyFeatures {
	s.Type = &v
	return s
}

func (s *ListProjectFeaturesResponseBodyFeatures) Validate() error {
	return dara.Validate(s)
}
