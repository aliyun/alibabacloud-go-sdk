// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGraphSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGraphSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListGraphSchemasResponseBody) *ListGraphSchemasResponse
	GetBody() *ListGraphSchemasResponseBody
}

type ListGraphSchemasResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGraphSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGraphSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGraphSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGraphSchemasResponse) GetBody() *ListGraphSchemasResponseBody {
	return s.Body
}

func (s *ListGraphSchemasResponse) SetHeaders(v map[string]*string) *ListGraphSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListGraphSchemasResponse) SetStatusCode(v int32) *ListGraphSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGraphSchemasResponse) SetBody(v *ListGraphSchemasResponseBody) *ListGraphSchemasResponse {
	s.Body = v
	return s
}

func (s *ListGraphSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
