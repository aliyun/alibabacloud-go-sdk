// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PbcSchema) *GetPbcSchemaResponse
	GetBody() *PbcSchema
}

type GetPbcSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcSchema         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetPbcSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcSchemaResponse) GetBody() *PbcSchema {
	return s.Body
}

func (s *GetPbcSchemaResponse) SetHeaders(v map[string]*string) *GetPbcSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetPbcSchemaResponse) SetStatusCode(v int32) *GetPbcSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcSchemaResponse) SetBody(v *PbcSchema) *GetPbcSchemaResponse {
	s.Body = v
	return s
}

func (s *GetPbcSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
