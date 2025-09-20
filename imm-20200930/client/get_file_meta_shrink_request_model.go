// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetFileMetaShrinkRequest
	GetDatasetName() *string
	SetProjectName(v string) *GetFileMetaShrinkRequest
	GetProjectName() *string
	SetURI(v string) *GetFileMetaShrinkRequest
	GetURI() *string
	SetWithFieldsShrink(v string) *GetFileMetaShrinkRequest
	GetWithFieldsShrink() *string
}

type GetFileMetaShrinkRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the file. Make sure that the file is indexed****.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// Specify the URI of the file in Photo and Drive Service in the pds://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI              *string `json:"URI,omitempty" xml:"URI,omitempty"`
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
}

func (s GetFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFileMetaShrinkRequest) GetURI() *string {
	return s.URI
}

func (s *GetFileMetaShrinkRequest) GetWithFieldsShrink() *string {
	return s.WithFieldsShrink
}

func (s *GetFileMetaShrinkRequest) SetDatasetName(v string) *GetFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFileMetaShrinkRequest) SetProjectName(v string) *GetFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileMetaShrinkRequest) SetURI(v string) *GetFileMetaShrinkRequest {
	s.URI = &v
	return s
}

func (s *GetFileMetaShrinkRequest) SetWithFieldsShrink(v string) *GetFileMetaShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *GetFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
