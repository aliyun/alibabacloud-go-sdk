// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProjectTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitProjectTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitProjectTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitProjectTaskResponseBody) *SubmitProjectTaskResponse
	GetBody() *SubmitProjectTaskResponseBody
}

type SubmitProjectTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitProjectTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitProjectTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitProjectTaskResponse) GetBody() *SubmitProjectTaskResponseBody {
	return s.Body
}

func (s *SubmitProjectTaskResponse) SetHeaders(v map[string]*string) *SubmitProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitProjectTaskResponse) SetStatusCode(v int32) *SubmitProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitProjectTaskResponse) SetBody(v *SubmitProjectTaskResponseBody) *SubmitProjectTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitProjectTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
