// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickAntStaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KickAntStaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KickAntStaResponse
	GetStatusCode() *int32
	SetBody(v *KickAntStaResponseBody) *KickAntStaResponse
	GetBody() *KickAntStaResponseBody
}

type KickAntStaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KickAntStaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KickAntStaResponse) String() string {
	return dara.Prettify(s)
}

func (s KickAntStaResponse) GoString() string {
	return s.String()
}

func (s *KickAntStaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KickAntStaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KickAntStaResponse) GetBody() *KickAntStaResponseBody {
	return s.Body
}

func (s *KickAntStaResponse) SetHeaders(v map[string]*string) *KickAntStaResponse {
	s.Headers = v
	return s
}

func (s *KickAntStaResponse) SetStatusCode(v int32) *KickAntStaResponse {
	s.StatusCode = &v
	return s
}

func (s *KickAntStaResponse) SetBody(v *KickAntStaResponseBody) *KickAntStaResponse {
	s.Body = v
	return s
}

func (s *KickAntStaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
