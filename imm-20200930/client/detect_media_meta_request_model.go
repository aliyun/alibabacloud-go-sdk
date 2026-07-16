// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMediaMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectMediaMetaRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectMediaMetaRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectMediaMetaRequest
	GetSourceURI() *string
}

type DetectMediaMetaRequest struct {
	// **Leave this parameter empty unless you have special requirements.**
	//
	// The chain authorization configuration. This parameter is optional. For more information, see [Use chain authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Object Storage Service (OSS) URI of the media file.
	//
	// The OSS URI follows the format oss://${Bucket}/${Object}, where `${Bucket}` is the name of an OSS bucket in the same region as the current project, and `${Object}` is the full path of the file including the file name extension.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.mp4
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectMediaMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectMediaMetaRequest) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectMediaMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectMediaMetaRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectMediaMetaRequest) SetCredentialConfig(v *CredentialConfig) *DetectMediaMetaRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectMediaMetaRequest) SetProjectName(v string) *DetectMediaMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectMediaMetaRequest) SetSourceURI(v string) *DetectMediaMetaRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectMediaMetaRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
