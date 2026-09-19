// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplateCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplateCacheResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplateCacheResponseBody) *DeleteTemplateCacheResponse
	GetBody() *DeleteTemplateCacheResponseBody
}

type DeleteTemplateCacheResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplateCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplateCacheResponse) GetBody() *DeleteTemplateCacheResponseBody {
	return s.Body
}

func (s *DeleteTemplateCacheResponse) SetHeaders(v map[string]*string) *DeleteTemplateCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateCacheResponse) SetStatusCode(v int32) *DeleteTemplateCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateCacheResponse) SetBody(v *DeleteTemplateCacheResponseBody) *DeleteTemplateCacheResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplateCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
