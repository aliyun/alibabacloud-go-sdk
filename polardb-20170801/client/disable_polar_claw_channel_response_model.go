// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisablePolarClawChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisablePolarClawChannelResponse
	GetStatusCode() *int32
	SetBody(v *DisablePolarClawChannelResponseBody) *DisablePolarClawChannelResponse
	GetBody() *DisablePolarClawChannelResponseBody
}

type DisablePolarClawChannelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisablePolarClawChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisablePolarClawChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawChannelResponse) GoString() string {
	return s.String()
}

func (s *DisablePolarClawChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisablePolarClawChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisablePolarClawChannelResponse) GetBody() *DisablePolarClawChannelResponseBody {
	return s.Body
}

func (s *DisablePolarClawChannelResponse) SetHeaders(v map[string]*string) *DisablePolarClawChannelResponse {
	s.Headers = v
	return s
}

func (s *DisablePolarClawChannelResponse) SetStatusCode(v int32) *DisablePolarClawChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DisablePolarClawChannelResponse) SetBody(v *DisablePolarClawChannelResponseBody) *DisablePolarClawChannelResponse {
	s.Body = v
	return s
}

func (s *DisablePolarClawChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
