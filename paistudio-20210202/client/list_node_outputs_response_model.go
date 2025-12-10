// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeOutputsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeOutputsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeOutputsResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeOutputsResponseBody) *ListNodeOutputsResponse
	GetBody() *ListNodeOutputsResponseBody
}

type ListNodeOutputsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeOutputsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeOutputsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeOutputsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeOutputsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeOutputsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeOutputsResponse) GetBody() *ListNodeOutputsResponseBody {
	return s.Body
}

func (s *ListNodeOutputsResponse) SetHeaders(v map[string]*string) *ListNodeOutputsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeOutputsResponse) SetStatusCode(v int32) *ListNodeOutputsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeOutputsResponse) SetBody(v *ListNodeOutputsResponseBody) *ListNodeOutputsResponse {
	s.Body = v
	return s
}

func (s *ListNodeOutputsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
