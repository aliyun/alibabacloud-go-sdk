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
	// **Leave this parameter empty unless you have special requirements.**
	//
	// The chain authorization configuration. This parameter is optional. For more information, see [Use chain authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
