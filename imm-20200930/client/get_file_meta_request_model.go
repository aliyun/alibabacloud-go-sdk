// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetFileMetaRequest
	GetDatasetName() *string
	SetProjectName(v string) *GetFileMetaRequest
	GetProjectName() *string
	SetURI(v string) *GetFileMetaRequest
	GetURI() *string
	SetWithFields(v []*string) *GetFileMetaRequest
	GetWithFields() []*string
}

type GetFileMetaRequest struct {
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
	URI        *string   `json:"URI,omitempty" xml:"URI,omitempty"`
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
}

func (s GetFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *GetFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFileMetaRequest) GetURI() *string {
	return s.URI
}

func (s *GetFileMetaRequest) GetWithFields() []*string {
	return s.WithFields
}

func (s *GetFileMetaRequest) SetDatasetName(v string) *GetFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *GetFileMetaRequest) SetProjectName(v string) *GetFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *GetFileMetaRequest) SetURI(v string) *GetFileMetaRequest {
	s.URI = &v
	return s
}

func (s *GetFileMetaRequest) SetWithFields(v []*string) *GetFileMetaRequest {
	s.WithFields = v
	return s
}

func (s *GetFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
