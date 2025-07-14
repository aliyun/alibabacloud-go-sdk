// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *CreateWebApplicationRequest
	GetNamespaceId() *string
	SetBody(v *CreateWebApplicationInput) *CreateWebApplicationRequest
	GetBody() *CreateWebApplicationInput
}

type CreateWebApplicationRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The information about the application.
	//
	// This parameter is required.
	Body *CreateWebApplicationInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateWebApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateWebApplicationRequest) GetBody() *CreateWebApplicationInput {
	return s.Body
}

func (s *CreateWebApplicationRequest) SetNamespaceId(v string) *CreateWebApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateWebApplicationRequest) SetBody(v *CreateWebApplicationInput) *CreateWebApplicationRequest {
	s.Body = v
	return s
}

func (s *CreateWebApplicationRequest) Validate() error {
	return dara.Validate(s)
}
