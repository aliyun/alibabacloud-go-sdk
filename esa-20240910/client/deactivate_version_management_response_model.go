// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateVersionManagementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateVersionManagementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateVersionManagementResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateVersionManagementResponseBody) *DeactivateVersionManagementResponse
	GetBody() *DeactivateVersionManagementResponseBody
}

type DeactivateVersionManagementResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateVersionManagementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateVersionManagementResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateVersionManagementResponse) GoString() string {
	return s.String()
}

func (s *DeactivateVersionManagementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateVersionManagementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateVersionManagementResponse) GetBody() *DeactivateVersionManagementResponseBody {
	return s.Body
}

func (s *DeactivateVersionManagementResponse) SetHeaders(v map[string]*string) *DeactivateVersionManagementResponse {
	s.Headers = v
	return s
}

func (s *DeactivateVersionManagementResponse) SetStatusCode(v int32) *DeactivateVersionManagementResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateVersionManagementResponse) SetBody(v *DeactivateVersionManagementResponseBody) *DeactivateVersionManagementResponse {
	s.Body = v
	return s
}

func (s *DeactivateVersionManagementResponse) Validate() error {
	return dara.Validate(s)
}
