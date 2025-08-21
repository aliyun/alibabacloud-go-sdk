// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenStoreUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpenStoreUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpenStoreUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetOpenStoreUsageResponseBody) *GetOpenStoreUsageResponse
	GetBody() *GetOpenStoreUsageResponseBody
}

type GetOpenStoreUsageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenStoreUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenStoreUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpenStoreUsageResponse) GoString() string {
	return s.String()
}

func (s *GetOpenStoreUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpenStoreUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpenStoreUsageResponse) GetBody() *GetOpenStoreUsageResponseBody {
	return s.Body
}

func (s *GetOpenStoreUsageResponse) SetHeaders(v map[string]*string) *GetOpenStoreUsageResponse {
	s.Headers = v
	return s
}

func (s *GetOpenStoreUsageResponse) SetStatusCode(v int32) *GetOpenStoreUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenStoreUsageResponse) SetBody(v *GetOpenStoreUsageResponseBody) *GetOpenStoreUsageResponse {
	s.Body = v
	return s
}

func (s *GetOpenStoreUsageResponse) Validate() error {
	return dara.Validate(s)
}
