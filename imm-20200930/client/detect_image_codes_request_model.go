// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageCodesRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageCodesRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageCodesRequest
	GetSourceURI() *string
}

type DetectImageCodesRequest struct {
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

func (s DetectImageCodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCodesRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCodesRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageCodesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageCodesRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageCodesRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageCodesRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageCodesRequest) SetProjectName(v string) *DetectImageCodesRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCodesRequest) SetSourceURI(v string) *DetectImageCodesRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageCodesRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
