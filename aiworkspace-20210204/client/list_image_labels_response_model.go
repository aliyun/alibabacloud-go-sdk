// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageLabelsResponse
	GetStatusCode() *int32
	SetBody(v *ListImageLabelsResponseBody) *ListImageLabelsResponse
	GetBody() *ListImageLabelsResponseBody
}

type ListImageLabelsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageLabelsResponse) GetBody() *ListImageLabelsResponseBody {
	return s.Body
}

func (s *ListImageLabelsResponse) SetHeaders(v map[string]*string) *ListImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListImageLabelsResponse) SetStatusCode(v int32) *ListImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageLabelsResponse) SetBody(v *ListImageLabelsResponseBody) *ListImageLabelsResponse {
	s.Body = v
	return s
}

func (s *ListImageLabelsResponse) Validate() error {
	return dara.Validate(s)
}
