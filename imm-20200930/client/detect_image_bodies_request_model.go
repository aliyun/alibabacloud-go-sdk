// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBodiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageBodiesRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageBodiesRequest
	GetProjectName() *string
	SetSensitivity(v float32) *DetectImageBodiesRequest
	GetSensitivity() *float32
	SetSourceURI(v string) *DetectImageBodiesRequest
	GetSourceURI() *string
}

type DetectImageBodiesRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s DetectImageBodiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBodiesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBodiesRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageBodiesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageBodiesRequest) GetSensitivity() *float32 {
	return s.Sensitivity
}

func (s *DetectImageBodiesRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageBodiesRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageBodiesRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageBodiesRequest) SetProjectName(v string) *DetectImageBodiesRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageBodiesRequest) SetSensitivity(v float32) *DetectImageBodiesRequest {
	s.Sensitivity = &v
	return s
}

func (s *DetectImageBodiesRequest) SetSourceURI(v string) *DetectImageBodiesRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageBodiesRequest) Validate() error {
	return dara.Validate(s)
}
