// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationAuditInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitOperationAuditInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitOperationAuditInfoResponse
	GetStatusCode() *int32
	SetBody(v *SubmitOperationAuditInfoResponseBody) *SubmitOperationAuditInfoResponse
	GetBody() *SubmitOperationAuditInfoResponseBody
}

type SubmitOperationAuditInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitOperationAuditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitOperationAuditInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationAuditInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitOperationAuditInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitOperationAuditInfoResponse) GetBody() *SubmitOperationAuditInfoResponseBody {
	return s.Body
}

func (s *SubmitOperationAuditInfoResponse) SetHeaders(v map[string]*string) *SubmitOperationAuditInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitOperationAuditInfoResponse) SetStatusCode(v int32) *SubmitOperationAuditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOperationAuditInfoResponse) SetBody(v *SubmitOperationAuditInfoResponseBody) *SubmitOperationAuditInfoResponse {
	s.Body = v
	return s
}

func (s *SubmitOperationAuditInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
