// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcVersionCmd) *UpdatePbcVersionRequest
	GetBody() *PbcVersionCmd
}

type UpdatePbcVersionRequest struct {
	Body *PbcVersionCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePbcVersionRequest) GetBody() *PbcVersionCmd {
	return s.Body
}

func (s *UpdatePbcVersionRequest) SetBody(v *PbcVersionCmd) *UpdatePbcVersionRequest {
	s.Body = v
	return s
}

func (s *UpdatePbcVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
