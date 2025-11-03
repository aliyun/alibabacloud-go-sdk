// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotStreamResponseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCopilotStreamResponseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCopilotStreamResponseResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCopilotStreamResponseResponseBody) *GenerateCopilotStreamResponseResponse
	GetBody() *GenerateCopilotStreamResponseResponseBody
}

type GenerateCopilotStreamResponseResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCopilotStreamResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCopilotStreamResponseResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotStreamResponseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCopilotStreamResponseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCopilotStreamResponseResponse) GetBody() *GenerateCopilotStreamResponseResponseBody {
	return s.Body
}

func (s *GenerateCopilotStreamResponseResponse) SetHeaders(v map[string]*string) *GenerateCopilotStreamResponseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCopilotStreamResponseResponse) SetStatusCode(v int32) *GenerateCopilotStreamResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponse) SetBody(v *GenerateCopilotStreamResponseResponseBody) *GenerateCopilotStreamResponseResponse {
	s.Body = v
	return s
}

func (s *GenerateCopilotStreamResponseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
