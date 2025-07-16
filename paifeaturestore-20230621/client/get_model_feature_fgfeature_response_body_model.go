// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelFeatureFGFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLookupFeatures(v []*GetModelFeatureFGFeatureResponseBodyLookupFeatures) *GetModelFeatureFGFeatureResponseBody
	GetLookupFeatures() []*GetModelFeatureFGFeatureResponseBodyLookupFeatures
	SetRawFeatures(v []*GetModelFeatureFGFeatureResponseBodyRawFeatures) *GetModelFeatureFGFeatureResponseBody
	GetRawFeatures() []*GetModelFeatureFGFeatureResponseBodyRawFeatures
	SetRequestId(v string) *GetModelFeatureFGFeatureResponseBody
	GetRequestId() *string
	SetReserves(v []*string) *GetModelFeatureFGFeatureResponseBody
	GetReserves() []*string
	SetSequenceFeatures(v []*GetModelFeatureFGFeatureResponseBodySequenceFeatures) *GetModelFeatureFGFeatureResponseBody
	GetSequenceFeatures() []*GetModelFeatureFGFeatureResponseBodySequenceFeatures
}

type GetModelFeatureFGFeatureResponseBody struct {
	LookupFeatures []*GetModelFeatureFGFeatureResponseBodyLookupFeatures `json:"LookupFeatures,omitempty" xml:"LookupFeatures,omitempty" type:"Repeated"`
	RawFeatures    []*GetModelFeatureFGFeatureResponseBodyRawFeatures    `json:"RawFeatures,omitempty" xml:"RawFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// E23EFF09-58AA-5420-934F-8453AE01548D
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Reserves         []*string                                               `json:"Reserves,omitempty" xml:"Reserves,omitempty" type:"Repeated"`
	SequenceFeatures []*GetModelFeatureFGFeatureResponseBodySequenceFeatures `json:"SequenceFeatures,omitempty" xml:"SequenceFeatures,omitempty" type:"Repeated"`
}

func (s GetModelFeatureFGFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBody) GetLookupFeatures() []*GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	return s.LookupFeatures
}

func (s *GetModelFeatureFGFeatureResponseBody) GetRawFeatures() []*GetModelFeatureFGFeatureResponseBodyRawFeatures {
	return s.RawFeatures
}

func (s *GetModelFeatureFGFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelFeatureFGFeatureResponseBody) GetReserves() []*string {
	return s.Reserves
}

func (s *GetModelFeatureFGFeatureResponseBody) GetSequenceFeatures() []*GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	return s.SequenceFeatures
}

func (s *GetModelFeatureFGFeatureResponseBody) SetLookupFeatures(v []*GetModelFeatureFGFeatureResponseBodyLookupFeatures) *GetModelFeatureFGFeatureResponseBody {
	s.LookupFeatures = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetRawFeatures(v []*GetModelFeatureFGFeatureResponseBodyRawFeatures) *GetModelFeatureFGFeatureResponseBody {
	s.RawFeatures = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetRequestId(v string) *GetModelFeatureFGFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetReserves(v []*string) *GetModelFeatureFGFeatureResponseBody {
	s.Reserves = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) SetSequenceFeatures(v []*GetModelFeatureFGFeatureResponseBodySequenceFeatures) *GetModelFeatureFGFeatureResponseBody {
	s.SequenceFeatures = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetModelFeatureFGFeatureResponseBodyLookupFeatures struct {
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// Item
	KeyFeatureDomain *string `json:"KeyFeatureDomain,omitempty" xml:"KeyFeatureDomain,omitempty"`
	// example:
	//
	// 1
	KeyFeatureName *string `json:"KeyFeatureName,omitempty" xml:"KeyFeatureName,omitempty"`
	// example:
	//
	// User
	MapFeatureDomain *string `json:"MapFeatureDomain,omitempty" xml:"MapFeatureDomain,omitempty"`
	// example:
	//
	// item_id
	MapFeatureName *string `json:"MapFeatureName,omitempty" xml:"MapFeatureName,omitempty"`
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetModelFeatureFGFeatureResponseBodyLookupFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodyLookupFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetKeyFeatureDomain() *string {
	return s.KeyFeatureDomain
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetKeyFeatureName() *string {
	return s.KeyFeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetMapFeatureDomain() *string {
	return s.MapFeatureDomain
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetMapFeatureName() *string {
	return s.MapFeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) GetValueType() *string {
	return s.ValueType
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetDefaultValue(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.DefaultValue = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetKeyFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.KeyFeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetKeyFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.KeyFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetMapFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.MapFeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetMapFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.MapFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) SetValueType(v string) *GetModelFeatureFGFeatureResponseBodyLookupFeatures {
	s.ValueType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyLookupFeatures) Validate() error {
	return dara.Validate(s)
}

type GetModelFeatureFGFeatureResponseBodyRawFeatures struct {
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// IdFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetModelFeatureFGFeatureResponseBodyRawFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodyRawFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) GetFeatureDomain() *string {
	return s.FeatureDomain
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) GetFeatureType() *string {
	return s.FeatureType
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) GetInputFeatureName() *string {
	return s.InputFeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) GetValueType() *string {
	return s.ValueType
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetDefaultValue(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.DefaultValue = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetFeatureType(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.FeatureType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetInputFeatureName(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) SetValueType(v string) *GetModelFeatureFGFeatureResponseBodyRawFeatures {
	s.ValueType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodyRawFeatures) Validate() error {
	return dara.Validate(s)
}

type GetModelFeatureFGFeatureResponseBodySequenceFeatures struct {
	// example:
	//
	// #
	AttributeDelim *string `json:"AttributeDelim,omitempty" xml:"AttributeDelim,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// ;
	SequenceDelim *string `json:"SequenceDelim,omitempty" xml:"SequenceDelim,omitempty"`
	// example:
	//
	// 50
	SequenceLength *int64                                                             `json:"SequenceLength,omitempty" xml:"SequenceLength,omitempty"`
	SubFeatures    []*GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures `json:"SubFeatures,omitempty" xml:"SubFeatures,omitempty" type:"Repeated"`
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) GetAttributeDelim() *string {
	return s.AttributeDelim
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) GetSequenceDelim() *string {
	return s.SequenceDelim
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) GetSequenceLength() *int64 {
	return s.SequenceLength
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) GetSubFeatures() []*GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	return s.SubFeatures
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetAttributeDelim(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.AttributeDelim = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetSequenceDelim(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.SequenceDelim = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetSequenceLength(v int64) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.SequenceLength = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) SetSubFeatures(v []*GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) *GetModelFeatureFGFeatureResponseBodySequenceFeatures {
	s.SubFeatures = v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeatures) Validate() error {
	return dara.Validate(s)
}

type GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures struct {
	// example:
	//
	// -1024
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// User
	FeatureDomain *string `json:"FeatureDomain,omitempty" xml:"FeatureDomain,omitempty"`
	// example:
	//
	// item_id
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// IdFeature
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// item_id
	InputFeatureName *string `json:"InputFeatureName,omitempty" xml:"InputFeatureName,omitempty"`
	// example:
	//
	// STRING
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GetFeatureDomain() *string {
	return s.FeatureDomain
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GetFeatureName() *string {
	return s.FeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GetFeatureType() *string {
	return s.FeatureType
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GetInputFeatureName() *string {
	return s.InputFeatureName
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) GetValueType() *string {
	return s.ValueType
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetDefaultValue(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.DefaultValue = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetFeatureDomain(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.FeatureDomain = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetFeatureName(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.FeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetFeatureType(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.FeatureType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetInputFeatureName(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.InputFeatureName = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) SetValueType(v string) *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures {
	s.ValueType = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponseBodySequenceFeaturesSubFeatures) Validate() error {
	return dara.Validate(s)
}
