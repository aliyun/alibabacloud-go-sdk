// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlatformUserInfoForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlatformUserInfoForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlatformUserInfoForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetPlatformUserInfoForPartnerResponseBody) *GetPlatformUserInfoForPartnerResponse
	GetBody() *GetPlatformUserInfoForPartnerResponseBody
}

type GetPlatformUserInfoForPartnerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlatformUserInfoForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlatformUserInfoForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlatformUserInfoForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetPlatformUserInfoForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlatformUserInfoForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlatformUserInfoForPartnerResponse) GetBody() *GetPlatformUserInfoForPartnerResponseBody {
	return s.Body
}

func (s *GetPlatformUserInfoForPartnerResponse) SetHeaders(v map[string]*string) *GetPlatformUserInfoForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponse) SetStatusCode(v int32) *GetPlatformUserInfoForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponse) SetBody(v *GetPlatformUserInfoForPartnerResponseBody) *GetPlatformUserInfoForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetPlatformUserInfoForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
