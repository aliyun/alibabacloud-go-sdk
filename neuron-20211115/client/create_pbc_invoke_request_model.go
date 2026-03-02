// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcInvokeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcInvokeCreateCmd) *CreatePbcInvokeRequest
	GetBody() *PbcInvokeCreateCmd
}

type CreatePbcInvokeRequest struct {
	Body *PbcInvokeCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcInvokeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcInvokeRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcInvokeRequest) GetBody() *PbcInvokeCreateCmd {
	return s.Body
}

func (s *CreatePbcInvokeRequest) SetBody(v *PbcInvokeCreateCmd) *CreatePbcInvokeRequest {
	s.Body = v
	return s
}

func (s *CreatePbcInvokeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
