// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIspFlushCacheTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIspFlushCacheTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIspFlushCacheTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIspFlushCacheTaskResponseBody) *SubmitIspFlushCacheTaskResponse
	GetBody() *SubmitIspFlushCacheTaskResponseBody
}

type SubmitIspFlushCacheTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIspFlushCacheTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIspFlushCacheTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIspFlushCacheTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitIspFlushCacheTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIspFlushCacheTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIspFlushCacheTaskResponse) GetBody() *SubmitIspFlushCacheTaskResponseBody {
	return s.Body
}

func (s *SubmitIspFlushCacheTaskResponse) SetHeaders(v map[string]*string) *SubmitIspFlushCacheTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitIspFlushCacheTaskResponse) SetStatusCode(v int32) *SubmitIspFlushCacheTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIspFlushCacheTaskResponse) SetBody(v *SubmitIspFlushCacheTaskResponseBody) *SubmitIspFlushCacheTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitIspFlushCacheTaskResponse) Validate() error {
	return dara.Validate(s)
}
