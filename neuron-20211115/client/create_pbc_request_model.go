// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcCreateCmd) *CreatePbcRequest
	GetBody() *PbcCreateCmd
}

type CreatePbcRequest struct {
	Body *PbcCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcRequest) GetBody() *PbcCreateCmd {
	return s.Body
}

func (s *CreatePbcRequest) SetBody(v *PbcCreateCmd) *CreatePbcRequest {
	s.Body = v
	return s
}

func (s *CreatePbcRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
