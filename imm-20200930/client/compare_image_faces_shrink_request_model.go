// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareImageFacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CompareImageFacesShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *CompareImageFacesShrinkRequest
	GetProjectName() *string
	SetSourceShrink(v string) *CompareImageFacesShrinkRequest
	GetSourceShrink() *string
}

type CompareImageFacesShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URLs of the two images for compression.
	SourceShrink *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CompareImageFacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareImageFacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CompareImageFacesShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CompareImageFacesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CompareImageFacesShrinkRequest) GetSourceShrink() *string {
	return s.SourceShrink
}

func (s *CompareImageFacesShrinkRequest) SetCredentialConfigShrink(v string) *CompareImageFacesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CompareImageFacesShrinkRequest) SetProjectName(v string) *CompareImageFacesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CompareImageFacesShrinkRequest) SetSourceShrink(v string) *CompareImageFacesShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *CompareImageFacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
