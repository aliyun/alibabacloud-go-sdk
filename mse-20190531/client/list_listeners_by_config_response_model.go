// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListListenersByConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListListenersByConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListListenersByConfigResponseBody) *ListListenersByConfigResponse
	GetBody() *ListListenersByConfigResponseBody
}

type ListListenersByConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListListenersByConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListListenersByConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByConfigResponse) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListListenersByConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListListenersByConfigResponse) GetBody() *ListListenersByConfigResponseBody {
	return s.Body
}

func (s *ListListenersByConfigResponse) SetHeaders(v map[string]*string) *ListListenersByConfigResponse {
	s.Headers = v
	return s
}

func (s *ListListenersByConfigResponse) SetStatusCode(v int32) *ListListenersByConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenersByConfigResponse) SetBody(v *ListListenersByConfigResponseBody) *ListListenersByConfigResponse {
	s.Body = v
	return s
}

func (s *ListListenersByConfigResponse) Validate() error {
	return dara.Validate(s)
}
