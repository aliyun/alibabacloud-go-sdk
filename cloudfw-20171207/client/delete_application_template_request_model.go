// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplateId(v string) *DeleteApplicationTemplateRequest
	GetApplicationTemplateId() *string
}

type DeleteApplicationTemplateRequest struct {
	// The ID of the application program template to delete. This parameter is required. If you leave this parameter empty, a parameter fault is returned. The template must exist within the current Alibaba Cloud account. You cannot delete a template if it contains custom application program rules or is referenced by an access control policy of the Internet border, NAT, or virtual private cloud (VPC) firewalls.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
}

func (s DeleteApplicationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationTemplateRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *DeleteApplicationTemplateRequest) SetApplicationTemplateId(v string) *DeleteApplicationTemplateRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *DeleteApplicationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
