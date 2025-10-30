// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicy2ApprovalProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPolicy2ApprovalProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPolicy2ApprovalProcessResponse
	GetStatusCode() *int32
	SetBody(v *DetachPolicy2ApprovalProcessResponseBody) *DetachPolicy2ApprovalProcessResponse
	GetBody() *DetachPolicy2ApprovalProcessResponseBody
}

type DetachPolicy2ApprovalProcessResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicy2ApprovalProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicy2ApprovalProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicy2ApprovalProcessResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicy2ApprovalProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPolicy2ApprovalProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPolicy2ApprovalProcessResponse) GetBody() *DetachPolicy2ApprovalProcessResponseBody {
	return s.Body
}

func (s *DetachPolicy2ApprovalProcessResponse) SetHeaders(v map[string]*string) *DetachPolicy2ApprovalProcessResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicy2ApprovalProcessResponse) SetStatusCode(v int32) *DetachPolicy2ApprovalProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicy2ApprovalProcessResponse) SetBody(v *DetachPolicy2ApprovalProcessResponseBody) *DetachPolicy2ApprovalProcessResponse {
	s.Body = v
	return s
}

func (s *DetachPolicy2ApprovalProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
