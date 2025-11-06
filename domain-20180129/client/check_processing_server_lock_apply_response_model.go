// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProcessingServerLockApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckProcessingServerLockApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckProcessingServerLockApplyResponse
	GetStatusCode() *int32
	SetBody(v *CheckProcessingServerLockApplyResponseBody) *CheckProcessingServerLockApplyResponse
	GetBody() *CheckProcessingServerLockApplyResponseBody
}

type CheckProcessingServerLockApplyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckProcessingServerLockApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckProcessingServerLockApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckProcessingServerLockApplyResponse) GoString() string {
	return s.String()
}

func (s *CheckProcessingServerLockApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckProcessingServerLockApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckProcessingServerLockApplyResponse) GetBody() *CheckProcessingServerLockApplyResponseBody {
	return s.Body
}

func (s *CheckProcessingServerLockApplyResponse) SetHeaders(v map[string]*string) *CheckProcessingServerLockApplyResponse {
	s.Headers = v
	return s
}

func (s *CheckProcessingServerLockApplyResponse) SetStatusCode(v int32) *CheckProcessingServerLockApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckProcessingServerLockApplyResponse) SetBody(v *CheckProcessingServerLockApplyResponseBody) *CheckProcessingServerLockApplyResponse {
	s.Body = v
	return s
}

func (s *CheckProcessingServerLockApplyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
