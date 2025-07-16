// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelFeatures(v []*ListModelFeaturesResponseBodyModelFeatures) *ListModelFeaturesResponseBody
	GetModelFeatures() []*ListModelFeaturesResponseBodyModelFeatures
	SetRequestId(v string) *ListModelFeaturesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListModelFeaturesResponseBody
	GetTotalCount() *int64
}

type ListModelFeaturesResponseBody struct {
	ModelFeatures []*ListModelFeaturesResponseBodyModelFeatures `json:"ModelFeatures,omitempty" xml:"ModelFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// 2CA11923-2A3D-5E5A-8314-E699D2DD15DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesResponseBody) GetModelFeatures() []*ListModelFeaturesResponseBodyModelFeatures {
	return s.ModelFeatures
}

func (s *ListModelFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelFeaturesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelFeaturesResponseBody) SetModelFeatures(v []*ListModelFeaturesResponseBodyModelFeatures) *ListModelFeaturesResponseBody {
	s.ModelFeatures = v
	return s
}

func (s *ListModelFeaturesResponseBody) SetRequestId(v string) *ListModelFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelFeaturesResponseBody) SetTotalCount(v int64) *ListModelFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListModelFeaturesResponseBodyModelFeatures struct {
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T14:46:22.227+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// label_table_1
	LabelTableName *string `json:"LabelTableName,omitempty" xml:"LabelTableName,omitempty"`
	// example:
	//
	// 3
	ModelFeatureId *string `json:"ModelFeatureId,omitempty" xml:"ModelFeatureId,omitempty"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1231243253****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 5
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListModelFeaturesResponseBodyModelFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeaturesResponseBodyModelFeatures) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetLabelTableName() *string {
	return s.LabelTableName
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetModelFeatureId() *string {
	return s.ModelFeatureId
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetName() *string {
	return s.Name
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetOwner() *string {
	return s.Owner
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModelFeaturesResponseBodyModelFeatures) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetGmtCreateTime(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.GmtCreateTime = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetGmtModifiedTime(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetLabelTableName(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.LabelTableName = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetModelFeatureId(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.ModelFeatureId = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetName(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.Name = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetOwner(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.Owner = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetProjectId(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.ProjectId = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) SetProjectName(v string) *ListModelFeaturesResponseBodyModelFeatures {
	s.ProjectName = &v
	return s
}

func (s *ListModelFeaturesResponseBodyModelFeatures) Validate() error {
	return dara.Validate(s)
}
