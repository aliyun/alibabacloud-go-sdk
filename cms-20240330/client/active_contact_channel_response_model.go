// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveContactChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveContactChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveContactChannelResponse
	GetStatusCode() *int32
	SetBody(v *ActiveContactChannelResponseBody) *ActiveContactChannelResponse
	GetBody() *ActiveContactChannelResponseBody
}

type ActiveContactChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveContactChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveContactChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveContactChannelResponse) GoString() string {
	return s.String()
}

func (s *ActiveContactChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveContactChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveContactChannelResponse) GetBody() *ActiveContactChannelResponseBody {
	return s.Body
}

func (s *ActiveContactChannelResponse) SetHeaders(v map[string]*string) *ActiveContactChannelResponse {
	s.Headers = v
	return s
}

func (s *ActiveContactChannelResponse) SetStatusCode(v int32) *ActiveContactChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveContactChannelResponse) SetBody(v *ActiveContactChannelResponseBody) *ActiveContactChannelResponse {
	s.Body = v
	return s
}

func (s *ActiveContactChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
