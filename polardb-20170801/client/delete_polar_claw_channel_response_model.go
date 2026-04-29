// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarClawChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarClawChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarClawChannelResponseBody) *DeletePolarClawChannelResponse
	GetBody() *DeletePolarClawChannelResponseBody
}

type DeletePolarClawChannelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarClawChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarClawChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawChannelResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarClawChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarClawChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarClawChannelResponse) GetBody() *DeletePolarClawChannelResponseBody {
	return s.Body
}

func (s *DeletePolarClawChannelResponse) SetHeaders(v map[string]*string) *DeletePolarClawChannelResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarClawChannelResponse) SetStatusCode(v int32) *DeletePolarClawChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarClawChannelResponse) SetBody(v *DeletePolarClawChannelResponseBody) *DeletePolarClawChannelResponse {
	s.Body = v
	return s
}

func (s *DeletePolarClawChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
