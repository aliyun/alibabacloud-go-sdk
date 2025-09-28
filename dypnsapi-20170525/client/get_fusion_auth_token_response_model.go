// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFusionAuthTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFusionAuthTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFusionAuthTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetFusionAuthTokenResponseBody) *GetFusionAuthTokenResponse
	GetBody() *GetFusionAuthTokenResponseBody
}

type GetFusionAuthTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFusionAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFusionAuthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFusionAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetFusionAuthTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFusionAuthTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFusionAuthTokenResponse) GetBody() *GetFusionAuthTokenResponseBody {
	return s.Body
}

func (s *GetFusionAuthTokenResponse) SetHeaders(v map[string]*string) *GetFusionAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetFusionAuthTokenResponse) SetStatusCode(v int32) *GetFusionAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFusionAuthTokenResponse) SetBody(v *GetFusionAuthTokenResponseBody) *GetFusionAuthTokenResponse {
	s.Body = v
	return s
}

func (s *GetFusionAuthTokenResponse) Validate() error {
	return dara.Validate(s)
}
