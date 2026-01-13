// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFormalServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ApplyFormalServiceCmd) *ApplyFormalServiceRequest
	GetBody() *ApplyFormalServiceCmd
}

type ApplyFormalServiceRequest struct {
	Body *ApplyFormalServiceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFormalServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyFormalServiceRequest) GoString() string {
	return s.String()
}

func (s *ApplyFormalServiceRequest) GetBody() *ApplyFormalServiceCmd {
	return s.Body
}

func (s *ApplyFormalServiceRequest) SetBody(v *ApplyFormalServiceCmd) *ApplyFormalServiceRequest {
	s.Body = v
	return s
}

func (s *ApplyFormalServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
