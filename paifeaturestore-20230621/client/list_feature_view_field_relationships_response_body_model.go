// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewFieldRelationshipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRelationships(v []*ListFeatureViewFieldRelationshipsResponseBodyRelationships) *ListFeatureViewFieldRelationshipsResponseBody
	GetRelationships() []*ListFeatureViewFieldRelationshipsResponseBodyRelationships
	SetRequestId(v string) *ListFeatureViewFieldRelationshipsResponseBody
	GetRequestId() *string
}

type ListFeatureViewFieldRelationshipsResponseBody struct {
	Relationships []*ListFeatureViewFieldRelationshipsResponseBodyRelationships `json:"Relationships,omitempty" xml:"Relationships,omitempty" type:"Repeated"`
	// example:
	//
	// BF349686-C932-55B5-9B31-DAFA395C0E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) GetRelationships() []*ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	return s.Relationships
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) SetRelationships(v []*ListFeatureViewFieldRelationshipsResponseBodyRelationships) *ListFeatureViewFieldRelationshipsResponseBody {
	s.Relationships = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) SetRequestId(v string) *ListFeatureViewFieldRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBody) Validate() error {
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

type ListFeatureViewFieldRelationshipsResponseBodyRelationships struct {
	// example:
	//
	// featureView1
	FeatureName *string                                                             `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	Models      []*ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// example:
	//
	// table2
	OfflineTableName *string `json:"OfflineTableName,omitempty" xml:"OfflineTableName,omitempty"`
	// example:
	//
	// table1
	OnlineTableName *string `json:"OnlineTableName,omitempty" xml:"OnlineTableName,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationships) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationships) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) GetModels() []*ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels {
	return s.Models
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) GetOfflineTableName() *string {
	return s.OfflineTableName
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) GetOnlineTableName() *string {
	return s.OnlineTableName
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetFeatureName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetModels(v []*ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.Models = v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetOfflineTableName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.OfflineTableName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) SetOnlineTableName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationships {
	s.OnlineTableName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationships) Validate() error {
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

type ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels struct {
	// example:
	//
	// f1
	FeatureAliasName *string `json:"FeatureAliasName,omitempty" xml:"FeatureAliasName,omitempty"`
	// example:
	//
	// 3
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// dbmtl
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) GoString() string {
	return s.String()
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) GetFeatureAliasName() *string {
	return s.FeatureAliasName
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) GetModelId() *string {
	return s.ModelId
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) GetModelName() *string {
	return s.ModelName
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) SetFeatureAliasName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels {
	s.FeatureAliasName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) SetModelId(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels {
	s.ModelId = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) SetModelName(v string) *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels {
	s.ModelName = &v
	return s
}

func (s *ListFeatureViewFieldRelationshipsResponseBodyRelationshipsModels) Validate() error {
	return dara.Validate(s)
}
