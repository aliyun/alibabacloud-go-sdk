// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExternalServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExternalServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListExternalServicesResponseBody) *ListExternalServicesResponse
	GetBody() *ListExternalServicesResponseBody
}

type ListExternalServicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExternalServicesResponse) GoString() string {
	return s.String()
}

func (s *ListExternalServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExternalServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExternalServicesResponse) GetBody() *ListExternalServicesResponseBody {
	return s.Body
}

func (s *ListExternalServicesResponse) SetHeaders(v map[string]*string) *ListExternalServicesResponse {
	s.Headers = v
	return s
}

func (s *ListExternalServicesResponse) SetStatusCode(v int32) *ListExternalServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalServicesResponse) SetBody(v *ListExternalServicesResponseBody) *ListExternalServicesResponse {
	s.Body = v
	return s
}

func (s *ListExternalServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
