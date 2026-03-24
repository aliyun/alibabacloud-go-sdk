// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmbodiedAIPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEmbodiedAIPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEmbodiedAIPlatformResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEmbodiedAIPlatformResponseBody) *ModifyEmbodiedAIPlatformResponse
	GetBody() *ModifyEmbodiedAIPlatformResponseBody
}

type ModifyEmbodiedAIPlatformResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEmbodiedAIPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEmbodiedAIPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformResponse) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEmbodiedAIPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEmbodiedAIPlatformResponse) GetBody() *ModifyEmbodiedAIPlatformResponseBody {
	return s.Body
}

func (s *ModifyEmbodiedAIPlatformResponse) SetHeaders(v map[string]*string) *ModifyEmbodiedAIPlatformResponse {
	s.Headers = v
	return s
}

func (s *ModifyEmbodiedAIPlatformResponse) SetStatusCode(v int32) *ModifyEmbodiedAIPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformResponse) SetBody(v *ModifyEmbodiedAIPlatformResponseBody) *ModifyEmbodiedAIPlatformResponse {
	s.Body = v
	return s
}

func (s *ModifyEmbodiedAIPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
