// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarClawChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarClawChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarClawChannelResponseBody) *CreatePolarClawChannelResponse
	GetBody() *CreatePolarClawChannelResponseBody
}

type CreatePolarClawChannelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarClawChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarClawChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawChannelResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarClawChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarClawChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarClawChannelResponse) GetBody() *CreatePolarClawChannelResponseBody {
	return s.Body
}

func (s *CreatePolarClawChannelResponse) SetHeaders(v map[string]*string) *CreatePolarClawChannelResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarClawChannelResponse) SetStatusCode(v int32) *CreatePolarClawChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarClawChannelResponse) SetBody(v *CreatePolarClawChannelResponseBody) *CreatePolarClawChannelResponse {
	s.Body = v
	return s
}

func (s *CreatePolarClawChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
