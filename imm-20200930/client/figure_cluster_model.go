// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFigureCluster interface {
	dara.Model
	String() string
	GoString() string
	SetAverageAge(v float32) *FigureCluster
	GetAverageAge() *float32
	SetCover(v *File) *FigureCluster
	GetCover() *File
	SetCreateTime(v string) *FigureCluster
	GetCreateTime() *string
	SetCustomId(v string) *FigureCluster
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *FigureCluster
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *FigureCluster
	GetDatasetName() *string
	SetFaceCount(v int64) *FigureCluster
	GetFaceCount() *int64
	SetGender(v string) *FigureCluster
	GetGender() *string
	SetImageCount(v int64) *FigureCluster
	GetImageCount() *int64
	SetMaxAge(v float32) *FigureCluster
	GetMaxAge() *float32
	SetMetaLockVersion(v int64) *FigureCluster
	GetMetaLockVersion() *int64
	SetMinAge(v float32) *FigureCluster
	GetMinAge() *float32
	SetName(v string) *FigureCluster
	GetName() *string
	SetObjectId(v string) *FigureCluster
	GetObjectId() *string
	SetObjectType(v string) *FigureCluster
	GetObjectType() *string
	SetOwnerId(v string) *FigureCluster
	GetOwnerId() *string
	SetProjectName(v string) *FigureCluster
	GetProjectName() *string
	SetUpdateTime(v string) *FigureCluster
	GetUpdateTime() *string
	SetVideoCount(v int64) *FigureCluster
	GetVideoCount() *int64
}

type FigureCluster struct {
	// The average age.
	//
	// example:
	//
	// 26
	AverageAge *float32 `json:"AverageAge,omitempty" xml:"AverageAge,omitempty"`
	// The cover image.
	Cover *File `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-14T10:10:52.83948013+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom ID.
	//
	// example:
	//
	// abc
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels. You can search for clusters by label.
	//
	// example:
	//
	// {"Bucket": "examplebucket"}
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The number of faces.
	//
	// example:
	//
	// 3
	FaceCount *int64 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// The gender.
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The number of images.
	//
	// example:
	//
	// 5
	ImageCount *int64 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// The maximum age.
	//
	// example:
	//
	// 44
	MaxAge *float32 `json:"MaxAge,omitempty" xml:"MaxAge,omitempty"`
	// The version of the metadata lock. A metadata lock version can be obtained by using a get or list operation. If you include the MetaLockVersion parameter in a request to update the cluster, the server checks consistency between the MetaLockVersion parameter value sent in the request and the one on the server side and updates the cluster only when they are consistent. This parameter prevents update conflicts in concurrent scenarios. The initial version is 0. The version is automatically increased by 1 after each successful update.
	//
	// example:
	//
	// 0
	MetaLockVersion *int64 `json:"MetaLockVersion,omitempty" xml:"MetaLockVersion,omitempty"`
	// The minimum age.
	//
	// example:
	//
	// 12
	MinAge *float32 `json:"MinAge,omitempty" xml:"MinAge,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// Cluster-ae6e3472-999e-410b-b54e-cd5dba****
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The type of the cluster.
	//
	// example:
	//
	// figure-cluster
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 102321002****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-01-14T10:10:52.83948013+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The number of videos.
	//
	// example:
	//
	// 3
	VideoCount *int64 `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
}

func (s FigureCluster) String() string {
	return dara.Prettify(s)
}

func (s FigureCluster) GoString() string {
	return s.String()
}

func (s *FigureCluster) GetAverageAge() *float32 {
	return s.AverageAge
}

func (s *FigureCluster) GetCover() *File {
	return s.Cover
}

func (s *FigureCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *FigureCluster) GetCustomId() *string {
	return s.CustomId
}

func (s *FigureCluster) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *FigureCluster) GetDatasetName() *string {
	return s.DatasetName
}

func (s *FigureCluster) GetFaceCount() *int64 {
	return s.FaceCount
}

func (s *FigureCluster) GetGender() *string {
	return s.Gender
}

func (s *FigureCluster) GetImageCount() *int64 {
	return s.ImageCount
}

func (s *FigureCluster) GetMaxAge() *float32 {
	return s.MaxAge
}

func (s *FigureCluster) GetMetaLockVersion() *int64 {
	return s.MetaLockVersion
}

func (s *FigureCluster) GetMinAge() *float32 {
	return s.MinAge
}

func (s *FigureCluster) GetName() *string {
	return s.Name
}

func (s *FigureCluster) GetObjectId() *string {
	return s.ObjectId
}

func (s *FigureCluster) GetObjectType() *string {
	return s.ObjectType
}

func (s *FigureCluster) GetOwnerId() *string {
	return s.OwnerId
}

func (s *FigureCluster) GetProjectName() *string {
	return s.ProjectName
}

func (s *FigureCluster) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *FigureCluster) GetVideoCount() *int64 {
	return s.VideoCount
}

func (s *FigureCluster) SetAverageAge(v float32) *FigureCluster {
	s.AverageAge = &v
	return s
}

func (s *FigureCluster) SetCover(v *File) *FigureCluster {
	s.Cover = v
	return s
}

func (s *FigureCluster) SetCreateTime(v string) *FigureCluster {
	s.CreateTime = &v
	return s
}

func (s *FigureCluster) SetCustomId(v string) *FigureCluster {
	s.CustomId = &v
	return s
}

func (s *FigureCluster) SetCustomLabels(v map[string]interface{}) *FigureCluster {
	s.CustomLabels = v
	return s
}

func (s *FigureCluster) SetDatasetName(v string) *FigureCluster {
	s.DatasetName = &v
	return s
}

func (s *FigureCluster) SetFaceCount(v int64) *FigureCluster {
	s.FaceCount = &v
	return s
}

func (s *FigureCluster) SetGender(v string) *FigureCluster {
	s.Gender = &v
	return s
}

func (s *FigureCluster) SetImageCount(v int64) *FigureCluster {
	s.ImageCount = &v
	return s
}

func (s *FigureCluster) SetMaxAge(v float32) *FigureCluster {
	s.MaxAge = &v
	return s
}

func (s *FigureCluster) SetMetaLockVersion(v int64) *FigureCluster {
	s.MetaLockVersion = &v
	return s
}

func (s *FigureCluster) SetMinAge(v float32) *FigureCluster {
	s.MinAge = &v
	return s
}

func (s *FigureCluster) SetName(v string) *FigureCluster {
	s.Name = &v
	return s
}

func (s *FigureCluster) SetObjectId(v string) *FigureCluster {
	s.ObjectId = &v
	return s
}

func (s *FigureCluster) SetObjectType(v string) *FigureCluster {
	s.ObjectType = &v
	return s
}

func (s *FigureCluster) SetOwnerId(v string) *FigureCluster {
	s.OwnerId = &v
	return s
}

func (s *FigureCluster) SetProjectName(v string) *FigureCluster {
	s.ProjectName = &v
	return s
}

func (s *FigureCluster) SetUpdateTime(v string) *FigureCluster {
	s.UpdateTime = &v
	return s
}

func (s *FigureCluster) SetVideoCount(v int64) *FigureCluster {
	s.VideoCount = &v
	return s
}

func (s *FigureCluster) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}
