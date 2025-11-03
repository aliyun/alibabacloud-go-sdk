// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotResponseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCopilotResponseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCopilotResponseResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCopilotResponseResponseBody) *GenerateCopilotResponseResponse
	GetBody() *GenerateCopilotResponseResponseBody
}

type GenerateCopilotResponseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCopilotResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCopilotResponseResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotResponseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCopilotResponseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCopilotResponseResponse) GetBody() *GenerateCopilotResponseResponseBody {
	return s.Body
}

func (s *GenerateCopilotResponseResponse) SetHeaders(v map[string]*string) *GenerateCopilotResponseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCopilotResponseResponse) SetStatusCode(v int32) *GenerateCopilotResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCopilotResponseResponse) SetBody(v *GenerateCopilotResponseResponseBody) *GenerateCopilotResponseResponse {
	s.Body = v
	return s
}

func (s *GenerateCopilotResponseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
