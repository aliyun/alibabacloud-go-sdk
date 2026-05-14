// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApprovePolarClawDevicePairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApprovePolarClawDevicePairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApprovePolarClawDevicePairResponse
	GetStatusCode() *int32
	SetBody(v *ApprovePolarClawDevicePairResponseBody) *ApprovePolarClawDevicePairResponse
	GetBody() *ApprovePolarClawDevicePairResponseBody
}

type ApprovePolarClawDevicePairResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApprovePolarClawDevicePairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApprovePolarClawDevicePairResponse) String() string {
	return dara.Prettify(s)
}

func (s ApprovePolarClawDevicePairResponse) GoString() string {
	return s.String()
}

func (s *ApprovePolarClawDevicePairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApprovePolarClawDevicePairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApprovePolarClawDevicePairResponse) GetBody() *ApprovePolarClawDevicePairResponseBody {
	return s.Body
}

func (s *ApprovePolarClawDevicePairResponse) SetHeaders(v map[string]*string) *ApprovePolarClawDevicePairResponse {
	s.Headers = v
	return s
}

func (s *ApprovePolarClawDevicePairResponse) SetStatusCode(v int32) *ApprovePolarClawDevicePairResponse {
	s.StatusCode = &v
	return s
}

func (s *ApprovePolarClawDevicePairResponse) SetBody(v *ApprovePolarClawDevicePairResponseBody) *ApprovePolarClawDevicePairResponse {
	s.Body = v
	return s
}

func (s *ApprovePolarClawDevicePairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
