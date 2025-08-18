// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheReserveSpecificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCacheReserveSpecificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCacheReserveSpecificationResponse
	GetStatusCode() *int32
	SetBody(v *GetCacheReserveSpecificationResponseBody) *GetCacheReserveSpecificationResponse
	GetBody() *GetCacheReserveSpecificationResponseBody
}

type GetCacheReserveSpecificationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCacheReserveSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCacheReserveSpecificationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCacheReserveSpecificationResponse) GoString() string {
	return s.String()
}

func (s *GetCacheReserveSpecificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCacheReserveSpecificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCacheReserveSpecificationResponse) GetBody() *GetCacheReserveSpecificationResponseBody {
	return s.Body
}

func (s *GetCacheReserveSpecificationResponse) SetHeaders(v map[string]*string) *GetCacheReserveSpecificationResponse {
	s.Headers = v
	return s
}

func (s *GetCacheReserveSpecificationResponse) SetStatusCode(v int32) *GetCacheReserveSpecificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCacheReserveSpecificationResponse) SetBody(v *GetCacheReserveSpecificationResponseBody) *GetCacheReserveSpecificationResponse {
	s.Body = v
	return s
}

func (s *GetCacheReserveSpecificationResponse) Validate() error {
	return dara.Validate(s)
}
