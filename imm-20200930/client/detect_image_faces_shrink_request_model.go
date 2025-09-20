// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageFacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectImageFacesShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectImageFacesShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageFacesShrinkRequest
	GetSourceURI() *string
}

type DetectImageFacesShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s DetectImageFacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageFacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageFacesShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectImageFacesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageFacesShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageFacesShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageFacesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageFacesShrinkRequest) SetProjectName(v string) *DetectImageFacesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageFacesShrinkRequest) SetSourceURI(v string) *DetectImageFacesShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageFacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
