// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListListenersByIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListListenersByIpResponse
	GetStatusCode() *int32
	SetBody(v *ListListenersByIpResponseBody) *ListListenersByIpResponse
	GetBody() *ListListenersByIpResponseBody
}

type ListListenersByIpResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListListenersByIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListListenersByIpResponse) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByIpResponse) GoString() string {
	return s.String()
}

func (s *ListListenersByIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListListenersByIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListListenersByIpResponse) GetBody() *ListListenersByIpResponseBody {
	return s.Body
}

func (s *ListListenersByIpResponse) SetHeaders(v map[string]*string) *ListListenersByIpResponse {
	s.Headers = v
	return s
}

func (s *ListListenersByIpResponse) SetStatusCode(v int32) *ListListenersByIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenersByIpResponse) SetBody(v *ListListenersByIpResponseBody) *ListListenersByIpResponse {
	s.Body = v
	return s
}

func (s *ListListenersByIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
