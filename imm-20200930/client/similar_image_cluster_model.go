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
	// The creation time.
	//
	// example:
	//
	// 2023-02-08T09:42:34.354969088+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom tag.
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The similar images.
	Files []*SimilarImage `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The ID of the cluster.
	//
	// example:
	//
	// SimilarImageCluster-748a041e-4ebc-4487-9e74-9c89b1****
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2023-02-08T09:42:34.354969088+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
