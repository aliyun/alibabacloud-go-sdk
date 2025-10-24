// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagResponse
	GetStatusCode() *int32
	SetBody(v *ListTagResponseBody) *ListTagResponse
	GetBody() *ListTagResponseBody
}

type ListTagResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponse) GoString() string {
	return s.String()
}

func (s *ListTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagResponse) GetBody() *ListTagResponseBody {
	return s.Body
}

func (s *ListTagResponse) SetHeaders(v map[string]*string) *ListTagResponse {
	s.Headers = v
	return s
}

func (s *ListTagResponse) SetStatusCode(v int32) *ListTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResponse) SetBody(v *ListTagResponseBody) *ListTagResponse {
	s.Body = v
	return s
}

func (s *ListTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
