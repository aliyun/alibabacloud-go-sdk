// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenServiceResponseBody) *OpenServiceResponse
	GetBody() *OpenServiceResponseBody
}

type OpenServiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenServiceResponse) GetBody() *OpenServiceResponseBody {
	return s.Body
}

func (s *OpenServiceResponse) SetHeaders(v map[string]*string) *OpenServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenServiceResponse) SetStatusCode(v int32) *OpenServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenServiceResponse) SetBody(v *OpenServiceResponseBody) *OpenServiceResponse {
	s.Body = v
	return s
}

func (s *OpenServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
