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
	// example:
	//
	// {\\"enable\\":false}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
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
