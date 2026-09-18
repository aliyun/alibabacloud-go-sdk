// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateRelayPollerScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateRelayPollerScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateRelayPollerScriptResponse
	GetStatusCode() *int32
	SetBody(v *GenerateRelayPollerScriptResponseBody) *GenerateRelayPollerScriptResponse
	GetBody() *GenerateRelayPollerScriptResponseBody
}

type GenerateRelayPollerScriptResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateRelayPollerScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateRelayPollerScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateRelayPollerScriptResponse) GoString() string {
	return s.String()
}

func (s *GenerateRelayPollerScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateRelayPollerScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateRelayPollerScriptResponse) GetBody() *GenerateRelayPollerScriptResponseBody {
	return s.Body
}

func (s *GenerateRelayPollerScriptResponse) SetHeaders(v map[string]*string) *GenerateRelayPollerScriptResponse {
	s.Headers = v
	return s
}

func (s *GenerateRelayPollerScriptResponse) SetStatusCode(v int32) *GenerateRelayPollerScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateRelayPollerScriptResponse) SetBody(v *GenerateRelayPollerScriptResponseBody) *GenerateRelayPollerScriptResponse {
	s.Body = v
	return s
}

func (s *GenerateRelayPollerScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
