// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourceSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceResourceSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceResourceSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceResourceSchemasResponseBody) *ListInstanceResourceSchemasResponse
	GetBody() *ListInstanceResourceSchemasResponseBody
}

type ListInstanceResourceSchemasResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResourceSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResourceSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceResourceSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceResourceSchemasResponse) GetBody() *ListInstanceResourceSchemasResponseBody {
	return s.Body
}

func (s *ListInstanceResourceSchemasResponse) SetHeaders(v map[string]*string) *ListInstanceResourceSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResourceSchemasResponse) SetStatusCode(v int32) *ListInstanceResourceSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResourceSchemasResponse) SetBody(v *ListInstanceResourceSchemasResponseBody) *ListInstanceResourceSchemasResponse {
	s.Body = v
	return s
}

func (s *ListInstanceResourceSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
