// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTensorboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTensorboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTensorboardsResponse
	GetStatusCode() *int32
	SetBody(v *ListTensorboardsResponseBody) *ListTensorboardsResponse
	GetBody() *ListTensorboardsResponseBody
}

type ListTensorboardsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTensorboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTensorboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTensorboardsResponse) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTensorboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTensorboardsResponse) GetBody() *ListTensorboardsResponseBody {
	return s.Body
}

func (s *ListTensorboardsResponse) SetHeaders(v map[string]*string) *ListTensorboardsResponse {
	s.Headers = v
	return s
}

func (s *ListTensorboardsResponse) SetStatusCode(v int32) *ListTensorboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTensorboardsResponse) SetBody(v *ListTensorboardsResponseBody) *ListTensorboardsResponse {
	s.Body = v
	return s
}

func (s *ListTensorboardsResponse) Validate() error {
	return dara.Validate(s)
}
