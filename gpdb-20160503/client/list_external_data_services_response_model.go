// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalDataServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExternalDataServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExternalDataServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListExternalDataServicesResponseBody) *ListExternalDataServicesResponse
	GetBody() *ListExternalDataServicesResponseBody
}

type ListExternalDataServicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalDataServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalDataServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataServicesResponse) GoString() string {
	return s.String()
}

func (s *ListExternalDataServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExternalDataServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExternalDataServicesResponse) GetBody() *ListExternalDataServicesResponseBody {
	return s.Body
}

func (s *ListExternalDataServicesResponse) SetHeaders(v map[string]*string) *ListExternalDataServicesResponse {
	s.Headers = v
	return s
}

func (s *ListExternalDataServicesResponse) SetStatusCode(v int32) *ListExternalDataServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalDataServicesResponse) SetBody(v *ListExternalDataServicesResponseBody) *ListExternalDataServicesResponse {
	s.Body = v
	return s
}

func (s *ListExternalDataServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
