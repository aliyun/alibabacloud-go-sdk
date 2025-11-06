// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserTmpIdentityForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserTmpIdentityForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserTmpIdentityForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetUserTmpIdentityForPartnerResponseBody) *GetUserTmpIdentityForPartnerResponse
	GetBody() *GetUserTmpIdentityForPartnerResponseBody
}

type GetUserTmpIdentityForPartnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserTmpIdentityForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserTmpIdentityForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserTmpIdentityForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetUserTmpIdentityForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserTmpIdentityForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserTmpIdentityForPartnerResponse) GetBody() *GetUserTmpIdentityForPartnerResponseBody {
	return s.Body
}

func (s *GetUserTmpIdentityForPartnerResponse) SetHeaders(v map[string]*string) *GetUserTmpIdentityForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponse) SetStatusCode(v int32) *GetUserTmpIdentityForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponse) SetBody(v *GetUserTmpIdentityForPartnerResponseBody) *GetUserTmpIdentityForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
