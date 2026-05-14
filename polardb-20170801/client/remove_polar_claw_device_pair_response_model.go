// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePolarClawDevicePairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePolarClawDevicePairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePolarClawDevicePairResponse
	GetStatusCode() *int32
	SetBody(v *RemovePolarClawDevicePairResponseBody) *RemovePolarClawDevicePairResponse
	GetBody() *RemovePolarClawDevicePairResponseBody
}

type RemovePolarClawDevicePairResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePolarClawDevicePairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePolarClawDevicePairResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePolarClawDevicePairResponse) GoString() string {
	return s.String()
}

func (s *RemovePolarClawDevicePairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePolarClawDevicePairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePolarClawDevicePairResponse) GetBody() *RemovePolarClawDevicePairResponseBody {
	return s.Body
}

func (s *RemovePolarClawDevicePairResponse) SetHeaders(v map[string]*string) *RemovePolarClawDevicePairResponse {
	s.Headers = v
	return s
}

func (s *RemovePolarClawDevicePairResponse) SetStatusCode(v int32) *RemovePolarClawDevicePairResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePolarClawDevicePairResponse) SetBody(v *RemovePolarClawDevicePairResponseBody) *RemovePolarClawDevicePairResponse {
	s.Body = v
	return s
}

func (s *RemovePolarClawDevicePairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
