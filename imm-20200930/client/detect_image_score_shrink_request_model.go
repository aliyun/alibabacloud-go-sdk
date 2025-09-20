// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageScoreShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *DetectImageScoreShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *DetectImageScoreShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageScoreShrinkRequest
	GetSourceURI() *string
}

type DetectImageScoreShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s DetectImageScoreShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageScoreShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectImageScoreShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *DetectImageScoreShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageScoreShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageScoreShrinkRequest) SetCredentialConfigShrink(v string) *DetectImageScoreShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *DetectImageScoreShrinkRequest) SetProjectName(v string) *DetectImageScoreShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageScoreShrinkRequest) SetSourceURI(v string) *DetectImageScoreShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageScoreShrinkRequest) Validate() error {
	return dara.Validate(s)
}
