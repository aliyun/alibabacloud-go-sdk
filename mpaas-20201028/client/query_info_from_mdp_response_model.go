// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInfoFromMdpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInfoFromMdpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInfoFromMdpResponse
	GetStatusCode() *int32
	SetBody(v *QueryInfoFromMdpResponseBody) *QueryInfoFromMdpResponse
	GetBody() *QueryInfoFromMdpResponseBody
}

type QueryInfoFromMdpResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInfoFromMdpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInfoFromMdpResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInfoFromMdpResponse) GoString() string {
	return s.String()
}

func (s *QueryInfoFromMdpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInfoFromMdpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInfoFromMdpResponse) GetBody() *QueryInfoFromMdpResponseBody {
	return s.Body
}

func (s *QueryInfoFromMdpResponse) SetHeaders(v map[string]*string) *QueryInfoFromMdpResponse {
	s.Headers = v
	return s
}

func (s *QueryInfoFromMdpResponse) SetStatusCode(v int32) *QueryInfoFromMdpResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInfoFromMdpResponse) SetBody(v *QueryInfoFromMdpResponseBody) *QueryInfoFromMdpResponse {
	s.Body = v
	return s
}

func (s *QueryInfoFromMdpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
