// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwaresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSoftwaresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSoftwaresResponse
	GetStatusCode() *int32
	SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse
	GetBody() *ListSoftwaresResponseBody
}

type ListSoftwaresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSoftwaresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSoftwaresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSoftwaresResponse) GetBody() *ListSoftwaresResponseBody {
	return s.Body
}

func (s *ListSoftwaresResponse) SetHeaders(v map[string]*string) *ListSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwaresResponse) SetStatusCode(v int32) *ListSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwaresResponse) SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse {
	s.Body = v
	return s
}

func (s *ListSoftwaresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
