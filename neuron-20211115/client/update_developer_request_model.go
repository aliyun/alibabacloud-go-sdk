// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeveloperInfoUpdateCmd) *UpdateDeveloperRequest
	GetBody() *DeveloperInfoUpdateCmd
}

type UpdateDeveloperRequest struct {
	// This parameter is required.
	Body *DeveloperInfoUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeveloperRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeveloperRequest) GetBody() *DeveloperInfoUpdateCmd {
	return s.Body
}

func (s *UpdateDeveloperRequest) SetBody(v *DeveloperInfoUpdateCmd) *UpdateDeveloperRequest {
	s.Body = v
	return s
}

func (s *UpdateDeveloperRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
