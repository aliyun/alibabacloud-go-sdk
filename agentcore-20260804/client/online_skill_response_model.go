// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineSkillResponse
	GetStatusCode() *int32
	SetBody(v *OnlineSkillResponseBody) *OnlineSkillResponse
	GetBody() *OnlineSkillResponseBody
}

type OnlineSkillResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineSkillResponse) GoString() string {
	return s.String()
}

func (s *OnlineSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineSkillResponse) GetBody() *OnlineSkillResponseBody {
	return s.Body
}

func (s *OnlineSkillResponse) SetHeaders(v map[string]*string) *OnlineSkillResponse {
	s.Headers = v
	return s
}

func (s *OnlineSkillResponse) SetStatusCode(v int32) *OnlineSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineSkillResponse) SetBody(v *OnlineSkillResponseBody) *OnlineSkillResponse {
	s.Body = v
	return s
}

func (s *OnlineSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
