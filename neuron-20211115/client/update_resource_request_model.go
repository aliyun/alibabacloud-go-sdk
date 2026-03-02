// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ResourceUpdateCmd) *UpdateResourceRequest
	GetBody() *ResourceUpdateCmd
}

type UpdateResourceRequest struct {
	Body *ResourceUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) GetBody() *ResourceUpdateCmd {
	return s.Body
}

func (s *UpdateResourceRequest) SetBody(v *ResourceUpdateCmd) *UpdateResourceRequest {
	s.Body = v
	return s
}

func (s *UpdateResourceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
