// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotQueryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KopilotQueryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KopilotQueryStatusResponse
	GetStatusCode() *int32
	SetBody(v *KopilotQueryStatusResponseBody) *KopilotQueryStatusResponse
	GetBody() *KopilotQueryStatusResponseBody
}

type KopilotQueryStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KopilotQueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KopilotQueryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s KopilotQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *KopilotQueryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KopilotQueryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KopilotQueryStatusResponse) GetBody() *KopilotQueryStatusResponseBody {
	return s.Body
}

func (s *KopilotQueryStatusResponse) SetHeaders(v map[string]*string) *KopilotQueryStatusResponse {
	s.Headers = v
	return s
}

func (s *KopilotQueryStatusResponse) SetStatusCode(v int32) *KopilotQueryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *KopilotQueryStatusResponse) SetBody(v *KopilotQueryStatusResponseBody) *KopilotQueryStatusResponse {
	s.Body = v
	return s
}

func (s *KopilotQueryStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
