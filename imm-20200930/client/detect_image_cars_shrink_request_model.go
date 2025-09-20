// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCarsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectImageCarsShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectImageCarsShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageCarsShrinkRequest
	GetSourceURI() *string
}

type DetectImageCarsShrinkRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The authorization chain. This parameter is optional. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket in which you store the image file.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the file that has an extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCarsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCarsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCarsShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectImageCarsShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageCarsShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageCarsShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageCarsShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageCarsShrinkRequest) SetProjectName(v string) *DetectImageCarsShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCarsShrinkRequest) SetSourceURI(v string) *DetectImageCarsShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageCarsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
