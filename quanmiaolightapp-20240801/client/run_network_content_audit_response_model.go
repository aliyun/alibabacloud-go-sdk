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
	SetId(v string) *RunNetworkContentAuditResponse
	GetId() *string
	SetEvent(v string) *RunNetworkContentAuditResponse
	GetEvent() *string
	SetBody(v *RunNetworkContentAuditResponseBody) *RunNetworkContentAuditResponse
	GetBody() *RunNetworkContentAuditResponseBody
}

type RunNetworkContentAuditResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
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

func (s *RunNetworkContentAuditResponse) GetId() *string {
	return s.Id
}

func (s *RunNetworkContentAuditResponse) GetEvent() *string {
	return s.Event
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

func (s *RunNetworkContentAuditResponse) SetId(v string) *RunNetworkContentAuditResponse {
	s.Id = &v
	return s
}

func (s *RunNetworkContentAuditResponse) SetEvent(v string) *RunNetworkContentAuditResponse {
	s.Event = &v
	return s
}

func (s *RunNetworkContentAuditResponse) SetBody(v *RunNetworkContentAuditResponseBody) *RunNetworkContentAuditResponse {
	s.Body = v
	return s
}

func (s *RunNetworkContentAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
