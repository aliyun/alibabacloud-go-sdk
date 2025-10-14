// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSiteMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSiteMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSiteMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *DisableSiteMonitorsResponseBody) *DisableSiteMonitorsResponse
	GetBody() *DisableSiteMonitorsResponseBody
}

type DisableSiteMonitorsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSiteMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSiteMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSiteMonitorsResponse) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSiteMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSiteMonitorsResponse) GetBody() *DisableSiteMonitorsResponseBody {
	return s.Body
}

func (s *DisableSiteMonitorsResponse) SetHeaders(v map[string]*string) *DisableSiteMonitorsResponse {
	s.Headers = v
	return s
}

func (s *DisableSiteMonitorsResponse) SetStatusCode(v int32) *DisableSiteMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSiteMonitorsResponse) SetBody(v *DisableSiteMonitorsResponseBody) *DisableSiteMonitorsResponse {
	s.Body = v
	return s
}

func (s *DisableSiteMonitorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
