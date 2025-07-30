// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeTairKVCacheCustomInstanceDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeTairKVCacheCustomInstanceDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeTairKVCacheCustomInstanceDiskResponse
	GetStatusCode() *int32
	SetBody(v *ResizeTairKVCacheCustomInstanceDiskResponseBody) *ResizeTairKVCacheCustomInstanceDiskResponse
	GetBody() *ResizeTairKVCacheCustomInstanceDiskResponseBody
}

type ResizeTairKVCacheCustomInstanceDiskResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeTairKVCacheCustomInstanceDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeTairKVCacheCustomInstanceDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeTairKVCacheCustomInstanceDiskResponse) GoString() string {
	return s.String()
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) GetBody() *ResizeTairKVCacheCustomInstanceDiskResponseBody {
	return s.Body
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) SetHeaders(v map[string]*string) *ResizeTairKVCacheCustomInstanceDiskResponse {
	s.Headers = v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) SetStatusCode(v int32) *ResizeTairKVCacheCustomInstanceDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) SetBody(v *ResizeTairKVCacheCustomInstanceDiskResponseBody) *ResizeTairKVCacheCustomInstanceDiskResponse {
	s.Body = v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponse) Validate() error {
	return dara.Validate(s)
}
