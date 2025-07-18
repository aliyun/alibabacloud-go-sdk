// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApprovalProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApprovalProcessesResponse
	GetStatusCode() *int32
	SetBody(v *ListApprovalProcessesResponseBody) *ListApprovalProcessesResponse
	GetBody() *ListApprovalProcessesResponseBody
}

type ListApprovalProcessesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApprovalProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApprovalProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApprovalProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApprovalProcessesResponse) GetBody() *ListApprovalProcessesResponseBody {
	return s.Body
}

func (s *ListApprovalProcessesResponse) SetHeaders(v map[string]*string) *ListApprovalProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListApprovalProcessesResponse) SetStatusCode(v int32) *ListApprovalProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApprovalProcessesResponse) SetBody(v *ListApprovalProcessesResponseBody) *ListApprovalProcessesResponse {
	s.Body = v
	return s
}

func (s *ListApprovalProcessesResponse) Validate() error {
	return dara.Validate(s)
}
