// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalProcessesForApprovalSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApprovalProcessesForApprovalSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApprovalProcessesForApprovalSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListApprovalProcessesForApprovalSchemasResponseBody) *ListApprovalProcessesForApprovalSchemasResponse
	GetBody() *ListApprovalProcessesForApprovalSchemasResponseBody
}

type ListApprovalProcessesForApprovalSchemasResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApprovalProcessesForApprovalSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApprovalProcessesForApprovalSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesForApprovalSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) GetBody() *ListApprovalProcessesForApprovalSchemasResponseBody {
	return s.Body
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) SetHeaders(v map[string]*string) *ListApprovalProcessesForApprovalSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) SetStatusCode(v int32) *ListApprovalProcessesForApprovalSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) SetBody(v *ListApprovalProcessesForApprovalSchemasResponseBody) *ListApprovalProcessesForApprovalSchemasResponse {
	s.Body = v
	return s
}

func (s *ListApprovalProcessesForApprovalSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
