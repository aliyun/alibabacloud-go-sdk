// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCacheServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCacheServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCacheServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCacheServicesResponseBody) *ListDataCacheServicesResponse
	GetBody() *ListDataCacheServicesResponseBody
}

type ListDataCacheServicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCacheServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCacheServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCacheServicesResponse) GoString() string {
	return s.String()
}

func (s *ListDataCacheServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCacheServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCacheServicesResponse) GetBody() *ListDataCacheServicesResponseBody {
	return s.Body
}

func (s *ListDataCacheServicesResponse) SetHeaders(v map[string]*string) *ListDataCacheServicesResponse {
	s.Headers = v
	return s
}

func (s *ListDataCacheServicesResponse) SetStatusCode(v int32) *ListDataCacheServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCacheServicesResponse) SetBody(v *ListDataCacheServicesResponseBody) *ListDataCacheServicesResponse {
	s.Body = v
	return s
}

func (s *ListDataCacheServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
