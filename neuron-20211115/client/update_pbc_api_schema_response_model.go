// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcApiSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePbcApiSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePbcApiSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PbcApiSchema) *UpdatePbcApiSchemaResponse
	GetBody() *PbcApiSchema
}

type UpdatePbcApiSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcApiSchema      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcApiSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcApiSchemaResponse) GoString() string {
	return s.String()
}

func (s *UpdatePbcApiSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePbcApiSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePbcApiSchemaResponse) GetBody() *PbcApiSchema {
	return s.Body
}

func (s *UpdatePbcApiSchemaResponse) SetHeaders(v map[string]*string) *UpdatePbcApiSchemaResponse {
	s.Headers = v
	return s
}

func (s *UpdatePbcApiSchemaResponse) SetStatusCode(v int32) *UpdatePbcApiSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePbcApiSchemaResponse) SetBody(v *PbcApiSchema) *UpdatePbcApiSchemaResponse {
	s.Body = v
	return s
}

func (s *UpdatePbcApiSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
