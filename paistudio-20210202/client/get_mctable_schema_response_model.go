// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMCTableSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMCTableSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMCTableSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetMCTableSchemaResponseBody) *GetMCTableSchemaResponse
	GetBody() *GetMCTableSchemaResponseBody
}

type GetMCTableSchemaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMCTableSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMCTableSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMCTableSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMCTableSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMCTableSchemaResponse) GetBody() *GetMCTableSchemaResponseBody {
	return s.Body
}

func (s *GetMCTableSchemaResponse) SetHeaders(v map[string]*string) *GetMCTableSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetMCTableSchemaResponse) SetStatusCode(v int32) *GetMCTableSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMCTableSchemaResponse) SetBody(v *GetMCTableSchemaResponseBody) *GetMCTableSchemaResponse {
	s.Body = v
	return s
}

func (s *GetMCTableSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
