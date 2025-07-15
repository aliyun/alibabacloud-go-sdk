// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeneratedContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGeneratedContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGeneratedContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListGeneratedContentsResponseBody) *ListGeneratedContentsResponse
	GetBody() *ListGeneratedContentsResponseBody
}

type ListGeneratedContentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGeneratedContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGeneratedContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGeneratedContentsResponse) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGeneratedContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGeneratedContentsResponse) GetBody() *ListGeneratedContentsResponseBody {
	return s.Body
}

func (s *ListGeneratedContentsResponse) SetHeaders(v map[string]*string) *ListGeneratedContentsResponse {
	s.Headers = v
	return s
}

func (s *ListGeneratedContentsResponse) SetStatusCode(v int32) *ListGeneratedContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGeneratedContentsResponse) SetBody(v *ListGeneratedContentsResponseBody) *ListGeneratedContentsResponse {
	s.Body = v
	return s
}

func (s *ListGeneratedContentsResponse) Validate() error {
	return dara.Validate(s)
}
