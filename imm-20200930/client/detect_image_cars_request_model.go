// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageCarsRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageCarsRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageCarsRequest
	GetSourceURI() *string
}

type DetectImageCarsRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The authorization chain. This parameter is optional. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s DetectImageCarsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCarsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCarsRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageCarsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageCarsRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageCarsRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageCarsRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageCarsRequest) SetProjectName(v string) *DetectImageCarsRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCarsRequest) SetSourceURI(v string) *DetectImageCarsRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageCarsRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
