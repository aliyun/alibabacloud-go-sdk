// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProcessApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProcessApprovalResponse
	GetStatusCode() *int32
	SetBody(v *ProcessApprovalResponseBody) *ProcessApprovalResponse
	GetBody() *ProcessApprovalResponseBody
}

type ProcessApprovalResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProcessApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProcessApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s ProcessApprovalResponse) GoString() string {
	return s.String()
}

func (s *ProcessApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProcessApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProcessApprovalResponse) GetBody() *ProcessApprovalResponseBody {
	return s.Body
}

func (s *ProcessApprovalResponse) SetHeaders(v map[string]*string) *ProcessApprovalResponse {
	s.Headers = v
	return s
}

func (s *ProcessApprovalResponse) SetStatusCode(v int32) *ProcessApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessApprovalResponse) SetBody(v *ProcessApprovalResponseBody) *ProcessApprovalResponse {
	s.Body = v
	return s
}

func (s *ProcessApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
