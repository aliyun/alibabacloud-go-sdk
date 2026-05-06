// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeylessServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKeylessServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKeylessServersResponse
	GetStatusCode() *int32
	SetBody(v *ListKeylessServersResponseBody) *ListKeylessServersResponse
	GetBody() *ListKeylessServersResponseBody
}

type ListKeylessServersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeylessServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeylessServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKeylessServersResponse) GoString() string {
	return s.String()
}

func (s *ListKeylessServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKeylessServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKeylessServersResponse) GetBody() *ListKeylessServersResponseBody {
	return s.Body
}

func (s *ListKeylessServersResponse) SetHeaders(v map[string]*string) *ListKeylessServersResponse {
	s.Headers = v
	return s
}

func (s *ListKeylessServersResponse) SetStatusCode(v int32) *ListKeylessServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeylessServersResponse) SetBody(v *ListKeylessServersResponseBody) *ListKeylessServersResponse {
	s.Body = v
	return s
}

func (s *ListKeylessServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
