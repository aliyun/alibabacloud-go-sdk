// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEditingProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEditingProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEditingProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListEditingProjectsResponseBody) *ListEditingProjectsResponse
	GetBody() *ListEditingProjectsResponseBody
}

type ListEditingProjectsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEditingProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEditingProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEditingProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListEditingProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEditingProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEditingProjectsResponse) GetBody() *ListEditingProjectsResponseBody {
	return s.Body
}

func (s *ListEditingProjectsResponse) SetHeaders(v map[string]*string) *ListEditingProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListEditingProjectsResponse) SetStatusCode(v int32) *ListEditingProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEditingProjectsResponse) SetBody(v *ListEditingProjectsResponseBody) *ListEditingProjectsResponse {
	s.Body = v
	return s
}

func (s *ListEditingProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
