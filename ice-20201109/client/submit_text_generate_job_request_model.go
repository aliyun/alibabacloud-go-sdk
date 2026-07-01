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
	// The job description, with a maximum length of 1,024 bytes (UTF-8 encoded).
	//
	// example:
	//
	// Test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The generation configuration, in JSON format.
	//
	// If `Type` is set to `Title` or `MarketingCopy`, specify the following fields:
	//
	// - `keywords`: The keywords used to generate the title or marketing copy. This parameter is required.
	//
	// - `textLength`: The target length of the generated text, in characters. Valid values: 5 to 1,000. The actual length of the output is less than or equal to this value. This parameter is required.
	//
	// - `targetCount`: The number of copy variations to generate. Valid values: 1 to 1,000. This parameter is required.
	//
	// If `Type` is set to `StoryboardScript`, specify the following field:
	//
	// - `originText`: The original text used to generate the storyboard script. This parameter is required.
	//
	// example:
	//
	// {"keywords":"New cake shop, animal cream","textLength":100,"targetCount":3}
	//
	// or
	//
	// {"originText": "A new cake shop just opened on the street, selling cream cakes, fruit cakes, bread, muffins, etc. A coffee shop opened across from the cake shop, with a steady stream of customers."}
	GenerateConfig *string `json:"GenerateConfig,omitempty" xml:"GenerateConfig,omitempty"`
	// The job title.
	//
	// \\- The maximum length is 128 bytes.
	//
	// \\- UTF-8 encoding is required.
	//
	// example:
	//
	// Test title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The job type.
	//
	// **Valid values:**
	//
	// - `MarketingCopy`: Generates marketing copy.
	//
	// - `Title`: Generates a short video title.
	//
	// - `StoryboardScript`: Generates a storyboard script from text.
	//
	// example:
	//
	// MarketingCopy
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The custom settings in JSON format. The maximum length is 512 bytes. You can use this parameter to specify a [custom callback address](https://help.aliyun.com/document_detail/451631.html).
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"} or {"NotifyAddress":"https://xx.xx.xxx"} or {"NotifyAddress":"ice-callback-demo"}
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
