// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAuditProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterAuditProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterAuditProjectResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterAuditProjectResponseBody) *GetClusterAuditProjectResponse
	GetBody() *GetClusterAuditProjectResponseBody
}

type GetClusterAuditProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterAuditProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterAuditProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAuditProjectResponse) GoString() string {
	return s.String()
}

func (s *GetClusterAuditProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterAuditProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterAuditProjectResponse) GetBody() *GetClusterAuditProjectResponseBody {
	return s.Body
}

func (s *GetClusterAuditProjectResponse) SetHeaders(v map[string]*string) *GetClusterAuditProjectResponse {
	s.Headers = v
	return s
}

func (s *GetClusterAuditProjectResponse) SetStatusCode(v int32) *GetClusterAuditProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterAuditProjectResponse) SetBody(v *GetClusterAuditProjectResponseBody) *GetClusterAuditProjectResponse {
	s.Body = v
	return s
}

func (s *GetClusterAuditProjectResponse) Validate() error {
	return dara.Validate(s)
}
