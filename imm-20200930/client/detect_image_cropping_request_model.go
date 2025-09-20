// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCroppingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatios(v string) *DetectImageCroppingRequest
	GetAspectRatios() *string
	SetCredentialConfig(v *CredentialConfig) *DetectImageCroppingRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageCroppingRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageCroppingRequest
	GetSourceURI() *string
}

type DetectImageCroppingRequest struct {
	// The cropping ratios. You can specify up to five cropping ratios. Take note of the following requirements:
	//
	// 	- The ratio must be an integer between 0 and 20.
	//
	// 	- The ratio must range from 0.5 to 2.
	//
	// 	- If you leave this parameter empty, the default processing logic is `["auto"]`.
	//
	// >  Errors are reported in one of the following cases:\\
	//
	// You specify more than five cropping ratios.\\
	//
	// You pass an empty list to the system.\\
	//
	// You specify a ratio that is not an integer, such as `4.1:3`.\\
	//
	// The ratio is beyond the range of 0.5 to 2.
	//
	// example:
	//
	// ["1:1"]
	AspectRatios *string `json:"AspectRatios,omitempty" xml:"AspectRatios,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket in which you store the image.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the image file that has an extension.
	//
	// example:
	//
	// oss://imm-test/testcases/facetest.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
}

func (s DetectImageCroppingRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCroppingRequest) GoString() string {
	return s.String()
}

func (s *DetectImageCroppingRequest) GetAspectRatios() *string {
	return s.AspectRatios
}

func (s *DetectImageCroppingRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageCroppingRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageCroppingRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageCroppingRequest) SetAspectRatios(v string) *DetectImageCroppingRequest {
	s.AspectRatios = &v
	return s
}

func (s *DetectImageCroppingRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageCroppingRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageCroppingRequest) SetProjectName(v string) *DetectImageCroppingRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageCroppingRequest) SetSourceURI(v string) *DetectImageCroppingRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageCroppingRequest) Validate() error {
	return dara.Validate(s)
}
