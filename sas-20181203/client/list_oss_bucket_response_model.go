// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOssBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOssBucketResponse
	GetStatusCode() *int32
	SetBody(v *ListOssBucketResponseBody) *ListOssBucketResponse
	GetBody() *ListOssBucketResponseBody
}

type ListOssBucketResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOssBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOssBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketResponse) GoString() string {
	return s.String()
}

func (s *ListOssBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOssBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOssBucketResponse) GetBody() *ListOssBucketResponseBody {
	return s.Body
}

func (s *ListOssBucketResponse) SetHeaders(v map[string]*string) *ListOssBucketResponse {
	s.Headers = v
	return s
}

func (s *ListOssBucketResponse) SetStatusCode(v int32) *ListOssBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOssBucketResponse) SetBody(v *ListOssBucketResponseBody) *ListOssBucketResponse {
	s.Body = v
	return s
}

func (s *ListOssBucketResponse) Validate() error {
	return dara.Validate(s)
}
