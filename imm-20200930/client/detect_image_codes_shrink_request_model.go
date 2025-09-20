// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectImageCodesShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectImageCodesShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageCodesShrinkRequest
	GetSourceURI() *string
}

type DetectImageCodesShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket in which the image file is stored.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the file that has an extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucketname/objectname
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCodesShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectImageCodesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageCodesShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageCodesShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageCodesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageCodesShrinkRequest) SetProjectName(v string) *DetectImageCodesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCodesShrinkRequest) SetSourceURI(v string) *DetectImageCodesShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageCodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
