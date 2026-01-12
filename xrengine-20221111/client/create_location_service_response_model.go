// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateLocationServiceResponseBody) *CreateLocationServiceResponse
	GetBody() *CreateLocationServiceResponseBody
}

type CreateLocationServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLocationServiceResponse) GetBody() *CreateLocationServiceResponseBody {
	return s.Body
}

func (s *CreateLocationServiceResponse) SetHeaders(v map[string]*string) *CreateLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateLocationServiceResponse) SetStatusCode(v int32) *CreateLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocationServiceResponse) SetBody(v *CreateLocationServiceResponseBody) *CreateLocationServiceResponse {
	s.Body = v
	return s
}

func (s *CreateLocationServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
