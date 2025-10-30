// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApprovalProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApprovalProcessResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApprovalProcessResponseBody) *UpdateApprovalProcessResponse
	GetBody() *UpdateApprovalProcessResponseBody
}

type UpdateApprovalProcessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApprovalProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApprovalProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponse) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApprovalProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApprovalProcessResponse) GetBody() *UpdateApprovalProcessResponseBody {
	return s.Body
}

func (s *UpdateApprovalProcessResponse) SetHeaders(v map[string]*string) *UpdateApprovalProcessResponse {
	s.Headers = v
	return s
}

func (s *UpdateApprovalProcessResponse) SetStatusCode(v int32) *UpdateApprovalProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApprovalProcessResponse) SetBody(v *UpdateApprovalProcessResponseBody) *UpdateApprovalProcessResponse {
	s.Body = v
	return s
}

func (s *UpdateApprovalProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
