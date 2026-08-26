// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentThemeResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentThemeResponseBody) *ListDataAgentThemeResponse
	GetBody() *ListDataAgentThemeResponseBody
}

type ListDataAgentThemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentThemeResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentThemeResponse) GetBody() *ListDataAgentThemeResponseBody {
	return s.Body
}

func (s *ListDataAgentThemeResponse) SetHeaders(v map[string]*string) *ListDataAgentThemeResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentThemeResponse) SetStatusCode(v int32) *ListDataAgentThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentThemeResponse) SetBody(v *ListDataAgentThemeResponseBody) *ListDataAgentThemeResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentThemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
