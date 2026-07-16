// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheReserveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCacheReserveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCacheReserveResponse
	GetStatusCode() *int32
	SetBody(v *GetCacheReserveResponseBody) *GetCacheReserveResponse
	GetBody() *GetCacheReserveResponseBody
}

type GetCacheReserveResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCacheReserveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCacheReserveResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCacheReserveResponse) GoString() string {
	return s.String()
}

func (s *GetCacheReserveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCacheReserveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCacheReserveResponse) GetBody() *GetCacheReserveResponseBody {
	return s.Body
}

func (s *GetCacheReserveResponse) SetHeaders(v map[string]*string) *GetCacheReserveResponse {
	s.Headers = v
	return s
}

func (s *GetCacheReserveResponse) SetStatusCode(v int32) *GetCacheReserveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCacheReserveResponse) SetBody(v *GetCacheReserveResponseBody) *GetCacheReserveResponse {
	s.Body = v
	return s
}

func (s *GetCacheReserveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
