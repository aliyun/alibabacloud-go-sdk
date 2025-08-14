// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTextGenerateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitTextGenerateJobRequest
	GetDescription() *string
	SetGenerateConfig(v string) *SubmitTextGenerateJobRequest
	GetGenerateConfig() *string
	SetTitle(v string) *SubmitTextGenerateJobRequest
	GetTitle() *string
	SetType(v string) *SubmitTextGenerateJobRequest
	GetType() *string
	SetUserData(v string) *SubmitTextGenerateJobRequest
	GetUserData() *string
}

type SubmitTextGenerateJobRequest struct {
	// The job description, which can be up to 1,024 bytes in length and must be encoded in UTF-8.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The text generation configurations, including keywords and the requirements for the word count and number of output copies.
	GenerateConfig *string `json:"GenerateConfig,omitempty" xml:"GenerateConfig,omitempty"`
	// The job title.
	//
	// The job title can be up to 128 bytes in length.
	//
	// The value must be encoded in UTF-8.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The job type.
	//
	// Valid values:
	//
	// 	- MarketingCopy: the marketing copy.
	//
	// 	- Title: the short video title.
	//
	// example:
	//
	// MarketingCopy
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user-defined data in the JSON format, which can be up to 512 bytes in length. You can specify a custom callback URL. For more information, see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTextGenerateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTextGenerateJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTextGenerateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitTextGenerateJobRequest) GetGenerateConfig() *string {
	return s.GenerateConfig
}

func (s *SubmitTextGenerateJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitTextGenerateJobRequest) GetType() *string {
	return s.Type
}

func (s *SubmitTextGenerateJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTextGenerateJobRequest) SetDescription(v string) *SubmitTextGenerateJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitTextGenerateJobRequest) SetGenerateConfig(v string) *SubmitTextGenerateJobRequest {
	s.GenerateConfig = &v
	return s
}

func (s *SubmitTextGenerateJobRequest) SetTitle(v string) *SubmitTextGenerateJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitTextGenerateJobRequest) SetType(v string) *SubmitTextGenerateJobRequest {
	s.Type = &v
	return s
}

func (s *SubmitTextGenerateJobRequest) SetUserData(v string) *SubmitTextGenerateJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTextGenerateJobRequest) Validate() error {
	return dara.Validate(s)
}
