// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNetworkContentAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunNetworkContentAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunNetworkContentAuditResponse
	GetStatusCode() *int32
	SetBody(v *RunNetworkContentAuditResponseBody) *RunNetworkContentAuditResponse
	GetBody() *RunNetworkContentAuditResponseBody
}

type RunNetworkContentAuditResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNetworkContentAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNetworkContentAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditResponse) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunNetworkContentAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunNetworkContentAuditResponse) GetBody() *RunNetworkContentAuditResponseBody {
	return s.Body
}

func (s *RunNetworkContentAuditResponse) SetHeaders(v map[string]*string) *RunNetworkContentAuditResponse {
	s.Headers = v
	return s
}

func (s *RunNetworkContentAuditResponse) SetStatusCode(v int32) *RunNetworkContentAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNetworkContentAuditResponse) SetBody(v *RunNetworkContentAuditResponseBody) *RunNetworkContentAuditResponse {
	s.Body = v
	return s
}

func (s *RunNetworkContentAuditResponse) Validate() error {
	return dara.Validate(s)
}
