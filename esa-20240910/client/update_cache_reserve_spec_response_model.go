// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheReserveSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCacheReserveSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCacheReserveSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCacheReserveSpecResponseBody) *UpdateCacheReserveSpecResponse
	GetBody() *UpdateCacheReserveSpecResponseBody
}

type UpdateCacheReserveSpecResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCacheReserveSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCacheReserveSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheReserveSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateCacheReserveSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCacheReserveSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCacheReserveSpecResponse) GetBody() *UpdateCacheReserveSpecResponseBody {
	return s.Body
}

func (s *UpdateCacheReserveSpecResponse) SetHeaders(v map[string]*string) *UpdateCacheReserveSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateCacheReserveSpecResponse) SetStatusCode(v int32) *UpdateCacheReserveSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCacheReserveSpecResponse) SetBody(v *UpdateCacheReserveSpecResponseBody) *UpdateCacheReserveSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateCacheReserveSpecResponse) Validate() error {
	return dara.Validate(s)
}
