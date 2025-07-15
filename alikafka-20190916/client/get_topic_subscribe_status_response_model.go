// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicSubscribeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTopicSubscribeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTopicSubscribeStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTopicSubscribeStatusResponseBody) *GetTopicSubscribeStatusResponse
	GetBody() *GetTopicSubscribeStatusResponseBody
}

type GetTopicSubscribeStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicSubscribeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicSubscribeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSubscribeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTopicSubscribeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTopicSubscribeStatusResponse) GetBody() *GetTopicSubscribeStatusResponseBody {
	return s.Body
}

func (s *GetTopicSubscribeStatusResponse) SetHeaders(v map[string]*string) *GetTopicSubscribeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTopicSubscribeStatusResponse) SetStatusCode(v int32) *GetTopicSubscribeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicSubscribeStatusResponse) SetBody(v *GetTopicSubscribeStatusResponseBody) *GetTopicSubscribeStatusResponse {
	s.Body = v
	return s
}

func (s *GetTopicSubscribeStatusResponse) Validate() error {
	return dara.Validate(s)
}
