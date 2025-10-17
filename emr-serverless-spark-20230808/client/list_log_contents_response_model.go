// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogContentsResponseBody) *ListLogContentsResponse
	GetBody() *ListLogContentsResponseBody
}

type ListLogContentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogContentsResponse) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogContentsResponse) GetBody() *ListLogContentsResponseBody {
	return s.Body
}

func (s *ListLogContentsResponse) SetHeaders(v map[string]*string) *ListLogContentsResponse {
	s.Headers = v
	return s
}

func (s *ListLogContentsResponse) SetStatusCode(v int32) *ListLogContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogContentsResponse) SetBody(v *ListLogContentsResponseBody) *ListLogContentsResponse {
	s.Body = v
	return s
}

func (s *ListLogContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
