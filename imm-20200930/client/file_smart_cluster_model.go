// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileSmartCluster interface {
	dara.Model
	String() string
	GoString() string
	SetSimilarity(v float32) *FileSmartCluster
	GetSimilarity() *float32
	SetSmartClusterId(v string) *FileSmartCluster
	GetSmartClusterId() *string
}

type FileSmartCluster struct {
	Similarity     *float32 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	SmartClusterId *string  `json:"SmartClusterId,omitempty" xml:"SmartClusterId,omitempty"`
}

func (s FileSmartCluster) String() string {
	return dara.Prettify(s)
}

func (s FileSmartCluster) GoString() string {
	return s.String()
}

func (s *FileSmartCluster) GetSimilarity() *float32 {
	return s.Similarity
}

func (s *FileSmartCluster) GetSmartClusterId() *string {
	return s.SmartClusterId
}

func (s *FileSmartCluster) SetSimilarity(v float32) *FileSmartCluster {
	s.Similarity = &v
	return s
}

func (s *FileSmartCluster) SetSmartClusterId(v string) *FileSmartCluster {
	s.SmartClusterId = &v
	return s
}

func (s *FileSmartCluster) Validate() error {
	return dara.Validate(s)
}
