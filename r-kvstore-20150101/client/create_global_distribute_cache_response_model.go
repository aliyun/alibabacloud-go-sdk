// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDistributeCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalDistributeCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalDistributeCacheResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalDistributeCacheResponseBody) *CreateGlobalDistributeCacheResponse
	GetBody() *CreateGlobalDistributeCacheResponseBody
}

type CreateGlobalDistributeCacheResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalDistributeCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalDistributeCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDistributeCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalDistributeCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalDistributeCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalDistributeCacheResponse) GetBody() *CreateGlobalDistributeCacheResponseBody {
	return s.Body
}

func (s *CreateGlobalDistributeCacheResponse) SetHeaders(v map[string]*string) *CreateGlobalDistributeCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalDistributeCacheResponse) SetStatusCode(v int32) *CreateGlobalDistributeCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalDistributeCacheResponse) SetBody(v *CreateGlobalDistributeCacheResponseBody) *CreateGlobalDistributeCacheResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalDistributeCacheResponse) Validate() error {
	return dara.Validate(s)
}
