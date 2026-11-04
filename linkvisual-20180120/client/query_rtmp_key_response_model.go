// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRtmpKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRtmpKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRtmpKeyResponse
	GetStatusCode() *int32
	SetBody(v *QueryRtmpKeyResponseBody) *QueryRtmpKeyResponse
	GetBody() *QueryRtmpKeyResponseBody
}

type QueryRtmpKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRtmpKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRtmpKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRtmpKeyResponse) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRtmpKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRtmpKeyResponse) GetBody() *QueryRtmpKeyResponseBody {
	return s.Body
}

func (s *QueryRtmpKeyResponse) SetHeaders(v map[string]*string) *QueryRtmpKeyResponse {
	s.Headers = v
	return s
}

func (s *QueryRtmpKeyResponse) SetStatusCode(v int32) *QueryRtmpKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRtmpKeyResponse) SetBody(v *QueryRtmpKeyResponseBody) *QueryRtmpKeyResponse {
	s.Body = v
	return s
}

func (s *QueryRtmpKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
