// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventsAttentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventsAttentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventsAttentionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventsAttentionResponseBody) *UpdateEventsAttentionResponse
	GetBody() *UpdateEventsAttentionResponseBody
}

type UpdateEventsAttentionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventsAttentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventsAttentionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventsAttentionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventsAttentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventsAttentionResponse) GetBody() *UpdateEventsAttentionResponseBody {
	return s.Body
}

func (s *UpdateEventsAttentionResponse) SetHeaders(v map[string]*string) *UpdateEventsAttentionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventsAttentionResponse) SetStatusCode(v int32) *UpdateEventsAttentionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventsAttentionResponse) SetBody(v *UpdateEventsAttentionResponseBody) *UpdateEventsAttentionResponse {
	s.Body = v
	return s
}

func (s *UpdateEventsAttentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
