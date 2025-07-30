// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceZoneFailOverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchInstanceZoneFailOverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchInstanceZoneFailOverResponse
	GetStatusCode() *int32
	SetBody(v *SwitchInstanceZoneFailOverResponseBody) *SwitchInstanceZoneFailOverResponse
	GetBody() *SwitchInstanceZoneFailOverResponseBody
}

type SwitchInstanceZoneFailOverResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchInstanceZoneFailOverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchInstanceZoneFailOverResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceZoneFailOverResponse) GoString() string {
	return s.String()
}

func (s *SwitchInstanceZoneFailOverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchInstanceZoneFailOverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchInstanceZoneFailOverResponse) GetBody() *SwitchInstanceZoneFailOverResponseBody {
	return s.Body
}

func (s *SwitchInstanceZoneFailOverResponse) SetHeaders(v map[string]*string) *SwitchInstanceZoneFailOverResponse {
	s.Headers = v
	return s
}

func (s *SwitchInstanceZoneFailOverResponse) SetStatusCode(v int32) *SwitchInstanceZoneFailOverResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchInstanceZoneFailOverResponse) SetBody(v *SwitchInstanceZoneFailOverResponseBody) *SwitchInstanceZoneFailOverResponse {
	s.Body = v
	return s
}

func (s *SwitchInstanceZoneFailOverResponse) Validate() error {
	return dara.Validate(s)
}
