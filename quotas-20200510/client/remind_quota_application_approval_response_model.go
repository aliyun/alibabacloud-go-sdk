// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemindQuotaApplicationApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemindQuotaApplicationApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemindQuotaApplicationApprovalResponse
	GetStatusCode() *int32
	SetBody(v *RemindQuotaApplicationApprovalResponseBody) *RemindQuotaApplicationApprovalResponse
	GetBody() *RemindQuotaApplicationApprovalResponseBody
}

type RemindQuotaApplicationApprovalResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemindQuotaApplicationApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemindQuotaApplicationApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s RemindQuotaApplicationApprovalResponse) GoString() string {
	return s.String()
}

func (s *RemindQuotaApplicationApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemindQuotaApplicationApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemindQuotaApplicationApprovalResponse) GetBody() *RemindQuotaApplicationApprovalResponseBody {
	return s.Body
}

func (s *RemindQuotaApplicationApprovalResponse) SetHeaders(v map[string]*string) *RemindQuotaApplicationApprovalResponse {
	s.Headers = v
	return s
}

func (s *RemindQuotaApplicationApprovalResponse) SetStatusCode(v int32) *RemindQuotaApplicationApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *RemindQuotaApplicationApprovalResponse) SetBody(v *RemindQuotaApplicationApprovalResponseBody) *RemindQuotaApplicationApprovalResponse {
	s.Body = v
	return s
}

func (s *RemindQuotaApplicationApprovalResponse) Validate() error {
	return dara.Validate(s)
}
