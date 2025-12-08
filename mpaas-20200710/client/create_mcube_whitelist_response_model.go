// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeWhitelistResponseBody) *CreateMcubeWhitelistResponse
	GetBody() *CreateMcubeWhitelistResponseBody
}

type CreateMcubeWhitelistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeWhitelistResponse) GetBody() *CreateMcubeWhitelistResponseBody {
	return s.Body
}

func (s *CreateMcubeWhitelistResponse) SetHeaders(v map[string]*string) *CreateMcubeWhitelistResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeWhitelistResponse) SetStatusCode(v int32) *CreateMcubeWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeWhitelistResponse) SetBody(v *CreateMcubeWhitelistResponseBody) *CreateMcubeWhitelistResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
