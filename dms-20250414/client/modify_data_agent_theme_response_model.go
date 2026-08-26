// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataAgentThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataAgentThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataAgentThemeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataAgentThemeResponseBody) *ModifyDataAgentThemeResponse
	GetBody() *ModifyDataAgentThemeResponseBody
}

type ModifyDataAgentThemeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataAgentThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataAgentThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataAgentThemeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataAgentThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataAgentThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataAgentThemeResponse) GetBody() *ModifyDataAgentThemeResponseBody {
	return s.Body
}

func (s *ModifyDataAgentThemeResponse) SetHeaders(v map[string]*string) *ModifyDataAgentThemeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataAgentThemeResponse) SetStatusCode(v int32) *ModifyDataAgentThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataAgentThemeResponse) SetBody(v *ModifyDataAgentThemeResponseBody) *ModifyDataAgentThemeResponse {
	s.Body = v
	return s
}

func (s *ModifyDataAgentThemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
