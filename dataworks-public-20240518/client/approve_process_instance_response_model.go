// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveProcessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveProcessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveProcessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ApproveProcessInstanceResponseBody) *ApproveProcessInstanceResponse
	GetBody() *ApproveProcessInstanceResponseBody
}

type ApproveProcessInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveProcessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *ApproveProcessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveProcessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveProcessInstanceResponse) GetBody() *ApproveProcessInstanceResponseBody {
	return s.Body
}

func (s *ApproveProcessInstanceResponse) SetHeaders(v map[string]*string) *ApproveProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *ApproveProcessInstanceResponse) SetStatusCode(v int32) *ApproveProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveProcessInstanceResponse) SetBody(v *ApproveProcessInstanceResponseBody) *ApproveProcessInstanceResponse {
	s.Body = v
	return s
}

func (s *ApproveProcessInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
