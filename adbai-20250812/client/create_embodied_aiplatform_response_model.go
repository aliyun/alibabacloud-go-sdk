// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEmbodiedAIPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEmbodiedAIPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEmbodiedAIPlatformResponse
	GetStatusCode() *int32
	SetBody(v *CreateEmbodiedAIPlatformResponseBody) *CreateEmbodiedAIPlatformResponse
	GetBody() *CreateEmbodiedAIPlatformResponseBody
}

type CreateEmbodiedAIPlatformResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEmbodiedAIPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEmbodiedAIPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformResponse) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEmbodiedAIPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEmbodiedAIPlatformResponse) GetBody() *CreateEmbodiedAIPlatformResponseBody {
	return s.Body
}

func (s *CreateEmbodiedAIPlatformResponse) SetHeaders(v map[string]*string) *CreateEmbodiedAIPlatformResponse {
	s.Headers = v
	return s
}

func (s *CreateEmbodiedAIPlatformResponse) SetStatusCode(v int32) *CreateEmbodiedAIPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEmbodiedAIPlatformResponse) SetBody(v *CreateEmbodiedAIPlatformResponseBody) *CreateEmbodiedAIPlatformResponse {
	s.Body = v
	return s
}

func (s *CreateEmbodiedAIPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
