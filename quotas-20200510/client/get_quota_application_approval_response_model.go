// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaApplicationApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaApplicationApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaApplicationApprovalResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaApplicationApprovalResponseBody) *GetQuotaApplicationApprovalResponse
	GetBody() *GetQuotaApplicationApprovalResponseBody
}

type GetQuotaApplicationApprovalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaApplicationApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaApplicationApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationApprovalResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaApplicationApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaApplicationApprovalResponse) GetBody() *GetQuotaApplicationApprovalResponseBody {
	return s.Body
}

func (s *GetQuotaApplicationApprovalResponse) SetHeaders(v map[string]*string) *GetQuotaApplicationApprovalResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaApplicationApprovalResponse) SetStatusCode(v int32) *GetQuotaApplicationApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponse) SetBody(v *GetQuotaApplicationApprovalResponseBody) *GetQuotaApplicationApprovalResponse {
	s.Body = v
	return s
}

func (s *GetQuotaApplicationApprovalResponse) Validate() error {
	return dara.Validate(s)
}
