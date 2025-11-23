// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOrderApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitOrderApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitOrderApprovalResponse
	GetStatusCode() *int32
	SetBody(v *SubmitOrderApprovalResponseBody) *SubmitOrderApprovalResponse
	GetBody() *SubmitOrderApprovalResponseBody
}

type SubmitOrderApprovalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitOrderApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitOrderApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitOrderApprovalResponse) GoString() string {
	return s.String()
}

func (s *SubmitOrderApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitOrderApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitOrderApprovalResponse) GetBody() *SubmitOrderApprovalResponseBody {
	return s.Body
}

func (s *SubmitOrderApprovalResponse) SetHeaders(v map[string]*string) *SubmitOrderApprovalResponse {
	s.Headers = v
	return s
}

func (s *SubmitOrderApprovalResponse) SetStatusCode(v int32) *SubmitOrderApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOrderApprovalResponse) SetBody(v *SubmitOrderApprovalResponseBody) *SubmitOrderApprovalResponse {
	s.Body = v
	return s
}

func (s *SubmitOrderApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
