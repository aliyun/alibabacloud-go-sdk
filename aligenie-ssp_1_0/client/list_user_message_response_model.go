// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserMessageResponse
	GetStatusCode() *int32
	SetBody(v *ListUserMessageResponseBody) *ListUserMessageResponse
	GetBody() *ListUserMessageResponseBody
}

type ListUserMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageResponse) GoString() string {
	return s.String()
}

func (s *ListUserMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserMessageResponse) GetBody() *ListUserMessageResponseBody {
	return s.Body
}

func (s *ListUserMessageResponse) SetHeaders(v map[string]*string) *ListUserMessageResponse {
	s.Headers = v
	return s
}

func (s *ListUserMessageResponse) SetStatusCode(v int32) *ListUserMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserMessageResponse) SetBody(v *ListUserMessageResponseBody) *ListUserMessageResponse {
	s.Body = v
	return s
}

func (s *ListUserMessageResponse) Validate() error {
	return dara.Validate(s)
}
