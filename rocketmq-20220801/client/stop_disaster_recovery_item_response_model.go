// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDisasterRecoveryItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDisasterRecoveryItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDisasterRecoveryItemResponse
	GetStatusCode() *int32
	SetBody(v *StopDisasterRecoveryItemResponseBody) *StopDisasterRecoveryItemResponse
	GetBody() *StopDisasterRecoveryItemResponseBody
}

type StopDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDisasterRecoveryItemResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *StopDisasterRecoveryItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDisasterRecoveryItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDisasterRecoveryItemResponse) GetBody() *StopDisasterRecoveryItemResponseBody {
	return s.Body
}

func (s *StopDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *StopDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *StopDisasterRecoveryItemResponse) SetStatusCode(v int32) *StopDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDisasterRecoveryItemResponse) SetBody(v *StopDisasterRecoveryItemResponseBody) *StopDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

func (s *StopDisasterRecoveryItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
