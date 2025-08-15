// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDisasterRecoveryItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDisasterRecoveryItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDisasterRecoveryItemResponse
	GetStatusCode() *int32
	SetBody(v *StartDisasterRecoveryItemResponseBody) *StartDisasterRecoveryItemResponse
	GetBody() *StartDisasterRecoveryItemResponseBody
}

type StartDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDisasterRecoveryItemResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *StartDisasterRecoveryItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDisasterRecoveryItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDisasterRecoveryItemResponse) GetBody() *StartDisasterRecoveryItemResponseBody {
	return s.Body
}

func (s *StartDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *StartDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *StartDisasterRecoveryItemResponse) SetStatusCode(v int32) *StartDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDisasterRecoveryItemResponse) SetBody(v *StartDisasterRecoveryItemResponseBody) *StartDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

func (s *StartDisasterRecoveryItemResponse) Validate() error {
	return dara.Validate(s)
}
