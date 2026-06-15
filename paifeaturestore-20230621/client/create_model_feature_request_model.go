// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatures(v []*CreateModelFeatureRequestFeatures) *CreateModelFeatureRequest
	GetFeatures() []*CreateModelFeatureRequestFeatures
	SetLabelPriorityLevel(v int64) *CreateModelFeatureRequest
	GetLabelPriorityLevel() *int64
	SetLabelTableId(v string) *CreateModelFeatureRequest
	GetLabelTableId() *string
	SetName(v string) *CreateModelFeatureRequest
	GetName() *string
	SetProjectId(v string) *CreateModelFeatureRequest
	GetProjectId() *string
	SetSequenceFeatureViewIds(v []*string) *CreateModelFeatureRequest
	GetSequenceFeatureViewIds() []*string
}

type CreateModelFeatureRequest struct {
	// The list of features.
	//
	// This parameter is required.
	Features []*CreateModelFeatureRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The priority level of the label table. Default value: 0, which indicates that conflicts between label table features and feature view features are not allowed. A value of 1 indicates that the label table takes precedence when conflicts occur. A value of 2 indicates that the feature view takes precedence.
	//
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// The label table ID. You can call the ListLabelTables operation to obtain the label table ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// The name of the model feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// model_feature_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project ID. You can call the ListProjects operation to obtain the project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sequence feature view IDs.
	SequenceFeatureViewIds []*string `json:"SequenceFeatureViewIds,omitempty" xml:"SequenceFeatureViewIds,omitempty" type:"Repeated"`
}

func (s CreateModelFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelFeatureRequest) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureRequest) GetFeatures() []*CreateModelFeatureRequestFeatures {
	return s.Features
}

func (s *CreateModelFeatureRequest) GetLabelPriorityLevel() *int64 {
	return s.LabelPriorityLevel
}

func (s *CreateModelFeatureRequest) GetLabelTableId() *string {
	return s.LabelTableId
}

func (s *CreateModelFeatureRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelFeatureRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateModelFeatureRequest) GetSequenceFeatureViewIds() []*string {
	return s.SequenceFeatureViewIds
}

func (s *CreateModelFeatureRequest) SetFeatures(v []*CreateModelFeatureRequestFeatures) *CreateModelFeatureRequest {
	s.Features = v
	return s
}

func (s *CreateModelFeatureRequest) SetLabelPriorityLevel(v int64) *CreateModelFeatureRequest {
	s.LabelPriorityLevel = &v
	return s
}

func (s *CreateModelFeatureRequest) SetLabelTableId(v string) *CreateModelFeatureRequest {
	s.LabelTableId = &v
	return s
}

func (s *CreateModelFeatureRequest) SetName(v string) *CreateModelFeatureRequest {
	s.Name = &v
	return s
}

func (s *CreateModelFeatureRequest) SetProjectId(v string) *CreateModelFeatureRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateModelFeatureRequest) SetSequenceFeatureViewIds(v []*string) *CreateModelFeatureRequest {
	s.SequenceFeatureViewIds = v
	return s
}

func (s *CreateModelFeatureRequest) Validate() error {
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

type CreateModelFeatureRequestFeatures struct {
	// The alias of the feature.
	//
	// example:
	//
	// userid
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The feature view ID. You can call the ListFeatureViews operation to obtain the feature view ID.
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
	// user_id
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

func (s CreateModelFeatureRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s CreateModelFeatureRequestFeatures) GoString() string {
	return s.String()
}

func (s *CreateModelFeatureRequestFeatures) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateModelFeatureRequestFeatures) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *CreateModelFeatureRequestFeatures) GetName() *string {
	return s.Name
}

func (s *CreateModelFeatureRequestFeatures) GetPrefixName() *string {
	return s.PrefixName
}

func (s *CreateModelFeatureRequestFeatures) GetType() *string {
	return s.Type
}

func (s *CreateModelFeatureRequestFeatures) SetAliasName(v string) *CreateModelFeatureRequestFeatures {
	s.AliasName = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetFeatureViewId(v string) *CreateModelFeatureRequestFeatures {
	s.FeatureViewId = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetName(v string) *CreateModelFeatureRequestFeatures {
	s.Name = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetPrefixName(v string) *CreateModelFeatureRequestFeatures {
	s.PrefixName = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) SetType(v string) *CreateModelFeatureRequestFeatures {
	s.Type = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) Validate() error {
	return dara.Validate(s)
}
