// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcApiSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcApiSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcApiSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PbcApiSchema) *CreatePbcApiSchemaResponse
	GetBody() *PbcApiSchema
}

type CreatePbcApiSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcApiSchema      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcApiSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcApiSchemaResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcApiSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcApiSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcApiSchemaResponse) GetBody() *PbcApiSchema {
	return s.Body
}

func (s *CreatePbcApiSchemaResponse) SetHeaders(v map[string]*string) *CreatePbcApiSchemaResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcApiSchemaResponse) SetStatusCode(v int32) *CreatePbcApiSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcApiSchemaResponse) SetBody(v *PbcApiSchema) *CreatePbcApiSchemaResponse {
	s.Body = v
	return s
}

func (s *CreatePbcApiSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
