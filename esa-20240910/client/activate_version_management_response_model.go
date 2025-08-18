// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateVersionManagementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateVersionManagementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateVersionManagementResponse
	GetStatusCode() *int32
	SetBody(v *ActivateVersionManagementResponseBody) *ActivateVersionManagementResponse
	GetBody() *ActivateVersionManagementResponseBody
}

type ActivateVersionManagementResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateVersionManagementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateVersionManagementResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateVersionManagementResponse) GoString() string {
	return s.String()
}

func (s *ActivateVersionManagementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateVersionManagementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateVersionManagementResponse) GetBody() *ActivateVersionManagementResponseBody {
	return s.Body
}

func (s *ActivateVersionManagementResponse) SetHeaders(v map[string]*string) *ActivateVersionManagementResponse {
	s.Headers = v
	return s
}

func (s *ActivateVersionManagementResponse) SetStatusCode(v int32) *ActivateVersionManagementResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateVersionManagementResponse) SetBody(v *ActivateVersionManagementResponseBody) *ActivateVersionManagementResponse {
	s.Body = v
	return s
}

func (s *ActivateVersionManagementResponse) Validate() error {
	return dara.Validate(s)
}
