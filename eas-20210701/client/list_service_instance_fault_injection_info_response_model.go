// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceFaultInjectionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceFaultInjectionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceFaultInjectionInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceFaultInjectionInfoResponseBody) *ListServiceInstanceFaultInjectionInfoResponse
	GetBody() *ListServiceInstanceFaultInjectionInfoResponseBody
}

type ListServiceInstanceFaultInjectionInfoResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceFaultInjectionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceFaultInjectionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceFaultInjectionInfoResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) GetBody() *ListServiceInstanceFaultInjectionInfoResponseBody {
	return s.Body
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) SetHeaders(v map[string]*string) *ListServiceInstanceFaultInjectionInfoResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) SetStatusCode(v int32) *ListServiceInstanceFaultInjectionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) SetBody(v *ListServiceInstanceFaultInjectionInfoResponseBody) *ListServiceInstanceFaultInjectionInfoResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceFaultInjectionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
