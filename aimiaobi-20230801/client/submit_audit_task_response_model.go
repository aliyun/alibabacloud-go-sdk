// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAuditTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAuditTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAuditTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAuditTaskResponseBody) *SubmitAuditTaskResponse
	GetBody() *SubmitAuditTaskResponseBody
}

type SubmitAuditTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAuditTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAuditTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAuditTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAuditTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAuditTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAuditTaskResponse) GetBody() *SubmitAuditTaskResponseBody {
	return s.Body
}

func (s *SubmitAuditTaskResponse) SetHeaders(v map[string]*string) *SubmitAuditTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAuditTaskResponse) SetStatusCode(v int32) *SubmitAuditTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAuditTaskResponse) SetBody(v *SubmitAuditTaskResponseBody) *SubmitAuditTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitAuditTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
