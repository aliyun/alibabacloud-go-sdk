// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTairKVCacheCustomInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTairKVCacheCustomInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTairKVCacheCustomInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopTairKVCacheCustomInstanceResponseBody) *StopTairKVCacheCustomInstanceResponse
	GetBody() *StopTairKVCacheCustomInstanceResponseBody
}

type StopTairKVCacheCustomInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTairKVCacheCustomInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTairKVCacheCustomInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTairKVCacheCustomInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopTairKVCacheCustomInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTairKVCacheCustomInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTairKVCacheCustomInstanceResponse) GetBody() *StopTairKVCacheCustomInstanceResponseBody {
	return s.Body
}

func (s *StopTairKVCacheCustomInstanceResponse) SetHeaders(v map[string]*string) *StopTairKVCacheCustomInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopTairKVCacheCustomInstanceResponse) SetStatusCode(v int32) *StopTairKVCacheCustomInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceResponse) SetBody(v *StopTairKVCacheCustomInstanceResponseBody) *StopTairKVCacheCustomInstanceResponse {
	s.Body = v
	return s
}

func (s *StopTairKVCacheCustomInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
