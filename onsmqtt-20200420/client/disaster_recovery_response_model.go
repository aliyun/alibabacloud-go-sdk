// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisasterRecoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisasterRecoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisasterRecoveryResponse
	GetStatusCode() *int32
	SetBody(v *DisasterRecoveryResponseBody) *DisasterRecoveryResponse
	GetBody() *DisasterRecoveryResponseBody
}

type DisasterRecoveryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisasterRecoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisasterRecoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DisasterRecoveryResponse) GoString() string {
	return s.String()
}

func (s *DisasterRecoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisasterRecoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisasterRecoveryResponse) GetBody() *DisasterRecoveryResponseBody {
	return s.Body
}

func (s *DisasterRecoveryResponse) SetHeaders(v map[string]*string) *DisasterRecoveryResponse {
	s.Headers = v
	return s
}

func (s *DisasterRecoveryResponse) SetStatusCode(v int32) *DisasterRecoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DisasterRecoveryResponse) SetBody(v *DisasterRecoveryResponseBody) *DisasterRecoveryResponse {
	s.Body = v
	return s
}

func (s *DisasterRecoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
