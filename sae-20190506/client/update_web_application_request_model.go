// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateWebApplicationRequest
	GetNamespaceId() *string
	SetBody(v *UpdateWebApplicationInput) *UpdateWebApplicationRequest
	GetBody() *UpdateWebApplicationInput
}

type UpdateWebApplicationRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// Updates the information about a web application.
	//
	// This parameter is required.
	Body *UpdateWebApplicationInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateWebApplicationRequest) GetBody() *UpdateWebApplicationInput {
	return s.Body
}

func (s *UpdateWebApplicationRequest) SetNamespaceId(v string) *UpdateWebApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateWebApplicationRequest) SetBody(v *UpdateWebApplicationInput) *UpdateWebApplicationRequest {
	s.Body = v
	return s
}

func (s *UpdateWebApplicationRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
