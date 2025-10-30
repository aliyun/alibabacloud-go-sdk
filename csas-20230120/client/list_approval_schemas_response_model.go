// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApprovalSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApprovalSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListApprovalSchemasResponseBody) *ListApprovalSchemasResponse
	GetBody() *ListApprovalSchemasResponseBody
}

type ListApprovalSchemasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApprovalSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApprovalSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListApprovalSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApprovalSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApprovalSchemasResponse) GetBody() *ListApprovalSchemasResponseBody {
	return s.Body
}

func (s *ListApprovalSchemasResponse) SetHeaders(v map[string]*string) *ListApprovalSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListApprovalSchemasResponse) SetStatusCode(v int32) *ListApprovalSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApprovalSchemasResponse) SetBody(v *ListApprovalSchemasResponseBody) *ListApprovalSchemasResponse {
	s.Body = v
	return s
}

func (s *ListApprovalSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
