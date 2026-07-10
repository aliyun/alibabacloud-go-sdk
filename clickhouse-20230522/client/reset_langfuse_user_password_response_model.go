// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLangfuseUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetLangfuseUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetLangfuseUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetLangfuseUserPasswordResponseBody) *ResetLangfuseUserPasswordResponse
	GetBody() *ResetLangfuseUserPasswordResponseBody
}

type ResetLangfuseUserPasswordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetLangfuseUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetLangfuseUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetLangfuseUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetLangfuseUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetLangfuseUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetLangfuseUserPasswordResponse) GetBody() *ResetLangfuseUserPasswordResponseBody {
	return s.Body
}

func (s *ResetLangfuseUserPasswordResponse) SetHeaders(v map[string]*string) *ResetLangfuseUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetLangfuseUserPasswordResponse) SetStatusCode(v int32) *ResetLangfuseUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetLangfuseUserPasswordResponse) SetBody(v *ResetLangfuseUserPasswordResponseBody) *ResetLangfuseUserPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetLangfuseUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
