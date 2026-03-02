// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PbcSchema) *CreatePbcSchemaResponse
	GetBody() *PbcSchema
}

type CreatePbcSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcSchema         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcSchemaResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcSchemaResponse) GetBody() *PbcSchema {
	return s.Body
}

func (s *CreatePbcSchemaResponse) SetHeaders(v map[string]*string) *CreatePbcSchemaResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcSchemaResponse) SetStatusCode(v int32) *CreatePbcSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcSchemaResponse) SetBody(v *PbcSchema) *CreatePbcSchemaResponse {
	s.Body = v
	return s
}

func (s *CreatePbcSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
