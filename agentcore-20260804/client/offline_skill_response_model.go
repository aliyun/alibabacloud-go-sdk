// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineSkillResponse
	GetStatusCode() *int32
	SetBody(v *OfflineSkillResponseBody) *OfflineSkillResponse
	GetBody() *OfflineSkillResponseBody
}

type OfflineSkillResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineSkillResponse) GoString() string {
	return s.String()
}

func (s *OfflineSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineSkillResponse) GetBody() *OfflineSkillResponseBody {
	return s.Body
}

func (s *OfflineSkillResponse) SetHeaders(v map[string]*string) *OfflineSkillResponse {
	s.Headers = v
	return s
}

func (s *OfflineSkillResponse) SetStatusCode(v int32) *OfflineSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineSkillResponse) SetBody(v *OfflineSkillResponseBody) *OfflineSkillResponse {
	s.Body = v
	return s
}

func (s *OfflineSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
