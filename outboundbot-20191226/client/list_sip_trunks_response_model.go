// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSipTrunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSipTrunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSipTrunksResponse
	GetStatusCode() *int32
	SetBody(v *ListSipTrunksResponseBody) *ListSipTrunksResponse
	GetBody() *ListSipTrunksResponseBody
}

type ListSipTrunksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSipTrunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSipTrunksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSipTrunksResponse) GoString() string {
	return s.String()
}

func (s *ListSipTrunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSipTrunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSipTrunksResponse) GetBody() *ListSipTrunksResponseBody {
	return s.Body
}

func (s *ListSipTrunksResponse) SetHeaders(v map[string]*string) *ListSipTrunksResponse {
	s.Headers = v
	return s
}

func (s *ListSipTrunksResponse) SetStatusCode(v int32) *ListSipTrunksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSipTrunksResponse) SetBody(v *ListSipTrunksResponseBody) *ListSipTrunksResponse {
	s.Body = v
	return s
}

func (s *ListSipTrunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
