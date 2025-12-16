// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSortScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSortScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSortScriptsResponse
	GetStatusCode() *int32
	SetBody(v *ListSortScriptsResponseBody) *ListSortScriptsResponse
	GetBody() *ListSortScriptsResponseBody
}

type ListSortScriptsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSortScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSortScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSortScriptsResponse) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSortScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSortScriptsResponse) GetBody() *ListSortScriptsResponseBody {
	return s.Body
}

func (s *ListSortScriptsResponse) SetHeaders(v map[string]*string) *ListSortScriptsResponse {
	s.Headers = v
	return s
}

func (s *ListSortScriptsResponse) SetStatusCode(v int32) *ListSortScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSortScriptsResponse) SetBody(v *ListSortScriptsResponseBody) *ListSortScriptsResponse {
	s.Body = v
	return s
}

func (s *ListSortScriptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
