// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *DeleteFileMetaRequest
	GetDatasetName() *string
	SetProjectName(v string) *DeleteFileMetaRequest
	GetProjectName() *string
	SetURI(v string) *DeleteFileMetaRequest
	GetURI() *string
}

type DeleteFileMetaRequest struct {
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-datset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the file in OSS.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the path of the object with the extension included.
	//
	// The URI of the file in Photo and Drive Service must be in the pds://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/exampleobject.txt
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DeleteFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DeleteFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteFileMetaRequest) GetURI() *string {
	return s.URI
}

func (s *DeleteFileMetaRequest) SetDatasetName(v string) *DeleteFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteFileMetaRequest) SetProjectName(v string) *DeleteFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteFileMetaRequest) SetURI(v string) *DeleteFileMetaRequest {
	s.URI = &v
	return s
}

func (s *DeleteFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
