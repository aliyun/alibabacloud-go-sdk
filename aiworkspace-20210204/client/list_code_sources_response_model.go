// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCodeSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCodeSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListCodeSourcesResponseBody) *ListCodeSourcesResponse
	GetBody() *ListCodeSourcesResponseBody
}

type ListCodeSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCodeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCodeSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCodeSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCodeSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCodeSourcesResponse) GetBody() *ListCodeSourcesResponseBody {
	return s.Body
}

func (s *ListCodeSourcesResponse) SetHeaders(v map[string]*string) *ListCodeSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListCodeSourcesResponse) SetStatusCode(v int32) *ListCodeSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCodeSourcesResponse) SetBody(v *ListCodeSourcesResponseBody) *ListCodeSourcesResponse {
	s.Body = v
	return s
}

func (s *ListCodeSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
