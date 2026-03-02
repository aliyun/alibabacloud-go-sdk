// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcSchemaCreateCmd) *CreatePbcSchemaRequest
	GetBody() *PbcSchemaCreateCmd
}

type CreatePbcSchemaRequest struct {
	Body *PbcSchemaCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcSchemaRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcSchemaRequest) GetBody() *PbcSchemaCreateCmd {
	return s.Body
}

func (s *CreatePbcSchemaRequest) SetBody(v *PbcSchemaCreateCmd) *CreatePbcSchemaRequest {
	s.Body = v
	return s
}

func (s *CreatePbcSchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
