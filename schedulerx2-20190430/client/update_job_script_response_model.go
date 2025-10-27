// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobScriptResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobScriptResponseBody) *UpdateJobScriptResponse
	GetBody() *UpdateJobScriptResponseBody
}

type UpdateJobScriptResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobScriptResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobScriptResponse) GetBody() *UpdateJobScriptResponseBody {
	return s.Body
}

func (s *UpdateJobScriptResponse) SetHeaders(v map[string]*string) *UpdateJobScriptResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobScriptResponse) SetStatusCode(v int32) *UpdateJobScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobScriptResponse) SetBody(v *UpdateJobScriptResponseBody) *UpdateJobScriptResponse {
	s.Body = v
	return s
}

func (s *UpdateJobScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
