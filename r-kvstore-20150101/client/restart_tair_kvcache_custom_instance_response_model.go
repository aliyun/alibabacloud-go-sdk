// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartTairKVCacheCustomInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartTairKVCacheCustomInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartTairKVCacheCustomInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartTairKVCacheCustomInstanceResponseBody) *RestartTairKVCacheCustomInstanceResponse
	GetBody() *RestartTairKVCacheCustomInstanceResponseBody
}

type RestartTairKVCacheCustomInstanceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartTairKVCacheCustomInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartTairKVCacheCustomInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartTairKVCacheCustomInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartTairKVCacheCustomInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartTairKVCacheCustomInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartTairKVCacheCustomInstanceResponse) GetBody() *RestartTairKVCacheCustomInstanceResponseBody {
	return s.Body
}

func (s *RestartTairKVCacheCustomInstanceResponse) SetHeaders(v map[string]*string) *RestartTairKVCacheCustomInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartTairKVCacheCustomInstanceResponse) SetStatusCode(v int32) *RestartTairKVCacheCustomInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceResponse) SetBody(v *RestartTairKVCacheCustomInstanceResponseBody) *RestartTairKVCacheCustomInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartTairKVCacheCustomInstanceResponse) Validate() error {
	return dara.Validate(s)
}
