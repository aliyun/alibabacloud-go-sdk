// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageFacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageFacesRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageFacesRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageFacesRequest
	GetSourceURI() *string
}

type DetectImageFacesRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the image object.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageFacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageFacesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageFacesRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageFacesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageFacesRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageFacesRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageFacesRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageFacesRequest) SetProjectName(v string) *DetectImageFacesRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageFacesRequest) SetSourceURI(v string) *DetectImageFacesRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageFacesRequest) Validate() error {
	return dara.Validate(s)
}
