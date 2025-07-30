// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTairKVCacheCustomInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTairKVCacheCustomInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTairKVCacheCustomInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartTairKVCacheCustomInstanceResponseBody) *StartTairKVCacheCustomInstanceResponse
	GetBody() *StartTairKVCacheCustomInstanceResponseBody
}

type StartTairKVCacheCustomInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTairKVCacheCustomInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTairKVCacheCustomInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTairKVCacheCustomInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartTairKVCacheCustomInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTairKVCacheCustomInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTairKVCacheCustomInstanceResponse) GetBody() *StartTairKVCacheCustomInstanceResponseBody {
	return s.Body
}

func (s *StartTairKVCacheCustomInstanceResponse) SetHeaders(v map[string]*string) *StartTairKVCacheCustomInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartTairKVCacheCustomInstanceResponse) SetStatusCode(v int32) *StartTairKVCacheCustomInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceResponse) SetBody(v *StartTairKVCacheCustomInstanceResponseBody) *StartTairKVCacheCustomInstanceResponse {
	s.Body = v
	return s
}

func (s *StartTairKVCacheCustomInstanceResponse) Validate() error {
	return dara.Validate(s)
}
