// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRestoreTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRestoreTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRestoreTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartRestoreTaskResponseBody) *StartRestoreTaskResponse
	GetBody() *StartRestoreTaskResponseBody
}

type StartRestoreTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRestoreTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRestoreTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRestoreTaskResponse) GoString() string {
	return s.String()
}

func (s *StartRestoreTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRestoreTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRestoreTaskResponse) GetBody() *StartRestoreTaskResponseBody {
	return s.Body
}

func (s *StartRestoreTaskResponse) SetHeaders(v map[string]*string) *StartRestoreTaskResponse {
	s.Headers = v
	return s
}

func (s *StartRestoreTaskResponse) SetStatusCode(v int32) *StartRestoreTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRestoreTaskResponse) SetBody(v *StartRestoreTaskResponseBody) *StartRestoreTaskResponse {
	s.Body = v
	return s
}

func (s *StartRestoreTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
