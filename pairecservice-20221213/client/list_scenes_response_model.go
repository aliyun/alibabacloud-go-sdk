// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScenesResponse
	GetStatusCode() *int32
	SetBody(v *ListScenesResponseBody) *ListScenesResponse
	GetBody() *ListScenesResponseBody
}

type ListScenesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScenesResponse) GoString() string {
	return s.String()
}

func (s *ListScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScenesResponse) GetBody() *ListScenesResponseBody {
	return s.Body
}

func (s *ListScenesResponse) SetHeaders(v map[string]*string) *ListScenesResponse {
	s.Headers = v
	return s
}

func (s *ListScenesResponse) SetStatusCode(v int32) *ListScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScenesResponse) SetBody(v *ListScenesResponseBody) *ListScenesResponse {
	s.Body = v
	return s
}

func (s *ListScenesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
