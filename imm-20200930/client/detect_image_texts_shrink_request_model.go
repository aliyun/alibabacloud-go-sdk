// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageTextsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectImageTextsShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectImageTextsShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageTextsShrinkRequest
	GetSourceURI() *string
}

type DetectImageTextsShrinkRequest struct {
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
	// The Object Storage Service (OSS) URI of the file.
	//
	// Specify the URI in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of an OSS bucket that is in the same region as the current project. ${Object} specifies the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageTextsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageTextsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageTextsShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectImageTextsShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageTextsShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageTextsShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageTextsShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageTextsShrinkRequest) SetProjectName(v string) *DetectImageTextsShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageTextsShrinkRequest) SetSourceURI(v string) *DetectImageTextsShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageTextsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
