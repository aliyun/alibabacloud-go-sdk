// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApprovalProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApprovalProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApprovalProcessesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApprovalProcessesResponseBody) *DeleteApprovalProcessesResponse
	GetBody() *DeleteApprovalProcessesResponseBody
}

type DeleteApprovalProcessesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApprovalProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApprovalProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApprovalProcessesResponse) GoString() string {
	return s.String()
}

func (s *DeleteApprovalProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApprovalProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApprovalProcessesResponse) GetBody() *DeleteApprovalProcessesResponseBody {
	return s.Body
}

func (s *DeleteApprovalProcessesResponse) SetHeaders(v map[string]*string) *DeleteApprovalProcessesResponse {
	s.Headers = v
	return s
}

func (s *DeleteApprovalProcessesResponse) SetStatusCode(v int32) *DeleteApprovalProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApprovalProcessesResponse) SetBody(v *DeleteApprovalProcessesResponseBody) *DeleteApprovalProcessesResponse {
	s.Body = v
	return s
}

func (s *DeleteApprovalProcessesResponse) Validate() error {
	return dara.Validate(s)
}
