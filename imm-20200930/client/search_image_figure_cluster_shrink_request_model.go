// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageFigureClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *SearchImageFigureClusterShrinkRequest
	GetCredentialConfigShrink() *string
	SetDatasetName(v string) *SearchImageFigureClusterShrinkRequest
	GetDatasetName() *string
	SetProjectName(v string) *SearchImageFigureClusterShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *SearchImageFigureClusterShrinkRequest
	GetSourceURI() *string
}

type SearchImageFigureClusterShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s SearchImageFigureClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageFigureClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *SearchImageFigureClusterShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SearchImageFigureClusterShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SearchImageFigureClusterShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *SearchImageFigureClusterShrinkRequest) SetCredentialConfigShrink(v string) *SearchImageFigureClusterShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) SetDatasetName(v string) *SearchImageFigureClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) SetProjectName(v string) *SearchImageFigureClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) SetSourceURI(v string) *SearchImageFigureClusterShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *SearchImageFigureClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
