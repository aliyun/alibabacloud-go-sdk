// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockEmbodiedAIPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnlockEmbodiedAIPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnlockEmbodiedAIPlatformResponse
	GetStatusCode() *int32
	SetBody(v *UnlockEmbodiedAIPlatformResponseBody) *UnlockEmbodiedAIPlatformResponse
	GetBody() *UnlockEmbodiedAIPlatformResponseBody
}

type UnlockEmbodiedAIPlatformResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockEmbodiedAIPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockEmbodiedAIPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s UnlockEmbodiedAIPlatformResponse) GoString() string {
	return s.String()
}

func (s *UnlockEmbodiedAIPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnlockEmbodiedAIPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnlockEmbodiedAIPlatformResponse) GetBody() *UnlockEmbodiedAIPlatformResponseBody {
	return s.Body
}

func (s *UnlockEmbodiedAIPlatformResponse) SetHeaders(v map[string]*string) *UnlockEmbodiedAIPlatformResponse {
	s.Headers = v
	return s
}

func (s *UnlockEmbodiedAIPlatformResponse) SetStatusCode(v int32) *UnlockEmbodiedAIPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformResponse) SetBody(v *UnlockEmbodiedAIPlatformResponseBody) *UnlockEmbodiedAIPlatformResponse {
	s.Body = v
	return s
}

func (s *UnlockEmbodiedAIPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
