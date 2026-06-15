// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v []*UpdateModelFeatureRequestFeatures) *UpdateModelFeatureRequest
	GetFeatures() []*UpdateModelFeatureRequestFeatures
	SetLabelPriorityLevel(v int64) *UpdateModelFeatureRequest
	GetLabelPriorityLevel() *int64
	SetLabelTableId(v string) *UpdateModelFeatureRequest
	GetLabelTableId() *string
	SetSequenceFeatureViewIds(v []*string) *UpdateModelFeatureRequest
	GetSequenceFeatureViewIds() []*string
}

type UpdateModelFeatureRequest struct {
	// The list of features.
	Features []*UpdateModelFeatureRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// The label table ID. You can call the ListLabelTables operation to query the label table ID.
	//
	// example:
	//
	// 4
	LabelTableId           *string   `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	SequenceFeatureViewIds []*string `json:"SequenceFeatureViewIds,omitempty" xml:"SequenceFeatureViewIds,omitempty" type:"Repeated"`
}

func (s UpdateModelFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureRequest) GetFeatures() []*UpdateModelFeatureRequestFeatures {
	return s.Features
}

func (s *UpdateModelFeatureRequest) GetLabelPriorityLevel() *int64 {
	return s.LabelPriorityLevel
}

func (s *UpdateModelFeatureRequest) GetLabelTableId() *string {
	return s.LabelTableId
}

func (s *UpdateModelFeatureRequest) GetSequenceFeatureViewIds() []*string {
	return s.SequenceFeatureViewIds
}

func (s *UpdateModelFeatureRequest) SetFeatures(v []*UpdateModelFeatureRequestFeatures) *UpdateModelFeatureRequest {
	s.Features = v
	return s
}

func (s *UpdateModelFeatureRequest) SetLabelPriorityLevel(v int64) *UpdateModelFeatureRequest {
	s.LabelPriorityLevel = &v
	return s
}

func (s *UpdateModelFeatureRequest) SetLabelTableId(v string) *UpdateModelFeatureRequest {
	s.LabelTableId = &v
	return s
}

func (s *UpdateModelFeatureRequest) SetSequenceFeatureViewIds(v []*string) *UpdateModelFeatureRequest {
	s.SequenceFeatureViewIds = v
	return s
}

func (s *UpdateModelFeatureRequest) Validate() error {
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

type UpdateModelFeatureRequestFeatures struct {
	// The feature alias.
	//
	// example:
	//
	// sex
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The feature view ID. You can call the ListFeatureViews operation to query the feature view ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// The feature name.
	//
	// This parameter is required.
	//
	// example:
	//
	// gender
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cand_seq__
	PrefixName *string `json:"PrefixName,omitempty" xml:"PrefixName,omitempty"`
	// The feature type. Valid values:
	//
	// - INT32
	//
	// - INT64
	//
	// - FLOAT
	//
	// - DOUBLE
	//
	// - STRING
	//
	// - BOOLEAN
	//
	// - TIMESTAMP.
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateModelFeatureRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureRequestFeatures) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureRequestFeatures) GetAliasName() *string {
	return s.AliasName
}

func (s *UpdateModelFeatureRequestFeatures) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *UpdateModelFeatureRequestFeatures) GetName() *string {
	return s.Name
}

func (s *UpdateModelFeatureRequestFeatures) GetPrefixName() *string {
	return s.PrefixName
}

func (s *UpdateModelFeatureRequestFeatures) GetType() *string {
	return s.Type
}

func (s *UpdateModelFeatureRequestFeatures) SetAliasName(v string) *UpdateModelFeatureRequestFeatures {
	s.AliasName = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetFeatureViewId(v string) *UpdateModelFeatureRequestFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetName(v string) *UpdateModelFeatureRequestFeatures {
	s.Name = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetPrefixName(v string) *UpdateModelFeatureRequestFeatures {
	s.PrefixName = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) SetType(v string) *UpdateModelFeatureRequestFeatures {
	s.Type = &v
	return s
}

func (s *UpdateModelFeatureRequestFeatures) Validate() error {
	return dara.Validate(s)
}
