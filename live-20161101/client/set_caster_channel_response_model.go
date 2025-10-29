// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCasterChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCasterChannelResponse
	GetStatusCode() *int32
	SetBody(v *SetCasterChannelResponseBody) *SetCasterChannelResponse
	GetBody() *SetCasterChannelResponseBody
}

type SetCasterChannelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCasterChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCasterChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCasterChannelResponse) GoString() string {
	return s.String()
}

func (s *SetCasterChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCasterChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCasterChannelResponse) GetBody() *SetCasterChannelResponseBody {
	return s.Body
}

func (s *SetCasterChannelResponse) SetHeaders(v map[string]*string) *SetCasterChannelResponse {
	s.Headers = v
	return s
}

func (s *SetCasterChannelResponse) SetStatusCode(v int32) *SetCasterChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCasterChannelResponse) SetBody(v *SetCasterChannelResponseBody) *SetCasterChannelResponse {
	s.Body = v
	return s
}

func (s *SetCasterChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
