// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppDonwloadUrlInMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAppDonwloadUrlInMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAppDonwloadUrlInMsaResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAppDonwloadUrlInMsaResponseBody) *GetUserAppDonwloadUrlInMsaResponse
	GetBody() *GetUserAppDonwloadUrlInMsaResponseBody
}

type GetUserAppDonwloadUrlInMsaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAppDonwloadUrlInMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAppDonwloadUrlInMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppDonwloadUrlInMsaResponse) GoString() string {
	return s.String()
}

func (s *GetUserAppDonwloadUrlInMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAppDonwloadUrlInMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAppDonwloadUrlInMsaResponse) GetBody() *GetUserAppDonwloadUrlInMsaResponseBody {
	return s.Body
}

func (s *GetUserAppDonwloadUrlInMsaResponse) SetHeaders(v map[string]*string) *GetUserAppDonwloadUrlInMsaResponse {
	s.Headers = v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponse) SetStatusCode(v int32) *GetUserAppDonwloadUrlInMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponse) SetBody(v *GetUserAppDonwloadUrlInMsaResponseBody) *GetUserAppDonwloadUrlInMsaResponse {
	s.Body = v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponse) Validate() error {
	return dara.Validate(s)
}
