// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatRoutingProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChatRoutingProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChatRoutingProfileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChatRoutingProfileResponseBody) *UpdateChatRoutingProfileResponse
	GetBody() *UpdateChatRoutingProfileResponseBody
}

type UpdateChatRoutingProfileResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChatRoutingProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChatRoutingProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatRoutingProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateChatRoutingProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChatRoutingProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChatRoutingProfileResponse) GetBody() *UpdateChatRoutingProfileResponseBody {
	return s.Body
}

func (s *UpdateChatRoutingProfileResponse) SetHeaders(v map[string]*string) *UpdateChatRoutingProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateChatRoutingProfileResponse) SetStatusCode(v int32) *UpdateChatRoutingProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChatRoutingProfileResponse) SetBody(v *UpdateChatRoutingProfileResponseBody) *UpdateChatRoutingProfileResponse {
	s.Body = v
	return s
}

func (s *UpdateChatRoutingProfileResponse) Validate() error {
	return dara.Validate(s)
}
