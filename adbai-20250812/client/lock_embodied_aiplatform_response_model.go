// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockEmbodiedAIPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockEmbodiedAIPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockEmbodiedAIPlatformResponse
	GetStatusCode() *int32
	SetBody(v *LockEmbodiedAIPlatformResponseBody) *LockEmbodiedAIPlatformResponse
	GetBody() *LockEmbodiedAIPlatformResponseBody
}

type LockEmbodiedAIPlatformResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockEmbodiedAIPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockEmbodiedAIPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s LockEmbodiedAIPlatformResponse) GoString() string {
	return s.String()
}

func (s *LockEmbodiedAIPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockEmbodiedAIPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockEmbodiedAIPlatformResponse) GetBody() *LockEmbodiedAIPlatformResponseBody {
	return s.Body
}

func (s *LockEmbodiedAIPlatformResponse) SetHeaders(v map[string]*string) *LockEmbodiedAIPlatformResponse {
	s.Headers = v
	return s
}

func (s *LockEmbodiedAIPlatformResponse) SetStatusCode(v int32) *LockEmbodiedAIPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *LockEmbodiedAIPlatformResponse) SetBody(v *LockEmbodiedAIPlatformResponseBody) *LockEmbodiedAIPlatformResponse {
	s.Body = v
	return s
}

func (s *LockEmbodiedAIPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
