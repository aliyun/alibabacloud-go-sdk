// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWebFilterTemplateRequest
	GetDescription() *string
	SetName(v string) *CreateWebFilterTemplateRequest
	GetName() *string
	SetStatus(v string) *CreateWebFilterTemplateRequest
	GetStatus() *string
}

type CreateWebFilterTemplateRequest struct {
	// The description of the web filtering template. The description can be up to 128 characters in length and can be empty.
	//
	// example:
	//
	// Web filtering template for URL access control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the web filtering template. The name must be 1 to 64 characters in length and can contain Chinese characters, letters, digits, hyphens (-), underscores (_), and periods (.). It cannot contain spaces or be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebFilterTemplate01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the web filtering template. Valid values:
	//
	// - **Enable**: enabled.
	//
	// - **Disable**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateWebFilterTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateWebFilterTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWebFilterTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateWebFilterTemplateRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateWebFilterTemplateRequest) SetDescription(v string) *CreateWebFilterTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateWebFilterTemplateRequest) SetName(v string) *CreateWebFilterTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateWebFilterTemplateRequest) SetStatus(v string) *CreateWebFilterTemplateRequest {
	s.Status = &v
	return s
}

func (s *CreateWebFilterTemplateRequest) Validate() error {
	return dara.Validate(s)
}
