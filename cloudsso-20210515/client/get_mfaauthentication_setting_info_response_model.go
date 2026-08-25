// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationSettingInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMFAAuthenticationSettingInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMFAAuthenticationSettingInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMFAAuthenticationSettingInfoResponseBody) *GetMFAAuthenticationSettingInfoResponse
	GetBody() *GetMFAAuthenticationSettingInfoResponseBody
}

type GetMFAAuthenticationSettingInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMFAAuthenticationSettingInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMFAAuthenticationSettingInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationSettingInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationSettingInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMFAAuthenticationSettingInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMFAAuthenticationSettingInfoResponse) GetBody() *GetMFAAuthenticationSettingInfoResponseBody {
	return s.Body
}

func (s *GetMFAAuthenticationSettingInfoResponse) SetHeaders(v map[string]*string) *GetMFAAuthenticationSettingInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponse) SetStatusCode(v int32) *GetMFAAuthenticationSettingInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponse) SetBody(v *GetMFAAuthenticationSettingInfoResponseBody) *GetMFAAuthenticationSettingInfoResponse {
	s.Body = v
	return s
}

func (s *GetMFAAuthenticationSettingInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
