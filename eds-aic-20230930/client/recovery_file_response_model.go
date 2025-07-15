// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoveryFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoveryFileResponse
	GetStatusCode() *int32
	SetBody(v *RecoveryFileResponseBody) *RecoveryFileResponse
	GetBody() *RecoveryFileResponseBody
}

type RecoveryFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoveryFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoveryFileResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoveryFileResponse) GoString() string {
	return s.String()
}

func (s *RecoveryFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoveryFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoveryFileResponse) GetBody() *RecoveryFileResponseBody {
	return s.Body
}

func (s *RecoveryFileResponse) SetHeaders(v map[string]*string) *RecoveryFileResponse {
	s.Headers = v
	return s
}

func (s *RecoveryFileResponse) SetStatusCode(v int32) *RecoveryFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryFileResponse) SetBody(v *RecoveryFileResponseBody) *RecoveryFileResponse {
	s.Body = v
	return s
}

func (s *RecoveryFileResponse) Validate() error {
	return dara.Validate(s)
}
