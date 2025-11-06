// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeWhitelistForIdeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeWhitelistForIdeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeWhitelistForIdeResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeWhitelistForIdeResponseBody) *CreateMcubeWhitelistForIdeResponse
	GetBody() *CreateMcubeWhitelistForIdeResponseBody
}

type CreateMcubeWhitelistForIdeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeWhitelistForIdeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeWhitelistForIdeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistForIdeResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistForIdeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeWhitelistForIdeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeWhitelistForIdeResponse) GetBody() *CreateMcubeWhitelistForIdeResponseBody {
	return s.Body
}

func (s *CreateMcubeWhitelistForIdeResponse) SetHeaders(v map[string]*string) *CreateMcubeWhitelistForIdeResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponse) SetStatusCode(v int32) *CreateMcubeWhitelistForIdeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponse) SetBody(v *CreateMcubeWhitelistForIdeResponseBody) *CreateMcubeWhitelistForIdeResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeWhitelistForIdeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
