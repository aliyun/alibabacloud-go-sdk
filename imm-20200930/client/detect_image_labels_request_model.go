// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *DetectImageLabelsRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *DetectImageLabelsRequest
	GetProjectName() *string
	SetSourceURI(v string) *DetectImageLabelsRequest
	GetSourceURI() *string
	SetThreshold(v float32) *DetectImageLabelsRequest
	GetThreshold() *float32
}

type DetectImageLabelsRequest struct {
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
	// immimagetest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket in which you store the image.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the image file that has an extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-test/testcases/facetest.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The threshold of the label confidence. Labels whose confidence is lower than the specified threshold are not returned in the response. Valid values: 0 to 1. If you leave this parameter empty, the algorithm provides a default threshold.
	//
	// example:
	//
	// 1
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectImageLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageLabelsRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *DetectImageLabelsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DetectImageLabelsRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *DetectImageLabelsRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DetectImageLabelsRequest) SetCredentialConfig(v *CredentialConfig) *DetectImageLabelsRequest {
	s.CredentialConfig = v
	return s
}

func (s *DetectImageLabelsRequest) SetProjectName(v string) *DetectImageLabelsRequest {
	s.ProjectName = &v
	return s
}

func (s *DetectImageLabelsRequest) SetSourceURI(v string) *DetectImageLabelsRequest {
	s.SourceURI = &v
	return s
}

func (s *DetectImageLabelsRequest) SetThreshold(v float32) *DetectImageLabelsRequest {
	s.Threshold = &v
	return s
}

func (s *DetectImageLabelsRequest) Validate() error {
	return dara.Validate(s)
}
