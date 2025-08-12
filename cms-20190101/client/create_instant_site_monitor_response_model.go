// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstantSiteMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstantSiteMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstantSiteMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstantSiteMonitorResponseBody) *CreateInstantSiteMonitorResponse
	GetBody() *CreateInstantSiteMonitorResponseBody
}

type CreateInstantSiteMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstantSiteMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstantSiteMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstantSiteMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateInstantSiteMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstantSiteMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstantSiteMonitorResponse) GetBody() *CreateInstantSiteMonitorResponseBody {
	return s.Body
}

func (s *CreateInstantSiteMonitorResponse) SetHeaders(v map[string]*string) *CreateInstantSiteMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateInstantSiteMonitorResponse) SetStatusCode(v int32) *CreateInstantSiteMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstantSiteMonitorResponse) SetBody(v *CreateInstantSiteMonitorResponseBody) *CreateInstantSiteMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateInstantSiteMonitorResponse) Validate() error {
	return dara.Validate(s)
}
