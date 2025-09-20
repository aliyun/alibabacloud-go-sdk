// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimilarImageCluster interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *SimilarImageCluster
	GetCreateTime() *string
	SetCustomLabels(v map[string]interface{}) *SimilarImageCluster
	GetCustomLabels() map[string]interface{}
	SetFiles(v []*SimilarImage) *SimilarImageCluster
	GetFiles() []*SimilarImage
	SetObjectId(v string) *SimilarImageCluster
	GetObjectId() *string
	SetUpdateTime(v string) *SimilarImageCluster
	GetUpdateTime() *string
}

type SimilarImageCluster struct {
	CreateTime   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Files        []*SimilarImage        `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	ObjectId     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	UpdateTime   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s SimilarImageCluster) String() string {
	return dara.Prettify(s)
}

func (s SimilarImageCluster) GoString() string {
	return s.String()
}

func (s *SimilarImageCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SimilarImageCluster) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *SimilarImageCluster) GetFiles() []*SimilarImage {
	return s.Files
}

func (s *SimilarImageCluster) GetObjectId() *string {
	return s.ObjectId
}

func (s *SimilarImageCluster) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SimilarImageCluster) SetCreateTime(v string) *SimilarImageCluster {
	s.CreateTime = &v
	return s
}

func (s *SimilarImageCluster) SetCustomLabels(v map[string]interface{}) *SimilarImageCluster {
	s.CustomLabels = v
	return s
}

func (s *SimilarImageCluster) SetFiles(v []*SimilarImage) *SimilarImageCluster {
	s.Files = v
	return s
}

func (s *SimilarImageCluster) SetObjectId(v string) *SimilarImageCluster {
	s.ObjectId = &v
	return s
}

func (s *SimilarImageCluster) SetUpdateTime(v string) *SimilarImageCluster {
	s.UpdateTime = &v
	return s
}

func (s *SimilarImageCluster) Validate() error {
	return dara.Validate(s)
}
