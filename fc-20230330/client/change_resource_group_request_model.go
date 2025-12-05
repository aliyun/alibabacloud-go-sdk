// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ChangeResourceGroupInput) *ChangeResourceGroupRequest
	GetBody() *ChangeResourceGroupInput
}

type ChangeResourceGroupRequest struct {
	// The request details for updating the resource group.
	Body *ChangeResourceGroupInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetBody() *ChangeResourceGroupInput {
	return s.Body
}

func (s *ChangeResourceGroupRequest) SetBody(v *ChangeResourceGroupInput) *ChangeResourceGroupRequest {
	s.Body = v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
