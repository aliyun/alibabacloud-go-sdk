// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebshellTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWebshellTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWebshellTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetWebshellTokenResponseBody) *GetWebshellTokenResponse
	GetBody() *GetWebshellTokenResponseBody
}

type GetWebshellTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebshellTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebshellTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWebshellTokenResponse) GoString() string {
	return s.String()
}

func (s *GetWebshellTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWebshellTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWebshellTokenResponse) GetBody() *GetWebshellTokenResponseBody {
	return s.Body
}

func (s *GetWebshellTokenResponse) SetHeaders(v map[string]*string) *GetWebshellTokenResponse {
	s.Headers = v
	return s
}

func (s *GetWebshellTokenResponse) SetStatusCode(v int32) *GetWebshellTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebshellTokenResponse) SetBody(v *GetWebshellTokenResponseBody) *GetWebshellTokenResponse {
	s.Body = v
	return s
}

func (s *GetWebshellTokenResponse) Validate() error {
	return dara.Validate(s)
}
