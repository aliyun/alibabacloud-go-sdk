// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvalidateApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvalidateApprovalResponse
	GetStatusCode() *int32
	SetBody(v *InvalidateApprovalResponseBody) *InvalidateApprovalResponse
	GetBody() *InvalidateApprovalResponseBody
}

type InvalidateApprovalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidateApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidateApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s InvalidateApprovalResponse) GoString() string {
	return s.String()
}

func (s *InvalidateApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvalidateApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvalidateApprovalResponse) GetBody() *InvalidateApprovalResponseBody {
	return s.Body
}

func (s *InvalidateApprovalResponse) SetHeaders(v map[string]*string) *InvalidateApprovalResponse {
	s.Headers = v
	return s
}

func (s *InvalidateApprovalResponse) SetStatusCode(v int32) *InvalidateApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidateApprovalResponse) SetBody(v *InvalidateApprovalResponseBody) *InvalidateApprovalResponse {
	s.Body = v
	return s
}

func (s *InvalidateApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
