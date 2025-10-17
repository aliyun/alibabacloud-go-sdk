// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchPhysicalDtsJobToCloudResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchPhysicalDtsJobToCloudResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchPhysicalDtsJobToCloudResponse
	GetStatusCode() *int32
	SetBody(v *SwitchPhysicalDtsJobToCloudResponseBody) *SwitchPhysicalDtsJobToCloudResponse
	GetBody() *SwitchPhysicalDtsJobToCloudResponseBody
}

type SwitchPhysicalDtsJobToCloudResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchPhysicalDtsJobToCloudResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchPhysicalDtsJobToCloudResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchPhysicalDtsJobToCloudResponse) GoString() string {
	return s.String()
}

func (s *SwitchPhysicalDtsJobToCloudResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchPhysicalDtsJobToCloudResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchPhysicalDtsJobToCloudResponse) GetBody() *SwitchPhysicalDtsJobToCloudResponseBody {
	return s.Body
}

func (s *SwitchPhysicalDtsJobToCloudResponse) SetHeaders(v map[string]*string) *SwitchPhysicalDtsJobToCloudResponse {
	s.Headers = v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponse) SetStatusCode(v int32) *SwitchPhysicalDtsJobToCloudResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponse) SetBody(v *SwitchPhysicalDtsJobToCloudResponseBody) *SwitchPhysicalDtsJobToCloudResponse {
	s.Body = v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
