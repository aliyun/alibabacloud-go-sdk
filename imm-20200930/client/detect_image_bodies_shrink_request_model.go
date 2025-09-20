// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBodiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectImageBodiesShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectImageBodiesShrinkRequest
	GetProjectName() *string
	SetSensitivity(v float32) *DetectImageBodiesShrinkRequest
	GetSensitivity() *float32
	SetSourceURI(v string) *DetectImageBodiesShrinkRequest
	GetSourceURI() *string
}

type DetectImageBodiesShrinkRequest struct {
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
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The accuracy level of detecting and recognizing specific content in the image. Valid values: 0 to 1. Default value: 0.6. A higher sensitivity specifies that more image details can be detected.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0.6
	Sensitivity *float32 `json:"Sensitivity,omitempty" xml:"Sensitivity,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket in which the image file is stored.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the file that has an extension.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageBodiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBodiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectImageBodiesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageBodiesShrinkRequest) GetSensitivity() *float32 {
	return s.Sensitivity
}

func (s *DetectImageBodiesShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageBodiesShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageBodiesShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) SetProjectName(v string) *DetectImageBodiesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) SetSensitivity(v float32) *DetectImageBodiesShrinkRequest {
	s.Sensitivity = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) SetSourceURI(v string) *DetectImageBodiesShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageBodiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
