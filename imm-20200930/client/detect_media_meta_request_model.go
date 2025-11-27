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
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the media object in Object Storage Service (OSS).
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
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
