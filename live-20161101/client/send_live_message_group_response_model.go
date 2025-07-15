// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *SendLiveMessageGroupResponseBody) *SendLiveMessageGroupResponse
	GetBody() *SendLiveMessageGroupResponseBody
}

type SendLiveMessageGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SendLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *SendLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendLiveMessageGroupResponse) GetBody() *SendLiveMessageGroupResponseBody {
	return s.Body
}

func (s *SendLiveMessageGroupResponse) SetHeaders(v map[string]*string) *SendLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *SendLiveMessageGroupResponse) SetStatusCode(v int32) *SendLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLiveMessageGroupResponse) SetBody(v *SendLiveMessageGroupResponseBody) *SendLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *SendLiveMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
