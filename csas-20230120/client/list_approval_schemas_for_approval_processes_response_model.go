// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalSchemasForApprovalProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApprovalSchemasForApprovalProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApprovalSchemasForApprovalProcessesResponse
	GetStatusCode() *int32
	SetBody(v *ListApprovalSchemasForApprovalProcessesResponseBody) *ListApprovalSchemasForApprovalProcessesResponse
	GetBody() *ListApprovalSchemasForApprovalProcessesResponseBody
}

type ListApprovalSchemasForApprovalProcessesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApprovalSchemasForApprovalProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApprovalSchemasForApprovalProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasForApprovalProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) GetBody() *ListApprovalSchemasForApprovalProcessesResponseBody {
	return s.Body
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) SetHeaders(v map[string]*string) *ListApprovalSchemasForApprovalProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) SetStatusCode(v int32) *ListApprovalSchemasForApprovalProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) SetBody(v *ListApprovalSchemasForApprovalProcessesResponseBody) *ListApprovalSchemasForApprovalProcessesResponse {
	s.Body = v
	return s
}

func (s *ListApprovalSchemasForApprovalProcessesResponse) Validate() error {
	return dara.Validate(s)
}
