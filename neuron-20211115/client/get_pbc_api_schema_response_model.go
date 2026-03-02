// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcApiSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcApiSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcApiSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PbcApiSchema) *GetPbcApiSchemaResponse
	GetBody() *PbcApiSchema
}

type GetPbcApiSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcApiSchema      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcApiSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcApiSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetPbcApiSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcApiSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcApiSchemaResponse) GetBody() *PbcApiSchema {
	return s.Body
}

func (s *GetPbcApiSchemaResponse) SetHeaders(v map[string]*string) *GetPbcApiSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetPbcApiSchemaResponse) SetStatusCode(v int32) *GetPbcApiSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcApiSchemaResponse) SetBody(v *PbcApiSchema) *GetPbcApiSchemaResponse {
	s.Body = v
	return s
}

func (s *GetPbcApiSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
