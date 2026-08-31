// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalTasksByUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApprovalTasksByUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApprovalTasksByUserResponse
	GetStatusCode() *int32
	SetBody(v *ListApprovalTasksByUserResponseBody) *ListApprovalTasksByUserResponse
	GetBody() *ListApprovalTasksByUserResponseBody
}

type ListApprovalTasksByUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApprovalTasksByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApprovalTasksByUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserResponse) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApprovalTasksByUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApprovalTasksByUserResponse) GetBody() *ListApprovalTasksByUserResponseBody {
	return s.Body
}

func (s *ListApprovalTasksByUserResponse) SetHeaders(v map[string]*string) *ListApprovalTasksByUserResponse {
	s.Headers = v
	return s
}

func (s *ListApprovalTasksByUserResponse) SetStatusCode(v int32) *ListApprovalTasksByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApprovalTasksByUserResponse) SetBody(v *ListApprovalTasksByUserResponseBody) *ListApprovalTasksByUserResponse {
	s.Body = v
	return s
}

func (s *ListApprovalTasksByUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
