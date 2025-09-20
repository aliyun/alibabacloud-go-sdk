// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageScoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageScoreRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageScoreRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageScoreRequest
	GetSourceURI() *string
}

type DetectImageScoreRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project.[](~~477051~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the input image.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://bucketname/objectname
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageScoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageScoreRequest) GoString() string {
	return s.String()
}

func (s *DetectImageScoreRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageScoreRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageScoreRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageScoreRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageScoreRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageScoreRequest) SetProjectName(v string) *DetectImageScoreRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageScoreRequest) SetSourceURI(v string) *DetectImageScoreRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageScoreRequest) Validate() error {
	return dara.Validate(s)
}
