// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRtcTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRtcTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRtcTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetRtcTokenResponseBody) *GetRtcTokenResponse
	GetBody() *GetRtcTokenResponseBody
}

type GetRtcTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRtcTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRtcTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRtcTokenResponse) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRtcTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRtcTokenResponse) GetBody() *GetRtcTokenResponseBody {
	return s.Body
}

func (s *GetRtcTokenResponse) SetHeaders(v map[string]*string) *GetRtcTokenResponse {
	s.Headers = v
	return s
}

func (s *GetRtcTokenResponse) SetStatusCode(v int32) *GetRtcTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRtcTokenResponse) SetBody(v *GetRtcTokenResponseBody) *GetRtcTokenResponse {
	s.Body = v
	return s
}

func (s *GetRtcTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
