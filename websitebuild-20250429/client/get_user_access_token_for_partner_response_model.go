// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAccessTokenForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAccessTokenForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAccessTokenForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAccessTokenForPartnerResponseBody) *GetUserAccessTokenForPartnerResponse
	GetBody() *GetUserAccessTokenForPartnerResponseBody
}

type GetUserAccessTokenForPartnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAccessTokenForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAccessTokenForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAccessTokenForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetUserAccessTokenForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAccessTokenForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAccessTokenForPartnerResponse) GetBody() *GetUserAccessTokenForPartnerResponseBody {
	return s.Body
}

func (s *GetUserAccessTokenForPartnerResponse) SetHeaders(v map[string]*string) *GetUserAccessTokenForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetUserAccessTokenForPartnerResponse) SetStatusCode(v int32) *GetUserAccessTokenForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAccessTokenForPartnerResponse) SetBody(v *GetUserAccessTokenForPartnerResponseBody) *GetUserAccessTokenForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetUserAccessTokenForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
