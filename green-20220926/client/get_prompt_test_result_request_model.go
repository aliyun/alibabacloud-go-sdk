// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetPromptTestResultRequest
	GetConfig() *string
	SetResourceType(v string) *GetPromptTestResultRequest
	GetResourceType() *string
	SetServiceCode(v string) *GetPromptTestResultRequest
	GetServiceCode() *string
	SetText(v string) *GetPromptTestResultRequest
	GetText() *string
	SetType(v string) *GetPromptTestResultRequest
	GetType() *string
}

type GetPromptTestResultRequest struct {
	// The configuration item. Note: This parameter is actually required. If it is not provided, the service returns a 400 error (config may not be empty).
	//
	// example:
	//
	// {\\"enable\\":false}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The resource type. Note: This parameter is actually required. If it is not provided, the service returns a 400 error (resourceType may not be empty).
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service code. Note: This parameter is actually required. If it is not provided, the service returns a 400 error (serviceCode may not be empty).
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The text content. Note: This parameter is actually required. If it is not provided, the service returns a 400 error.
	//
	// example:
	//
	// Test text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The business type. Note: This parameter is actually required. If it is not provided, the service returns a 400 error.
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPromptTestResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTestResultRequest) GoString() string {
	return s.String()
}

func (s *GetPromptTestResultRequest) GetConfig() *string {
	return s.Config
}

func (s *GetPromptTestResultRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetPromptTestResultRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetPromptTestResultRequest) GetText() *string {
	return s.Text
}

func (s *GetPromptTestResultRequest) GetType() *string {
	return s.Type
}

func (s *GetPromptTestResultRequest) SetConfig(v string) *GetPromptTestResultRequest {
	s.Config = &v
	return s
}

func (s *GetPromptTestResultRequest) SetResourceType(v string) *GetPromptTestResultRequest {
	s.ResourceType = &v
	return s
}

func (s *GetPromptTestResultRequest) SetServiceCode(v string) *GetPromptTestResultRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetPromptTestResultRequest) SetText(v string) *GetPromptTestResultRequest {
	s.Text = &v
	return s
}

func (s *GetPromptTestResultRequest) SetType(v string) *GetPromptTestResultRequest {
	s.Type = &v
	return s
}

func (s *GetPromptTestResultRequest) Validate() error {
	return dara.Validate(s)
}
