// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatRoutingProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatRoutingProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatRoutingProfileResponse
	GetStatusCode() *int32
	SetBody(v *GetChatRoutingProfileResponseBody) *GetChatRoutingProfileResponse
	GetBody() *GetChatRoutingProfileResponseBody
}

type GetChatRoutingProfileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatRoutingProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatRoutingProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatRoutingProfileResponse) GoString() string {
	return s.String()
}

func (s *GetChatRoutingProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatRoutingProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatRoutingProfileResponse) GetBody() *GetChatRoutingProfileResponseBody {
	return s.Body
}

func (s *GetChatRoutingProfileResponse) SetHeaders(v map[string]*string) *GetChatRoutingProfileResponse {
	s.Headers = v
	return s
}

func (s *GetChatRoutingProfileResponse) SetStatusCode(v int32) *GetChatRoutingProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatRoutingProfileResponse) SetBody(v *GetChatRoutingProfileResponseBody) *GetChatRoutingProfileResponse {
	s.Body = v
	return s
}

func (s *GetChatRoutingProfileResponse) Validate() error {
	return dara.Validate(s)
}
