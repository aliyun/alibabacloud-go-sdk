// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWritingStylesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWritingStylesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWritingStylesResponse
	GetStatusCode() *int32
	SetBody(v *ListWritingStylesResponseBody) *ListWritingStylesResponse
	GetBody() *ListWritingStylesResponseBody
}

type ListWritingStylesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWritingStylesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWritingStylesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWritingStylesResponse) GoString() string {
	return s.String()
}

func (s *ListWritingStylesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWritingStylesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWritingStylesResponse) GetBody() *ListWritingStylesResponseBody {
	return s.Body
}

func (s *ListWritingStylesResponse) SetHeaders(v map[string]*string) *ListWritingStylesResponse {
	s.Headers = v
	return s
}

func (s *ListWritingStylesResponse) SetStatusCode(v int32) *ListWritingStylesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWritingStylesResponse) SetBody(v *ListWritingStylesResponseBody) *ListWritingStylesResponse {
	s.Body = v
	return s
}

func (s *ListWritingStylesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
