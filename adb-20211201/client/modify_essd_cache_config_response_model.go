// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEssdCacheConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEssdCacheConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEssdCacheConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEssdCacheConfigResponseBody) *ModifyEssdCacheConfigResponse
	GetBody() *ModifyEssdCacheConfigResponseBody
}

type ModifyEssdCacheConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEssdCacheConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEssdCacheConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEssdCacheConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyEssdCacheConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEssdCacheConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEssdCacheConfigResponse) GetBody() *ModifyEssdCacheConfigResponseBody {
	return s.Body
}

func (s *ModifyEssdCacheConfigResponse) SetHeaders(v map[string]*string) *ModifyEssdCacheConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyEssdCacheConfigResponse) SetStatusCode(v int32) *ModifyEssdCacheConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEssdCacheConfigResponse) SetBody(v *ModifyEssdCacheConfigResponseBody) *ModifyEssdCacheConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyEssdCacheConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
