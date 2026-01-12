// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *QueryLocationServiceResponseBody) *QueryLocationServiceResponse
	GetBody() *QueryLocationServiceResponseBody
}

type QueryLocationServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLocationServiceResponse) GetBody() *QueryLocationServiceResponseBody {
	return s.Body
}

func (s *QueryLocationServiceResponse) SetHeaders(v map[string]*string) *QueryLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryLocationServiceResponse) SetStatusCode(v int32) *QueryLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocationServiceResponse) SetBody(v *QueryLocationServiceResponseBody) *QueryLocationServiceResponse {
	s.Body = v
	return s
}

func (s *QueryLocationServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
