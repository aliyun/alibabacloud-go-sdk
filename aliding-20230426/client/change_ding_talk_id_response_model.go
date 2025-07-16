// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDingTalkIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDingTalkIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDingTalkIdResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDingTalkIdResponseBody) *ChangeDingTalkIdResponse
	GetBody() *ChangeDingTalkIdResponseBody
}

type ChangeDingTalkIdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDingTalkIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDingTalkIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdResponse) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDingTalkIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDingTalkIdResponse) GetBody() *ChangeDingTalkIdResponseBody {
	return s.Body
}

func (s *ChangeDingTalkIdResponse) SetHeaders(v map[string]*string) *ChangeDingTalkIdResponse {
	s.Headers = v
	return s
}

func (s *ChangeDingTalkIdResponse) SetStatusCode(v int32) *ChangeDingTalkIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDingTalkIdResponse) SetBody(v *ChangeDingTalkIdResponseBody) *ChangeDingTalkIdResponse {
	s.Body = v
	return s
}

func (s *ChangeDingTalkIdResponse) Validate() error {
	return dara.Validate(s)
}
