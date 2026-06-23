// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectOperationTaskApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectOperationTaskApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectOperationTaskApprovalResponse
	GetStatusCode() *int32
	SetBody(v *RejectOperationTaskApprovalResponseBody) *RejectOperationTaskApprovalResponse
	GetBody() *RejectOperationTaskApprovalResponseBody
}

type RejectOperationTaskApprovalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectOperationTaskApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectOperationTaskApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectOperationTaskApprovalResponse) GoString() string {
	return s.String()
}

func (s *RejectOperationTaskApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectOperationTaskApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectOperationTaskApprovalResponse) GetBody() *RejectOperationTaskApprovalResponseBody {
	return s.Body
}

func (s *RejectOperationTaskApprovalResponse) SetHeaders(v map[string]*string) *RejectOperationTaskApprovalResponse {
	s.Headers = v
	return s
}

func (s *RejectOperationTaskApprovalResponse) SetStatusCode(v int32) *RejectOperationTaskApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectOperationTaskApprovalResponse) SetBody(v *RejectOperationTaskApprovalResponseBody) *RejectOperationTaskApprovalResponse {
	s.Body = v
	return s
}

func (s *RejectOperationTaskApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
