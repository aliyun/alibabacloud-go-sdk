// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewRelationshipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRelationships(v []*ListFeatureViewRelationshipsResponseBodyRelationships) *ListFeatureViewRelationshipsResponseBody
	GetRelationships() []*ListFeatureViewRelationshipsResponseBodyRelationships
	SetRequestId(v string) *ListFeatureViewRelationshipsResponseBody
	GetRequestId() *string
}

type ListFeatureViewRelationshipsResponseBody struct {
	Relationships []*ListFeatureViewRelationshipsResponseBodyRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
	// example:
	//
	// 0FBBE454-9BD1-5D8F-9129-D14DB7FAFE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureViewRelationshipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponseBody) GetRelationships() []*ListFeatureViewRelationshipsResponseBodyRelationships {
	return s.Relationships
}

func (s *ListFeatureViewRelationshipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureViewRelationshipsResponseBody) SetRelationships(v []*ListFeatureViewRelationshipsResponseBodyRelationships) *ListFeatureViewRelationshipsResponseBody {
	s.Relationships = v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBody) SetRequestId(v string) *ListFeatureViewRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBody) Validate() error {
	if s.Relationships != nil {
		for _, item := range s.Relationships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeatureViewRelationshipsResponseBodyRelationships struct {
	// example:
	//
	// fv1
	FeatureViewName *string                                                        `json:"FeatureViewName,omitempty" xml:"FeatureViewName,omitempty"`
	Models          []*ListFeatureViewRelationshipsResponseBodyRelationshipsModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListFeatureViewRelationshipsResponseBodyRelationships) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponseBodyRelationships) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) GetFeatureViewName() *string {
	return s.FeatureViewName
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) GetModels() []*ListFeatureViewRelationshipsResponseBodyRelationshipsModels {
	return s.Models
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) SetFeatureViewName(v string) *ListFeatureViewRelationshipsResponseBodyRelationships {
	s.FeatureViewName = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) SetModels(v []*ListFeatureViewRelationshipsResponseBodyRelationshipsModels) *ListFeatureViewRelationshipsResponseBodyRelationships {
	s.Models = v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) SetProjectName(v string) *ListFeatureViewRelationshipsResponseBodyRelationships {
	s.ProjectName = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationships) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeatureViewRelationshipsResponseBodyRelationshipsModels struct {
	// example:
	//
	// 3
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// dbmtl
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s ListFeatureViewRelationshipsResponseBodyRelationshipsModels) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewRelationshipsResponseBodyRelationshipsModels) GoString() string {
	return s.String()
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) GetModelId() *string {
	return s.ModelId
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) GetModelName() *string {
	return s.ModelName
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) SetModelId(v string) *ListFeatureViewRelationshipsResponseBodyRelationshipsModels {
	s.ModelId = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) SetModelName(v string) *ListFeatureViewRelationshipsResponseBodyRelationshipsModels {
	s.ModelName = &v
	return s
}

func (s *ListFeatureViewRelationshipsResponseBodyRelationshipsModels) Validate() error {
	return dara.Validate(s)
}
