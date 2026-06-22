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
	SetInclusionHints(v []*string) *DetectImageCroppingRequest
	GetInclusionHints() []*string
	SetProjectName(v string) *DetectImageCroppingRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageCroppingRequest
	GetSourceURI() *string
}

type DetectImageCroppingRequest struct {
	// The list of cropping aspect ratios. You can specify up to 5 ratios. Each ratio must meet the following requirements:
	//
	// - The ratio must consist of integers in the range of (0, 20).
	//
	// - The ratio value must be in the range of [0.5, 2].
	//
	// - If you do not specify this parameter, the default value `["auto"]` is used.
	//
	// > The following cases cause an error:<br>- More than 5 ratios are specified.<br>- An empty list is passed.<br>- The ratio contains non-integer values, such as `4.1:3`.<br>- The ratio value is less than 0.5 or greater than 2.
	//
	// example:
	//
	// ["1:1"]
	AspectRatios *string `json:"AspectRatios,omitempty" xml:"AspectRatios,omitempty"`
	// **Leave this parameter empty unless otherwise required.**
	//
	// The China authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The list of semantic text descriptions for objects that you want the cropping result to include. Each element is free-form text, such as "signboard" or "dish".
	//
	// > Usage limits of the InclusionHints parameter:
	//
	// <br>- You can pass in up to 1 hint in the array to specify the type of object to include in the cropping result, such as "signboard".
	//
	// <br>- The algorithm detects all objects in the image that match the hint and generates cropping regions that include as many matched objects as possible.
	//
	// <br>- Each cropping region includes up to 10 matched objects. If more than 10 objects match in the image, the cropping region includes up to 10 of them.
	//
	// <br>- You can use the MatchedInclusionHints response field to determine whether the hint was successfully matched.
	//
	// <br>- This parameter is supported only in regions in the Chinese mainland.
	//
	// example:
	//
	// ["sign"]
	InclusionHints []*string `json:"InclusionHints,omitempty" xml:"InclusionHints,omitempty" type:"Repeated"`
	// The project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the image.
	//
	// The OSS URI follows the format oss://${Bucket}/${Object}, where `${Bucket}` is the name of an OSS bucket in the same region as the current project, and `${Object}` is the full path of the file including the file name extension.
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

func (s *DetectImageCroppingRequest) GetInclusionHints() []*string {
	return s.InclusionHints
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

func (s *DetectImageCroppingRequest) SetInclusionHints(v []*string) *DetectImageCroppingRequest {
	s.InclusionHints = v
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
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
