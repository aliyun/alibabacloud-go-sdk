// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcVersionCmd) *CreatePbcVersionRequest
	GetBody() *PbcVersionCmd
}

type CreatePbcVersionRequest struct {
	Body *PbcVersionCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcVersionRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcVersionRequest) GetBody() *PbcVersionCmd {
	return s.Body
}

func (s *CreatePbcVersionRequest) SetBody(v *PbcVersionCmd) *CreatePbcVersionRequest {
	s.Body = v
	return s
}

func (s *CreatePbcVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
