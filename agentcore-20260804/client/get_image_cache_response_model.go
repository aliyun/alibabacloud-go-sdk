// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageCacheResponse
	GetStatusCode() *int32
	SetBody(v *GetImageCacheResponseBody) *GetImageCacheResponse
	GetBody() *GetImageCacheResponseBody
}

type GetImageCacheResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponse) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageCacheResponse) GetBody() *GetImageCacheResponseBody {
	return s.Body
}

func (s *GetImageCacheResponse) SetHeaders(v map[string]*string) *GetImageCacheResponse {
	s.Headers = v
	return s
}

func (s *GetImageCacheResponse) SetStatusCode(v int32) *GetImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageCacheResponse) SetBody(v *GetImageCacheResponseBody) *GetImageCacheResponse {
	s.Body = v
	return s
}

func (s *GetImageCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
