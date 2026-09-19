// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateCacheResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateCacheResponseBody) *ListTemplateCacheResponse
	GetBody() *ListTemplateCacheResponseBody
}

type ListTemplateCacheResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateCacheResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateCacheResponse) GetBody() *ListTemplateCacheResponseBody {
	return s.Body
}

func (s *ListTemplateCacheResponse) SetHeaders(v map[string]*string) *ListTemplateCacheResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateCacheResponse) SetStatusCode(v int32) *ListTemplateCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateCacheResponse) SetBody(v *ListTemplateCacheResponseBody) *ListTemplateCacheResponse {
	s.Body = v
	return s
}

func (s *ListTemplateCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
