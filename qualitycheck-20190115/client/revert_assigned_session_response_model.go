// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAssignedSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertAssignedSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertAssignedSessionResponse
	GetStatusCode() *int32
	SetBody(v *RevertAssignedSessionResponseBody) *RevertAssignedSessionResponse
	GetBody() *RevertAssignedSessionResponseBody
}

type RevertAssignedSessionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertAssignedSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertAssignedSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionResponse) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertAssignedSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertAssignedSessionResponse) GetBody() *RevertAssignedSessionResponseBody {
	return s.Body
}

func (s *RevertAssignedSessionResponse) SetHeaders(v map[string]*string) *RevertAssignedSessionResponse {
	s.Headers = v
	return s
}

func (s *RevertAssignedSessionResponse) SetStatusCode(v int32) *RevertAssignedSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertAssignedSessionResponse) SetBody(v *RevertAssignedSessionResponseBody) *RevertAssignedSessionResponse {
	s.Body = v
	return s
}

func (s *RevertAssignedSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
