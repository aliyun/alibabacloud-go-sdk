// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebApplicationScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateWebApplicationScalingConfigRequest
	GetNamespaceId() *string
	SetBody(v *UpdateWebApplicationScalingConfigInput) *UpdateWebApplicationScalingConfigRequest
	GetBody() *UpdateWebApplicationScalingConfigInput
}

type UpdateWebApplicationScalingConfigRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The information about scaling configurations.
	//
	// This parameter is required.
	Body *UpdateWebApplicationScalingConfigInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebApplicationScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebApplicationScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebApplicationScalingConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateWebApplicationScalingConfigRequest) GetBody() *UpdateWebApplicationScalingConfigInput {
	return s.Body
}

func (s *UpdateWebApplicationScalingConfigRequest) SetNamespaceId(v string) *UpdateWebApplicationScalingConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateWebApplicationScalingConfigRequest) SetBody(v *UpdateWebApplicationScalingConfigInput) *UpdateWebApplicationScalingConfigRequest {
	s.Body = v
	return s
}

func (s *UpdateWebApplicationScalingConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
