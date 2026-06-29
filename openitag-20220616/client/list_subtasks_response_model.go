// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubtasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubtasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubtasksResponse
	GetStatusCode() *int32
	SetBody(v *ListSubtasksResponseBody) *ListSubtasksResponse
	GetBody() *ListSubtasksResponseBody
}

type ListSubtasksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubtasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubtasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubtasksResponse) GoString() string {
	return s.String()
}

func (s *ListSubtasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubtasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubtasksResponse) GetBody() *ListSubtasksResponseBody {
	return s.Body
}

func (s *ListSubtasksResponse) SetHeaders(v map[string]*string) *ListSubtasksResponse {
	s.Headers = v
	return s
}

func (s *ListSubtasksResponse) SetStatusCode(v int32) *ListSubtasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubtasksResponse) SetBody(v *ListSubtasksResponseBody) *ListSubtasksResponse {
	s.Body = v
	return s
}

func (s *ListSubtasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
