// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicy2ApprovalProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPolicy2ApprovalProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPolicy2ApprovalProcessResponse
	GetStatusCode() *int32
	SetBody(v *AttachPolicy2ApprovalProcessResponseBody) *AttachPolicy2ApprovalProcessResponse
	GetBody() *AttachPolicy2ApprovalProcessResponseBody
}

type AttachPolicy2ApprovalProcessResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicy2ApprovalProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicy2ApprovalProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicy2ApprovalProcessResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicy2ApprovalProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPolicy2ApprovalProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPolicy2ApprovalProcessResponse) GetBody() *AttachPolicy2ApprovalProcessResponseBody {
	return s.Body
}

func (s *AttachPolicy2ApprovalProcessResponse) SetHeaders(v map[string]*string) *AttachPolicy2ApprovalProcessResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicy2ApprovalProcessResponse) SetStatusCode(v int32) *AttachPolicy2ApprovalProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicy2ApprovalProcessResponse) SetBody(v *AttachPolicy2ApprovalProcessResponseBody) *AttachPolicy2ApprovalProcessResponse {
	s.Body = v
	return s
}

func (s *AttachPolicy2ApprovalProcessResponse) Validate() error {
	return dara.Validate(s)
}
