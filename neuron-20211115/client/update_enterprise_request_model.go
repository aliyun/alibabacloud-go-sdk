// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *EnterpriseInfoUpdateCmd) *UpdateEnterpriseRequest
	GetBody() *EnterpriseInfoUpdateCmd
}

type UpdateEnterpriseRequest struct {
	// This parameter is required.
	Body *EnterpriseInfoUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnterpriseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseRequest) GetBody() *EnterpriseInfoUpdateCmd {
	return s.Body
}

func (s *UpdateEnterpriseRequest) SetBody(v *EnterpriseInfoUpdateCmd) *UpdateEnterpriseRequest {
	s.Body = v
	return s
}

func (s *UpdateEnterpriseRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
