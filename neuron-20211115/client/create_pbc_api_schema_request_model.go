// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcApiSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcApiSchemaCreateCmd) *CreatePbcApiSchemaRequest
	GetBody() *PbcApiSchemaCreateCmd
}

type CreatePbcApiSchemaRequest struct {
	Body *PbcApiSchemaCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcApiSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcApiSchemaRequest) GoString() string {
	return s.String()
}

func (s *CreatePbcApiSchemaRequest) GetBody() *PbcApiSchemaCreateCmd {
	return s.Body
}

func (s *CreatePbcApiSchemaRequest) SetBody(v *PbcApiSchemaCreateCmd) *CreatePbcApiSchemaRequest {
	s.Body = v
	return s
}

func (s *CreatePbcApiSchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
