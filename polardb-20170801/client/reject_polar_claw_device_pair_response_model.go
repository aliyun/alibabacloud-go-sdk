// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectPolarClawDevicePairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectPolarClawDevicePairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectPolarClawDevicePairResponse
	GetStatusCode() *int32
	SetBody(v *RejectPolarClawDevicePairResponseBody) *RejectPolarClawDevicePairResponse
	GetBody() *RejectPolarClawDevicePairResponseBody
}

type RejectPolarClawDevicePairResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectPolarClawDevicePairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectPolarClawDevicePairResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectPolarClawDevicePairResponse) GoString() string {
	return s.String()
}

func (s *RejectPolarClawDevicePairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectPolarClawDevicePairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectPolarClawDevicePairResponse) GetBody() *RejectPolarClawDevicePairResponseBody {
	return s.Body
}

func (s *RejectPolarClawDevicePairResponse) SetHeaders(v map[string]*string) *RejectPolarClawDevicePairResponse {
	s.Headers = v
	return s
}

func (s *RejectPolarClawDevicePairResponse) SetStatusCode(v int32) *RejectPolarClawDevicePairResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectPolarClawDevicePairResponse) SetBody(v *RejectPolarClawDevicePairResponseBody) *RejectPolarClawDevicePairResponse {
	s.Body = v
	return s
}

func (s *RejectPolarClawDevicePairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
