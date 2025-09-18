// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishedMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishedMmAppResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishedMmAppResponseBody) *ListPublishedMmAppResponse
	GetBody() *ListPublishedMmAppResponseBody
}

type ListPublishedMmAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedMmAppResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishedMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishedMmAppResponse) GetBody() *ListPublishedMmAppResponseBody {
	return s.Body
}

func (s *ListPublishedMmAppResponse) SetHeaders(v map[string]*string) *ListPublishedMmAppResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedMmAppResponse) SetStatusCode(v int32) *ListPublishedMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedMmAppResponse) SetBody(v *ListPublishedMmAppResponseBody) *ListPublishedMmAppResponse {
	s.Body = v
	return s
}

func (s *ListPublishedMmAppResponse) Validate() error {
	return dara.Validate(s)
}
