// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectMediaMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectMediaMetaShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectMediaMetaShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectMediaMetaShrinkRequest
	GetSourceURI() *string
}

type DetectMediaMetaShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s DetectMediaMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectMediaMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectMediaMetaShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectMediaMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectMediaMetaShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectMediaMetaShrinkRequest) SetCredentialConfigShrink(v string) *DetectMediaMetaShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectMediaMetaShrinkRequest) SetProjectName(v string) *DetectMediaMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectMediaMetaShrinkRequest) SetSourceURI(v string) *DetectMediaMetaShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectMediaMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
