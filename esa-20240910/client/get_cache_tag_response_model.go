// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCacheTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCacheTagResponse
	GetStatusCode() *int32
	SetBody(v *GetCacheTagResponseBody) *GetCacheTagResponse
	GetBody() *GetCacheTagResponseBody
}

type GetCacheTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCacheTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCacheTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCacheTagResponse) GoString() string {
	return s.String()
}

func (s *GetCacheTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCacheTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCacheTagResponse) GetBody() *GetCacheTagResponseBody {
	return s.Body
}

func (s *GetCacheTagResponse) SetHeaders(v map[string]*string) *GetCacheTagResponse {
	s.Headers = v
	return s
}

func (s *GetCacheTagResponse) SetStatusCode(v int32) *GetCacheTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCacheTagResponse) SetBody(v *GetCacheTagResponseBody) *GetCacheTagResponse {
	s.Body = v
	return s
}

func (s *GetCacheTagResponse) Validate() error {
	return dara.Validate(s)
}
