// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickLiveMessageGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KickLiveMessageGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KickLiveMessageGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *KickLiveMessageGroupUserResponseBody) *KickLiveMessageGroupUserResponse
	GetBody() *KickLiveMessageGroupUserResponseBody
}

type KickLiveMessageGroupUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KickLiveMessageGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KickLiveMessageGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s KickLiveMessageGroupUserResponse) GoString() string {
	return s.String()
}

func (s *KickLiveMessageGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KickLiveMessageGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KickLiveMessageGroupUserResponse) GetBody() *KickLiveMessageGroupUserResponseBody {
	return s.Body
}

func (s *KickLiveMessageGroupUserResponse) SetHeaders(v map[string]*string) *KickLiveMessageGroupUserResponse {
	s.Headers = v
	return s
}

func (s *KickLiveMessageGroupUserResponse) SetStatusCode(v int32) *KickLiveMessageGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *KickLiveMessageGroupUserResponse) SetBody(v *KickLiveMessageGroupUserResponseBody) *KickLiveMessageGroupUserResponse {
	s.Body = v
	return s
}

func (s *KickLiveMessageGroupUserResponse) Validate() error {
	return dara.Validate(s)
}
