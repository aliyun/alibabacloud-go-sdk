// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationTrafficConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateWebApplicationTrafficConfigRequest
	GetNamespaceId() *string
	SetBody(v *UpdateWebApplicationTrafficConfigInput) *UpdateWebApplicationTrafficConfigRequest
	GetBody() *UpdateWebApplicationTrafficConfigInput
}

type UpdateWebApplicationTrafficConfigRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The traffic configurations.
	//
	// This parameter is required.
	Body *UpdateWebApplicationTrafficConfigInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebApplicationTrafficConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationTrafficConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationTrafficConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateWebApplicationTrafficConfigRequest) GetBody() *UpdateWebApplicationTrafficConfigInput {
	return s.Body
}

func (s *UpdateWebApplicationTrafficConfigRequest) SetNamespaceId(v string) *UpdateWebApplicationTrafficConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateWebApplicationTrafficConfigRequest) SetBody(v *UpdateWebApplicationTrafficConfigInput) *UpdateWebApplicationTrafficConfigRequest {
	s.Body = v
	return s
}

func (s *UpdateWebApplicationTrafficConfigRequest) Validate() error {
	return dara.Validate(s)
}
