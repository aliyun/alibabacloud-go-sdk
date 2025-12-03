// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSprintsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSprintsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSprintsResponse
	GetStatusCode() *int32
	SetBody(v *ListSprintsResponseBody) *ListSprintsResponse
	GetBody() *ListSprintsResponseBody
}

type ListSprintsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSprintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSprintsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSprintsResponse) GoString() string {
	return s.String()
}

func (s *ListSprintsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSprintsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSprintsResponse) GetBody() *ListSprintsResponseBody {
	return s.Body
}

func (s *ListSprintsResponse) SetHeaders(v map[string]*string) *ListSprintsResponse {
	s.Headers = v
	return s
}

func (s *ListSprintsResponse) SetStatusCode(v int32) *ListSprintsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSprintsResponse) SetBody(v *ListSprintsResponseBody) *ListSprintsResponse {
	s.Body = v
	return s
}

func (s *ListSprintsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
