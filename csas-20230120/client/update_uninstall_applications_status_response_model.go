// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUninstallApplicationsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUninstallApplicationsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUninstallApplicationsStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUninstallApplicationsStatusResponseBody) *UpdateUninstallApplicationsStatusResponse
	GetBody() *UpdateUninstallApplicationsStatusResponseBody
}

type UpdateUninstallApplicationsStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUninstallApplicationsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUninstallApplicationsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUninstallApplicationsStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUninstallApplicationsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUninstallApplicationsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUninstallApplicationsStatusResponse) GetBody() *UpdateUninstallApplicationsStatusResponseBody {
	return s.Body
}

func (s *UpdateUninstallApplicationsStatusResponse) SetHeaders(v map[string]*string) *UpdateUninstallApplicationsStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponse) SetStatusCode(v int32) *UpdateUninstallApplicationsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponse) SetBody(v *UpdateUninstallApplicationsStatusResponseBody) *UpdateUninstallApplicationsStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponse) Validate() error {
	return dara.Validate(s)
}
