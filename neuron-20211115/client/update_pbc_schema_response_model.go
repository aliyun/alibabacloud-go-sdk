// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePbcSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePbcSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PbcSchema) *UpdatePbcSchemaResponse
	GetBody() *PbcSchema
}

type UpdatePbcSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcSchema         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcSchemaResponse) GoString() string {
	return s.String()
}

func (s *UpdatePbcSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePbcSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePbcSchemaResponse) GetBody() *PbcSchema {
	return s.Body
}

func (s *UpdatePbcSchemaResponse) SetHeaders(v map[string]*string) *UpdatePbcSchemaResponse {
	s.Headers = v
	return s
}

func (s *UpdatePbcSchemaResponse) SetStatusCode(v int32) *UpdatePbcSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePbcSchemaResponse) SetBody(v *PbcSchema) *UpdatePbcSchemaResponse {
	s.Body = v
	return s
}

func (s *UpdatePbcSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
