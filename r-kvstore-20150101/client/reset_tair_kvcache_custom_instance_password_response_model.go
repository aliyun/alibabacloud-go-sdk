// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTairKVCacheCustomInstancePasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetTairKVCacheCustomInstancePasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetTairKVCacheCustomInstancePasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetTairKVCacheCustomInstancePasswordResponseBody) *ResetTairKVCacheCustomInstancePasswordResponse
	GetBody() *ResetTairKVCacheCustomInstancePasswordResponseBody
}

type ResetTairKVCacheCustomInstancePasswordResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetTairKVCacheCustomInstancePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetTairKVCacheCustomInstancePasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetTairKVCacheCustomInstancePasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) GetBody() *ResetTairKVCacheCustomInstancePasswordResponseBody {
	return s.Body
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) SetHeaders(v map[string]*string) *ResetTairKVCacheCustomInstancePasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) SetStatusCode(v int32) *ResetTairKVCacheCustomInstancePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) SetBody(v *ResetTairKVCacheCustomInstancePasswordResponseBody) *ResetTairKVCacheCustomInstancePasswordResponse {
	s.Body = v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordResponse) Validate() error {
	return dara.Validate(s)
}
