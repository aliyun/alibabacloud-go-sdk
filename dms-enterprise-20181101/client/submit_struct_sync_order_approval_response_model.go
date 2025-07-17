// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitStructSyncOrderApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitStructSyncOrderApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitStructSyncOrderApprovalResponse
	GetStatusCode() *int32
	SetBody(v *SubmitStructSyncOrderApprovalResponseBody) *SubmitStructSyncOrderApprovalResponse
	GetBody() *SubmitStructSyncOrderApprovalResponseBody
}

type SubmitStructSyncOrderApprovalResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitStructSyncOrderApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitStructSyncOrderApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitStructSyncOrderApprovalResponse) GoString() string {
	return s.String()
}

func (s *SubmitStructSyncOrderApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitStructSyncOrderApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitStructSyncOrderApprovalResponse) GetBody() *SubmitStructSyncOrderApprovalResponseBody {
	return s.Body
}

func (s *SubmitStructSyncOrderApprovalResponse) SetHeaders(v map[string]*string) *SubmitStructSyncOrderApprovalResponse {
	s.Headers = v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponse) SetStatusCode(v int32) *SubmitStructSyncOrderApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponse) SetBody(v *SubmitStructSyncOrderApprovalResponseBody) *SubmitStructSyncOrderApprovalResponse {
	s.Body = v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponse) Validate() error {
	return dara.Validate(s)
}
