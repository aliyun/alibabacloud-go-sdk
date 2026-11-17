// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateWebFilterTemplateRequest
	GetDescription() *string
	SetName(v string) *UpdateWebFilterTemplateRequest
	GetName() *string
	SetStatus(v string) *UpdateWebFilterTemplateRequest
	GetStatus() *string
	SetWebFilterTemplateId(v string) *UpdateWebFilterTemplateRequest
	GetWebFilterTemplateId() *string
}

type UpdateWebFilterTemplateRequest struct {
	// The description of the web filtering template. The description can be up to 128 characters in length. If this parameter is not specified, the original value is retained. If an empty string is specified, the description is cleared.
	//
	// example:
	//
	// Web filtering template for URL access control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the web filtering template. The name must be 1 to 64 characters in length and can contain Chinese characters, English letters, digits, hyphens (-), underscores (_), and periods (.). It cannot contain spaces. If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Web filtering template 02
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the web filtering template. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the web filtering template to be modified. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s UpdateWebFilterTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWebFilterTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWebFilterTemplateRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateWebFilterTemplateRequest) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *UpdateWebFilterTemplateRequest) SetDescription(v string) *UpdateWebFilterTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateWebFilterTemplateRequest) SetName(v string) *UpdateWebFilterTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateWebFilterTemplateRequest) SetStatus(v string) *UpdateWebFilterTemplateRequest {
	s.Status = &v
	return s
}

func (s *UpdateWebFilterTemplateRequest) SetWebFilterTemplateId(v string) *UpdateWebFilterTemplateRequest {
	s.WebFilterTemplateId = &v
	return s
}

func (s *UpdateWebFilterTemplateRequest) Validate() error {
	return dara.Validate(s)
}
