// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConsumerChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConsumerChannelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConsumerChannelResponseBody) *ModifyConsumerChannelResponse
	GetBody() *ModifyConsumerChannelResponseBody
}

type ModifyConsumerChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConsumerChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumerChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConsumerChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConsumerChannelResponse) GetBody() *ModifyConsumerChannelResponseBody {
	return s.Body
}

func (s *ModifyConsumerChannelResponse) SetHeaders(v map[string]*string) *ModifyConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumerChannelResponse) SetStatusCode(v int32) *ModifyConsumerChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConsumerChannelResponse) SetBody(v *ModifyConsumerChannelResponseBody) *ModifyConsumerChannelResponse {
	s.Body = v
	return s
}

func (s *ModifyConsumerChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
