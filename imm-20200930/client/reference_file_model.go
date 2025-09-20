// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReferenceFile interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *ReferenceFile
	GetDatasetName() *string
	SetObjectId(v string) *ReferenceFile
	GetObjectId() *string
	SetProjectName(v string) *ReferenceFile
	GetProjectName() *string
	SetURI(v string) *ReferenceFile
	GetURI() *string
}

type ReferenceFile struct {
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// 75d5de2c50754e3dadd5c35dbca5f9949369e37eb342a73821f690c94c36c7f7
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// oss://test-bucket/test-object.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ReferenceFile) String() string {
	return dara.Prettify(s)
}

func (s ReferenceFile) GoString() string {
	return s.String()
}

func (s *ReferenceFile) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ReferenceFile) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReferenceFile) GetProjectName() *string {
	return s.ProjectName
}

func (s *ReferenceFile) GetURI() *string {
	return s.URI
}

func (s *ReferenceFile) SetDatasetName(v string) *ReferenceFile {
	s.DatasetName = &v
	return s
}

func (s *ReferenceFile) SetObjectId(v string) *ReferenceFile {
	s.ObjectId = &v
	return s
}

func (s *ReferenceFile) SetProjectName(v string) *ReferenceFile {
	s.ProjectName = &v
	return s
}

func (s *ReferenceFile) SetURI(v string) *ReferenceFile {
	s.URI = &v
	return s
}

func (s *ReferenceFile) Validate() error {
	return dara.Validate(s)
}
