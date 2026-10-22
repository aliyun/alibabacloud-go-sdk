// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAutoAuditUnblockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAutoAuditUnblockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAutoAuditUnblockResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAutoAuditUnblockResponseBody) *SubmitAutoAuditUnblockResponse
	GetBody() *SubmitAutoAuditUnblockResponseBody
}

type SubmitAutoAuditUnblockResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAutoAuditUnblockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAutoAuditUnblockResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAutoAuditUnblockResponse) GoString() string {
	return s.String()
}

func (s *SubmitAutoAuditUnblockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAutoAuditUnblockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAutoAuditUnblockResponse) GetBody() *SubmitAutoAuditUnblockResponseBody {
	return s.Body
}

func (s *SubmitAutoAuditUnblockResponse) SetHeaders(v map[string]*string) *SubmitAutoAuditUnblockResponse {
	s.Headers = v
	return s
}

func (s *SubmitAutoAuditUnblockResponse) SetStatusCode(v int32) *SubmitAutoAuditUnblockResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAutoAuditUnblockResponse) SetBody(v *SubmitAutoAuditUnblockResponseBody) *SubmitAutoAuditUnblockResponse {
	s.Body = v
	return s
}

func (s *SubmitAutoAuditUnblockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
