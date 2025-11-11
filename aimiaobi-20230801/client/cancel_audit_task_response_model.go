// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAuditTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAuditTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAuditTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelAuditTaskResponseBody) *CancelAuditTaskResponse
	GetBody() *CancelAuditTaskResponseBody
}

type CancelAuditTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAuditTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAuditTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAuditTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelAuditTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAuditTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAuditTaskResponse) GetBody() *CancelAuditTaskResponseBody {
	return s.Body
}

func (s *CancelAuditTaskResponse) SetHeaders(v map[string]*string) *CancelAuditTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelAuditTaskResponse) SetStatusCode(v int32) *CancelAuditTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAuditTaskResponse) SetBody(v *CancelAuditTaskResponseBody) *CancelAuditTaskResponse {
	s.Body = v
	return s
}

func (s *CancelAuditTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
