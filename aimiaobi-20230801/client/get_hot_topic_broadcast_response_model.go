// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotTopicBroadcastResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotTopicBroadcastResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotTopicBroadcastResponse
	GetStatusCode() *int32
	SetBody(v *GetHotTopicBroadcastResponseBody) *GetHotTopicBroadcastResponse
	GetBody() *GetHotTopicBroadcastResponseBody
}

type GetHotTopicBroadcastResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotTopicBroadcastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotTopicBroadcastResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponse) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotTopicBroadcastResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotTopicBroadcastResponse) GetBody() *GetHotTopicBroadcastResponseBody {
	return s.Body
}

func (s *GetHotTopicBroadcastResponse) SetHeaders(v map[string]*string) *GetHotTopicBroadcastResponse {
	s.Headers = v
	return s
}

func (s *GetHotTopicBroadcastResponse) SetStatusCode(v int32) *GetHotTopicBroadcastResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotTopicBroadcastResponse) SetBody(v *GetHotTopicBroadcastResponseBody) *GetHotTopicBroadcastResponse {
	s.Body = v
	return s
}

func (s *GetHotTopicBroadcastResponse) Validate() error {
	return dara.Validate(s)
}
