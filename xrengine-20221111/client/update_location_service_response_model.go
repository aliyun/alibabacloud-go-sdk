// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLocationServiceResponseBody) *UpdateLocationServiceResponse
	GetBody() *UpdateLocationServiceResponseBody
}

type UpdateLocationServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLocationServiceResponse) GetBody() *UpdateLocationServiceResponseBody {
	return s.Body
}

func (s *UpdateLocationServiceResponse) SetHeaders(v map[string]*string) *UpdateLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateLocationServiceResponse) SetStatusCode(v int32) *UpdateLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLocationServiceResponse) SetBody(v *UpdateLocationServiceResponseBody) *UpdateLocationServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateLocationServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
