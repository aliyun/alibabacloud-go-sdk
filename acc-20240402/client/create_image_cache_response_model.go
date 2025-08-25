// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageCacheResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageCacheResponseBody) *CreateImageCacheResponse
	GetBody() *CreateImageCacheResponseBody
}

type CreateImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageCacheResponse) GetBody() *CreateImageCacheResponseBody {
	return s.Body
}

func (s *CreateImageCacheResponse) SetHeaders(v map[string]*string) *CreateImageCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateImageCacheResponse) SetStatusCode(v int32) *CreateImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageCacheResponse) SetBody(v *CreateImageCacheResponseBody) *CreateImageCacheResponse {
	s.Body = v
	return s
}

func (s *CreateImageCacheResponse) Validate() error {
	return dara.Validate(s)
}
