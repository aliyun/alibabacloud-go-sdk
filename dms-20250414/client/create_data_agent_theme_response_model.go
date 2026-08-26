// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentThemeResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentThemeResponseBody) *CreateDataAgentThemeResponse
	GetBody() *CreateDataAgentThemeResponseBody
}

type CreateDataAgentThemeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentThemeResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentThemeResponse) GetBody() *CreateDataAgentThemeResponseBody {
	return s.Body
}

func (s *CreateDataAgentThemeResponse) SetHeaders(v map[string]*string) *CreateDataAgentThemeResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentThemeResponse) SetStatusCode(v int32) *CreateDataAgentThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentThemeResponse) SetBody(v *CreateDataAgentThemeResponseBody) *CreateDataAgentThemeResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentThemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
