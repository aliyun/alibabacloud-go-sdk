// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllowOperationTaskApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllowOperationTaskApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllowOperationTaskApprovalResponse
	GetStatusCode() *int32
	SetBody(v *AllowOperationTaskApprovalResponseBody) *AllowOperationTaskApprovalResponse
	GetBody() *AllowOperationTaskApprovalResponseBody
}

type AllowOperationTaskApprovalResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllowOperationTaskApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllowOperationTaskApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s AllowOperationTaskApprovalResponse) GoString() string {
	return s.String()
}

func (s *AllowOperationTaskApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllowOperationTaskApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllowOperationTaskApprovalResponse) GetBody() *AllowOperationTaskApprovalResponseBody {
	return s.Body
}

func (s *AllowOperationTaskApprovalResponse) SetHeaders(v map[string]*string) *AllowOperationTaskApprovalResponse {
	s.Headers = v
	return s
}

func (s *AllowOperationTaskApprovalResponse) SetStatusCode(v int32) *AllowOperationTaskApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *AllowOperationTaskApprovalResponse) SetBody(v *AllowOperationTaskApprovalResponseBody) *AllowOperationTaskApprovalResponse {
	s.Body = v
	return s
}

func (s *AllowOperationTaskApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
