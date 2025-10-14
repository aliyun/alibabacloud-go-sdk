// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSiteMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSiteMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateSiteMonitorResponseBody) *CreateSiteMonitorResponse
	GetBody() *CreateSiteMonitorResponseBody
}

type CreateSiteMonitorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSiteMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSiteMonitorResponse) GetBody() *CreateSiteMonitorResponseBody {
	return s.Body
}

func (s *CreateSiteMonitorResponse) SetHeaders(v map[string]*string) *CreateSiteMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteMonitorResponse) SetStatusCode(v int32) *CreateSiteMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteMonitorResponse) SetBody(v *CreateSiteMonitorResponseBody) *CreateSiteMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateSiteMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
