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
	// This parameter is required.
	Features []*CreateModelFeatureRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LabelPriorityLevel *int64 `json:"LabelPriorityLevel,omitempty" xml:"LabelPriorityLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// model_feature_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId              *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
	// example:
	//
	// userid
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *CreateModelFeatureRequestFeatures) SetType(v string) *CreateModelFeatureRequestFeatures {
	s.Type = &v
	return s
}

func (s *CreateModelFeatureRequestFeatures) Validate() error {
	return dara.Validate(s)
}
