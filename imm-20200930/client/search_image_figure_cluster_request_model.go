// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageFigureClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *SearchImageFigureClusterRequest
	GetCredentialConfig() *CredentialConfig
	SetDatasetName(v string) *SearchImageFigureClusterRequest
	GetDatasetName() *string
	SetProjectName(v string) *SearchImageFigureClusterRequest
	GetProjectName() *string
	SetSourceURI(v string) *SearchImageFigureClusterRequest
	GetSourceURI() *string
}

type SearchImageFigureClusterRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	// The OSS URI of the image.
	//
	// Specify the OSS URI in the `oss://${Bucket}/${Object}` format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s SearchImageFigureClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageFigureClusterRequest) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *SearchImageFigureClusterRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SearchImageFigureClusterRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SearchImageFigureClusterRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *SearchImageFigureClusterRequest) SetCredentialConfig(v *CredentialConfig) *SearchImageFigureClusterRequest {
	s.CredentialConfig = v
	return s
}

func (s *SearchImageFigureClusterRequest) SetDatasetName(v string) *SearchImageFigureClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchImageFigureClusterRequest) SetProjectName(v string) *SearchImageFigureClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *SearchImageFigureClusterRequest) SetSourceURI(v string) *SearchImageFigureClusterRequest {
	s.SourceURI = &v
	return s
}

func (s *SearchImageFigureClusterRequest) Validate() error {
	return dara.Validate(s)
}
