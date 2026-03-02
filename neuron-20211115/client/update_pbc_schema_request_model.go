// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcSchema) *UpdatePbcSchemaRequest
	GetBody() *PbcSchema
}

type UpdatePbcSchemaRequest struct {
	// This parameter is required.
	Body *PbcSchema `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcSchemaRequest) GoString() string {
	return s.String()
}

func (s *UpdatePbcSchemaRequest) GetBody() *PbcSchema {
	return s.Body
}

func (s *UpdatePbcSchemaRequest) SetBody(v *PbcSchema) *UpdatePbcSchemaRequest {
	s.Body = v
	return s
}

func (s *UpdatePbcSchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
