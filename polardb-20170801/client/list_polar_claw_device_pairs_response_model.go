// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawDevicePairsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolarClawDevicePairsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolarClawDevicePairsResponse
	GetStatusCode() *int32
	SetBody(v *ListPolarClawDevicePairsResponseBody) *ListPolarClawDevicePairsResponse
	GetBody() *ListPolarClawDevicePairsResponseBody
}

type ListPolarClawDevicePairsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolarClawDevicePairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolarClawDevicePairsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawDevicePairsResponse) GoString() string {
	return s.String()
}

func (s *ListPolarClawDevicePairsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolarClawDevicePairsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolarClawDevicePairsResponse) GetBody() *ListPolarClawDevicePairsResponseBody {
	return s.Body
}

func (s *ListPolarClawDevicePairsResponse) SetHeaders(v map[string]*string) *ListPolarClawDevicePairsResponse {
	s.Headers = v
	return s
}

func (s *ListPolarClawDevicePairsResponse) SetStatusCode(v int32) *ListPolarClawDevicePairsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolarClawDevicePairsResponse) SetBody(v *ListPolarClawDevicePairsResponseBody) *ListPolarClawDevicePairsResponse {
	s.Body = v
	return s
}

func (s *ListPolarClawDevicePairsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
