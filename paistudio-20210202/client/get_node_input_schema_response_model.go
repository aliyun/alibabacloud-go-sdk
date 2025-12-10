// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeInputSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeInputSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeInputSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeInputSchemaResponseBody) *GetNodeInputSchemaResponse
	GetBody() *GetNodeInputSchemaResponseBody
}

type GetNodeInputSchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeInputSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeInputSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInputSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetNodeInputSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeInputSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeInputSchemaResponse) GetBody() *GetNodeInputSchemaResponseBody {
	return s.Body
}

func (s *GetNodeInputSchemaResponse) SetHeaders(v map[string]*string) *GetNodeInputSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetNodeInputSchemaResponse) SetStatusCode(v int32) *GetNodeInputSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeInputSchemaResponse) SetBody(v *GetNodeInputSchemaResponseBody) *GetNodeInputSchemaResponse {
	s.Body = v
	return s
}

func (s *GetNodeInputSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
