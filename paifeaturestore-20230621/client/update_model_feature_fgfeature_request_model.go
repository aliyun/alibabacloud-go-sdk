// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelFeatureFGFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLookupFeatures(v []*UpdateModelFeatureFGFeatureRequestLookupFeatures) *UpdateModelFeatureFGFeatureRequest
	GetLookupFeatures() []*UpdateModelFeatureFGFeatureRequestLookupFeatures
	SetRawFeatures(v []*UpdateModelFeatureFGFeatureRequestRawFeatures) *UpdateModelFeatureFGFeatureRequest
	GetRawFeatures() []*UpdateModelFeatureFGFeatureRequestRawFeatures
	SetReserves(v []*string) *UpdateModelFeatureFGFeatureRequest
	GetReserves() []*string
	SetSequenceFeatures(v []*UpdateModelFeatureFGFeatureRequestSequenceFeatures) *UpdateModelFeatureFGFeatureRequest
	GetSequenceFeatures() []*UpdateModelFeatureFGFeatureRequestSequenceFeatures
}

type UpdateModelFeatureFGFeatureRequest struct {
	LookupFeatures   []*UpdateModelFeatureFGFeatureRequestLookupFeatures   `json:"LookupFeatures,omitempty" xml:"LookupFeatures,omitempty" type:"Repeated"`
	RawFeatures      []*UpdateModelFeatureFGFeatureRequestRawFeatures      `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty" type:"Repeated"`
	Reserves         []*string                                             `json:"Reserves,omitempty" xml:"Reserves,omitempty" type:"Repeated"`
	SequenceFeatures []*UpdateModelFeatureFGFeatureRequestSequenceFeatures `json:"SequenceFeatures,omitempty" xml:"SequenceFeatures,omitempty" type:"Repeated"`
}

func (s UpdateModelFeatureFGFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequest) GetLookupFeatures() []*UpdateModelFeatureFGFeatureRequestLookupFeatures {
	return s.LookupFeatures
}

func (s *UpdateModelFeatureFGFeatureRequest) GetRawFeatures() []*UpdateModelFeatureFGFeatureRequestRawFeatures {
	return s.RawFeatures
}

func (s *UpdateModelFeatureFGFeatureRequest) GetReserves() []*string {
	return s.Reserves
}

func (s *UpdateModelFeatureFGFeatureRequest) GetSequenceFeatures() []*UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	return s.SequenceFeatures
}

func (s *UpdateModelFeatureFGFeatureRequest) SetLookupFeatures(v []*UpdateModelFeatureFGFeatureRequestLookupFeatures) *UpdateModelFeatureFGFeatureRequest {
	s.LookupFeatures = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) SetRawFeatures(v []*UpdateModelFeatureFGFeatureRequestRawFeatures) *UpdateModelFeatureFGFeatureRequest {
	s.RawFeatures = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) SetReserves(v []*string) *UpdateModelFeatureFGFeatureRequest {
	s.Reserves = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) SetSequenceFeatures(v []*UpdateModelFeatureFGFeatureRequestSequenceFeatures) *UpdateModelFeatureFGFeatureRequest {
	s.SequenceFeatures = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequest) Validate() error {
	if s.LookupFeatures != nil {
		for _, item := range s.LookupFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RawFeatures != nil {
		for _, item := range s.RawFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SequenceFeatures != nil {
		for _, item := range s.SequenceFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateModelFeatureFGFeatureRequestLookupFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Item
	KeyFeatureDomain *string `json:"KeyFeatureDomain,omitempty" xml:"KeyFeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	KeyFeatureName *string `json:"KeyFeatureName,omitempty" xml:"KeyFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// User
	MapFeatureDomain *string `json:"MapFeatureDomain,omitempty" xml:"MapFeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	MapFeatureName *string `json:"MapFeatureName,omitempty" xml:"MapFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateModelFeatureFGFeatureRequestLookupFeatures) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestLookupFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetKeyFeatureDomain() *string {
	return s.KeyFeatureDomain
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetKeyFeatureName() *string {
	return s.KeyFeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetMapFeatureDomain() *string {
	return s.MapFeatureDomain
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetMapFeatureName() *string {
	return s.MapFeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetDefaultValue(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.DefaultValue = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetKeyFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.KeyFeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetKeyFeatureName(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.KeyFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetMapFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.MapFeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetMapFeatureName(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.MapFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) SetValueType(v string) *UpdateModelFeatureFGFeatureRequestLookupFeatures {
	s.ValueType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestLookupFeatures) Validate() error {
	return dara.Validate(s)
}

type UpdateModelFeatureFGFeatureRequestRawFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IdFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateModelFeatureFGFeatureRequestRawFeatures) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestRawFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) GetFeatureDomain() *string {
	return s.FeatureDomain
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) GetFeatureType() *string {
	return s.FeatureType
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) GetInputFeatureName() *string {
	return s.InputFeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetDefaultValue(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.DefaultValue = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetFeatureType(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.FeatureType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetInputFeatureName(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) SetValueType(v string) *UpdateModelFeatureFGFeatureRequestRawFeatures {
	s.ValueType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestRawFeatures) Validate() error {
	return dara.Validate(s)
}

type UpdateModelFeatureFGFeatureRequestSequenceFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// #
	AttributeDelim *string `json:"AttributeDelim,omitempty" xml:"AttributeDelim,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ;
	SequenceDelim *string `json:"SequenceDelim,omitempty" xml:"SequenceDelim,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	SequenceLength *int64                                                           `json:"SequenceLength,omitempty" xml:"SequenceLength,omitempty"`
	SubFeatures    []*UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures `json:"SubFeatures,omitempty" xml:"SubFeatures,omitempty" type:"Repeated"`
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeatures) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) GetAttributeDelim() *string {
	return s.AttributeDelim
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) GetSequenceDelim() *string {
	return s.SequenceDelim
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) GetSequenceLength() *int64 {
	return s.SequenceLength
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) GetSubFeatures() []*UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	return s.SubFeatures
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetAttributeDelim(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.AttributeDelim = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetSequenceDelim(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.SequenceDelim = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetSequenceLength(v int64) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.SequenceLength = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) SetSubFeatures(v []*UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) *UpdateModelFeatureFGFeatureRequestSequenceFeatures {
	s.SubFeatures = v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeatures) Validate() error {
	if s.SubFeatures != nil {
		for _, item := range s.SubFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures struct {
	// This parameter is required.
	//
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RawFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GetFeatureDomain() *string {
	return s.FeatureDomain
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GetFeatureType() *string {
	return s.FeatureType
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GetInputFeatureName() *string {
	return s.InputFeatureName
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetDefaultValue(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.DefaultValue = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetFeatureDomain(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetFeatureName(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.FeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetFeatureType(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.FeatureType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetInputFeatureName(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) SetValueType(v string) *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures {
	s.ValueType = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureRequestSequenceFeaturesSubFeatures) Validate() error {
	return dara.Validate(s)
}
