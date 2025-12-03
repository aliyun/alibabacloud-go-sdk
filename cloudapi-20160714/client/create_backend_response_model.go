// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackendResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackendResponseBody) *CreateBackendResponse
	GetBody() *CreateBackendResponseBody
}

type CreateBackendResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackendResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendResponse) GoString() string {
	return s.String()
}

func (s *CreateBackendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackendResponse) GetBody() *CreateBackendResponseBody {
	return s.Body
}

func (s *CreateBackendResponse) SetHeaders(v map[string]*string) *CreateBackendResponse {
	s.Headers = v
	return s
}

func (s *CreateBackendResponse) SetStatusCode(v int32) *CreateBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackendResponse) SetBody(v *CreateBackendResponseBody) *CreateBackendResponse {
	s.Body = v
	return s
}

func (s *CreateBackendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
