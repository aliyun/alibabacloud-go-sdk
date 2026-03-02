// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcApiSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PbcApiSchemaUpdateCmd) *UpdatePbcApiSchemaRequest
	GetBody() *PbcApiSchemaUpdateCmd
}

type UpdatePbcApiSchemaRequest struct {
	Body *PbcApiSchemaUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcApiSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcApiSchemaRequest) GoString() string {
	return s.String()
}

func (s *UpdatePbcApiSchemaRequest) GetBody() *PbcApiSchemaUpdateCmd {
	return s.Body
}

func (s *UpdatePbcApiSchemaRequest) SetBody(v *PbcApiSchemaUpdateCmd) *UpdatePbcApiSchemaRequest {
	s.Body = v
	return s
}

func (s *UpdatePbcApiSchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
