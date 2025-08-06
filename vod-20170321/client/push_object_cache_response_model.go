// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushObjectCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushObjectCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushObjectCacheResponse
	GetStatusCode() *int32
	SetBody(v *PushObjectCacheResponseBody) *PushObjectCacheResponse
	GetBody() *PushObjectCacheResponseBody
}

type PushObjectCacheResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushObjectCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushObjectCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s PushObjectCacheResponse) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushObjectCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushObjectCacheResponse) GetBody() *PushObjectCacheResponseBody {
	return s.Body
}

func (s *PushObjectCacheResponse) SetHeaders(v map[string]*string) *PushObjectCacheResponse {
	s.Headers = v
	return s
}

func (s *PushObjectCacheResponse) SetStatusCode(v int32) *PushObjectCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *PushObjectCacheResponse) SetBody(v *PushObjectCacheResponseBody) *PushObjectCacheResponse {
	s.Body = v
	return s
}

func (s *PushObjectCacheResponse) Validate() error {
	return dara.Validate(s)
}
