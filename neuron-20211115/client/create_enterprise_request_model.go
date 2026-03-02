// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Enterprise) *CreateEnterpriseRequest
	GetBody() *Enterprise
}

type CreateEnterpriseRequest struct {
	// This parameter is required.
	Body *Enterprise `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseRequest) GetBody() *Enterprise {
	return s.Body
}

func (s *CreateEnterpriseRequest) SetBody(v *Enterprise) *CreateEnterpriseRequest {
	s.Body = v
	return s
}

func (s *CreateEnterpriseRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
