// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendHotlineHeartBeatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendHotlineHeartBeatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendHotlineHeartBeatResponse
	GetStatusCode() *int32
	SetBody(v *SendHotlineHeartBeatResponseBody) *SendHotlineHeartBeatResponse
	GetBody() *SendHotlineHeartBeatResponseBody
}

type SendHotlineHeartBeatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendHotlineHeartBeatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendHotlineHeartBeatResponse) String() string {
	return dara.Prettify(s)
}

func (s SendHotlineHeartBeatResponse) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendHotlineHeartBeatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendHotlineHeartBeatResponse) GetBody() *SendHotlineHeartBeatResponseBody {
	return s.Body
}

func (s *SendHotlineHeartBeatResponse) SetHeaders(v map[string]*string) *SendHotlineHeartBeatResponse {
	s.Headers = v
	return s
}

func (s *SendHotlineHeartBeatResponse) SetStatusCode(v int32) *SendHotlineHeartBeatResponse {
	s.StatusCode = &v
	return s
}

func (s *SendHotlineHeartBeatResponse) SetBody(v *SendHotlineHeartBeatResponseBody) *SendHotlineHeartBeatResponse {
	s.Body = v
	return s
}

func (s *SendHotlineHeartBeatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
