// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPauseProtectionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPauseProtectionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPauseProtectionStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPauseProtectionStatusResponseBody) *ModifyPauseProtectionStatusResponse
	GetBody() *ModifyPauseProtectionStatusResponseBody
}

type ModifyPauseProtectionStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPauseProtectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPauseProtectionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPauseProtectionStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyPauseProtectionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPauseProtectionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPauseProtectionStatusResponse) GetBody() *ModifyPauseProtectionStatusResponseBody {
	return s.Body
}

func (s *ModifyPauseProtectionStatusResponse) SetHeaders(v map[string]*string) *ModifyPauseProtectionStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyPauseProtectionStatusResponse) SetStatusCode(v int32) *ModifyPauseProtectionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPauseProtectionStatusResponse) SetBody(v *ModifyPauseProtectionStatusResponseBody) *ModifyPauseProtectionStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyPauseProtectionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
