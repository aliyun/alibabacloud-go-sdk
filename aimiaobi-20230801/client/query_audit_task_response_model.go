// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuditTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAuditTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAuditTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryAuditTaskResponseBody) *QueryAuditTaskResponse
	GetBody() *QueryAuditTaskResponseBody
}

type QueryAuditTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuditTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuditTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAuditTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAuditTaskResponse) GetBody() *QueryAuditTaskResponseBody {
	return s.Body
}

func (s *QueryAuditTaskResponse) SetHeaders(v map[string]*string) *QueryAuditTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryAuditTaskResponse) SetStatusCode(v int32) *QueryAuditTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuditTaskResponse) SetBody(v *QueryAuditTaskResponseBody) *QueryAuditTaskResponse {
	s.Body = v
	return s
}

func (s *QueryAuditTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
