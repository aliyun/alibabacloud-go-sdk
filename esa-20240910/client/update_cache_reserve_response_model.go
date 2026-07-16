// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheReserveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCacheReserveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCacheReserveResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCacheReserveResponseBody) *UpdateCacheReserveResponse
	GetBody() *UpdateCacheReserveResponseBody
}

type UpdateCacheReserveResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCacheReserveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCacheReserveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheReserveResponse) GoString() string {
	return s.String()
}

func (s *UpdateCacheReserveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCacheReserveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCacheReserveResponse) GetBody() *UpdateCacheReserveResponseBody {
	return s.Body
}

func (s *UpdateCacheReserveResponse) SetHeaders(v map[string]*string) *UpdateCacheReserveResponse {
	s.Headers = v
	return s
}

func (s *UpdateCacheReserveResponse) SetStatusCode(v int32) *UpdateCacheReserveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCacheReserveResponse) SetBody(v *UpdateCacheReserveResponseBody) *UpdateCacheReserveResponse {
	s.Body = v
	return s
}

func (s *UpdateCacheReserveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
