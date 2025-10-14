// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIspFlushCacheInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIspFlushCacheInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIspFlushCacheInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIspFlushCacheInstanceConfigResponseBody) *UpdateIspFlushCacheInstanceConfigResponse
	GetBody() *UpdateIspFlushCacheInstanceConfigResponseBody
}

type UpdateIspFlushCacheInstanceConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIspFlushCacheInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIspFlushCacheInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIspFlushCacheInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) GetBody() *UpdateIspFlushCacheInstanceConfigResponseBody {
	return s.Body
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) SetHeaders(v map[string]*string) *UpdateIspFlushCacheInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) SetStatusCode(v int32) *UpdateIspFlushCacheInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) SetBody(v *UpdateIspFlushCacheInstanceConfigResponseBody) *UpdateIspFlushCacheInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateIspFlushCacheInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
