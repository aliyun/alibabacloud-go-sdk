// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryOriginImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoveryOriginImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoveryOriginImageResponse
	GetStatusCode() *int32
	SetBody(v *RecoveryOriginImageResponseBody) *RecoveryOriginImageResponse
	GetBody() *RecoveryOriginImageResponseBody
}

type RecoveryOriginImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoveryOriginImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoveryOriginImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoveryOriginImageResponse) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoveryOriginImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoveryOriginImageResponse) GetBody() *RecoveryOriginImageResponseBody {
	return s.Body
}

func (s *RecoveryOriginImageResponse) SetHeaders(v map[string]*string) *RecoveryOriginImageResponse {
	s.Headers = v
	return s
}

func (s *RecoveryOriginImageResponse) SetStatusCode(v int32) *RecoveryOriginImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryOriginImageResponse) SetBody(v *RecoveryOriginImageResponseBody) *RecoveryOriginImageResponse {
	s.Body = v
	return s
}

func (s *RecoveryOriginImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
