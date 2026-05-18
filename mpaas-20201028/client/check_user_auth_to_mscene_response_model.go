// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserAuthToMsceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUserAuthToMsceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUserAuthToMsceneResponse
	GetStatusCode() *int32
	SetBody(v *CheckUserAuthToMsceneResponseBody) *CheckUserAuthToMsceneResponse
	GetBody() *CheckUserAuthToMsceneResponseBody
}

type CheckUserAuthToMsceneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserAuthToMsceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserAuthToMsceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserAuthToMsceneResponse) GoString() string {
	return s.String()
}

func (s *CheckUserAuthToMsceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUserAuthToMsceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUserAuthToMsceneResponse) GetBody() *CheckUserAuthToMsceneResponseBody {
	return s.Body
}

func (s *CheckUserAuthToMsceneResponse) SetHeaders(v map[string]*string) *CheckUserAuthToMsceneResponse {
	s.Headers = v
	return s
}

func (s *CheckUserAuthToMsceneResponse) SetStatusCode(v int32) *CheckUserAuthToMsceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserAuthToMsceneResponse) SetBody(v *CheckUserAuthToMsceneResponseBody) *CheckUserAuthToMsceneResponse {
	s.Body = v
	return s
}

func (s *CheckUserAuthToMsceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
