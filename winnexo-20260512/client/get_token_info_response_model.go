// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenInfoResponseBody) *GetTokenInfoResponse
	GetBody() *GetTokenInfoResponseBody
}

type GetTokenInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTokenInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenInfoResponse) GetBody() *GetTokenInfoResponseBody {
	return s.Body
}

func (s *GetTokenInfoResponse) SetHeaders(v map[string]*string) *GetTokenInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTokenInfoResponse) SetStatusCode(v int32) *GetTokenInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenInfoResponse) SetBody(v *GetTokenInfoResponseBody) *GetTokenInfoResponse {
	s.Body = v
	return s
}

func (s *GetTokenInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
