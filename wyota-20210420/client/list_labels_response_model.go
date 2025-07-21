// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLabelsResponse
	GetStatusCode() *int32
	SetBody(v *ListLabelsResponseBody) *ListLabelsResponse
	GetBody() *ListLabelsResponseBody
}

type ListLabelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLabelsResponse) GetBody() *ListLabelsResponseBody {
	return s.Body
}

func (s *ListLabelsResponse) SetHeaders(v map[string]*string) *ListLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListLabelsResponse) SetStatusCode(v int32) *ListLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLabelsResponse) SetBody(v *ListLabelsResponseBody) *ListLabelsResponse {
	s.Body = v
	return s
}

func (s *ListLabelsResponse) Validate() error {
	return dara.Validate(s)
}
