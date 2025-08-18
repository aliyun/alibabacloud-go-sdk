// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTieredCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTieredCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTieredCacheResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTieredCacheResponseBody) *UpdateTieredCacheResponse
	GetBody() *UpdateTieredCacheResponseBody
}

type UpdateTieredCacheResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTieredCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTieredCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTieredCacheResponse) GoString() string {
	return s.String()
}

func (s *UpdateTieredCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTieredCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTieredCacheResponse) GetBody() *UpdateTieredCacheResponseBody {
	return s.Body
}

func (s *UpdateTieredCacheResponse) SetHeaders(v map[string]*string) *UpdateTieredCacheResponse {
	s.Headers = v
	return s
}

func (s *UpdateTieredCacheResponse) SetStatusCode(v int32) *UpdateTieredCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTieredCacheResponse) SetBody(v *UpdateTieredCacheResponseBody) *UpdateTieredCacheResponse {
	s.Body = v
	return s
}

func (s *UpdateTieredCacheResponse) Validate() error {
	return dara.Validate(s)
}
