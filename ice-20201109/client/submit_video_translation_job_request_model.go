// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitVideoTranslationJobRequest
	GetClientToken() *string
	SetDescription(v string) *SubmitVideoTranslationJobRequest
	GetDescription() *string
	SetEditingConfig(v string) *SubmitVideoTranslationJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitVideoTranslationJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitVideoTranslationJobRequest
	GetOutputConfig() *string
	SetSignature(v string) *SubmitVideoTranslationJobRequest
	GetSignature() *string
	SetSignatureMehtod(v string) *SubmitVideoTranslationJobRequest
	GetSignatureMehtod() *string
	SetSignatureNonce(v string) *SubmitVideoTranslationJobRequest
	GetSignatureNonce() *string
	SetSignatureType(v string) *SubmitVideoTranslationJobRequest
	GetSignatureType() *string
	SetSignatureVersion(v string) *SubmitVideoTranslationJobRequest
	GetSignatureVersion() *string
	SetTitle(v string) *SubmitVideoTranslationJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitVideoTranslationJobRequest
	GetUserData() *string
}

type SubmitVideoTranslationJobRequest struct {
	// 	- The client token.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 	- The job description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 	- The configuration parameters of the video translation job.
	//
	// 	- The value must be in the JSON format.
	//
	// example:
	//
	// {"SourceLanguage":"zh","TargetLanguage":"en","DetextArea":"Auto"}
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// 	- The input parameters of the video translation job.
	//
	// 	- A video translation job takes a video or subtitle file as the input.
	//
	// 	- The value must be in the JSON format.
	//
	// example:
	//
	// {"Type":"Video","Media":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/xxx.mp4"}
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// 	- The output parameters of the video translation job.
	//
	// 	- A video translation job can generate a video or subtitle file as the output.
	//
	// example:
	//
	// {"MediaURL": "https://your-bucket.oss-cn-shanghai.aliyuncs.com/your-object.mp4"}
	OutputConfig     *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	Signature        *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	SignatureMehtod  *string `json:"SignatureMehtod,omitempty" xml:"SignatureMehtod,omitempty"`
	SignatureNonce   *string `json:"SignatureNonce,omitempty" xml:"SignatureNonce,omitempty"`
	SignatureType    *string `json:"SignatureType,omitempty" xml:"SignatureType,omitempty"`
	SignatureVersion *string `json:"SignatureVersion,omitempty" xml:"SignatureVersion,omitempty"`
	// 	- The job title.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 	- The user-defined data.
	//
	// 	- The data must be in the JSON format, and can be up to 512 characters in length.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoTranslationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitVideoTranslationJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitVideoTranslationJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitVideoTranslationJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitVideoTranslationJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitVideoTranslationJobRequest) GetSignature() *string {
	return s.Signature
}

func (s *SubmitVideoTranslationJobRequest) GetSignatureMehtod() *string {
	return s.SignatureMehtod
}

func (s *SubmitVideoTranslationJobRequest) GetSignatureNonce() *string {
	return s.SignatureNonce
}

func (s *SubmitVideoTranslationJobRequest) GetSignatureType() *string {
	return s.SignatureType
}

func (s *SubmitVideoTranslationJobRequest) GetSignatureVersion() *string {
	return s.SignatureVersion
}

func (s *SubmitVideoTranslationJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitVideoTranslationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoTranslationJobRequest) SetClientToken(v string) *SubmitVideoTranslationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetDescription(v string) *SubmitVideoTranslationJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetEditingConfig(v string) *SubmitVideoTranslationJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetInputConfig(v string) *SubmitVideoTranslationJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetOutputConfig(v string) *SubmitVideoTranslationJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetSignature(v string) *SubmitVideoTranslationJobRequest {
	s.Signature = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetSignatureMehtod(v string) *SubmitVideoTranslationJobRequest {
	s.SignatureMehtod = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetSignatureNonce(v string) *SubmitVideoTranslationJobRequest {
	s.SignatureNonce = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetSignatureType(v string) *SubmitVideoTranslationJobRequest {
	s.SignatureType = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetSignatureVersion(v string) *SubmitVideoTranslationJobRequest {
	s.SignatureVersion = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetTitle(v string) *SubmitVideoTranslationJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetUserData(v string) *SubmitVideoTranslationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) Validate() error {
	return dara.Validate(s)
}
