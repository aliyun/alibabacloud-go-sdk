// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupByResourceGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorGroupByResourceGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorGroupByResourceGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorGroupByResourceGroupIdResponseBody) *CreateMonitorGroupByResourceGroupIdResponse
	GetBody() *CreateMonitorGroupByResourceGroupIdResponseBody
}

type CreateMonitorGroupByResourceGroupIdResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorGroupByResourceGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorGroupByResourceGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupByResourceGroupIdResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) GetBody() *CreateMonitorGroupByResourceGroupIdResponseBody {
	return s.Body
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupByResourceGroupIdResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) SetStatusCode(v int32) *CreateMonitorGroupByResourceGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) SetBody(v *CreateMonitorGroupByResourceGroupIdResponseBody) *CreateMonitorGroupByResourceGroupIdResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) Validate() error {
	return dara.Validate(s)
}
