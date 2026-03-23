// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickStaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KickStaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KickStaResponse
	GetStatusCode() *int32
	SetBody(v *KickStaResponseBody) *KickStaResponse
	GetBody() *KickStaResponseBody
}

type KickStaResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KickStaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KickStaResponse) String() string {
	return dara.Prettify(s)
}

func (s KickStaResponse) GoString() string {
	return s.String()
}

func (s *KickStaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KickStaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KickStaResponse) GetBody() *KickStaResponseBody {
	return s.Body
}

func (s *KickStaResponse) SetHeaders(v map[string]*string) *KickStaResponse {
	s.Headers = v
	return s
}

func (s *KickStaResponse) SetStatusCode(v int32) *KickStaResponse {
	s.StatusCode = &v
	return s
}

func (s *KickStaResponse) SetBody(v *KickStaResponseBody) *KickStaResponse {
	s.Body = v
	return s
}

func (s *KickStaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
