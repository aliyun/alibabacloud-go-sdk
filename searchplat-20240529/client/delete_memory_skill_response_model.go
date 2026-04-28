// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemorySkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemorySkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemorySkillResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemorySkillResponseBody) *DeleteMemorySkillResponse
	GetBody() *DeleteMemorySkillResponseBody
}

type DeleteMemorySkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemorySkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemorySkillResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySkillResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemorySkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemorySkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemorySkillResponse) GetBody() *DeleteMemorySkillResponseBody {
	return s.Body
}

func (s *DeleteMemorySkillResponse) SetHeaders(v map[string]*string) *DeleteMemorySkillResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemorySkillResponse) SetStatusCode(v int32) *DeleteMemorySkillResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemorySkillResponse) SetBody(v *DeleteMemorySkillResponseBody) *DeleteMemorySkillResponse {
	s.Body = v
	return s
}

func (s *DeleteMemorySkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
