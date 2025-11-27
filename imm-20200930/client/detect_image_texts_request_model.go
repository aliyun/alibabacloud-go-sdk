// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageTextsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageTextsRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageTextsRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageTextsRequest
	GetSourceURI() *string
}

type DetectImageTextsRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s DetectImageTextsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageTextsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageTextsRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageTextsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageTextsRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageTextsRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageTextsRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageTextsRequest) SetProjectName(v string) *DetectImageTextsRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageTextsRequest) SetSourceURI(v string) *DetectImageTextsRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageTextsRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
