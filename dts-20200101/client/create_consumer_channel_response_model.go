// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConsumerChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConsumerChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateConsumerChannelResponseBody) *CreateConsumerChannelResponse
	GetBody() *CreateConsumerChannelResponseBody
}

type CreateConsumerChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConsumerChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConsumerChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConsumerChannelResponse) GetBody() *CreateConsumerChannelResponseBody {
	return s.Body
}

func (s *CreateConsumerChannelResponse) SetHeaders(v map[string]*string) *CreateConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerChannelResponse) SetStatusCode(v int32) *CreateConsumerChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerChannelResponse) SetBody(v *CreateConsumerChannelResponseBody) *CreateConsumerChannelResponse {
	s.Body = v
	return s
}

func (s *CreateConsumerChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
