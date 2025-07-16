// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchMainOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchMainOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchMainOrgResponse
	GetStatusCode() *int32
	SetBody(v *SwitchMainOrgResponseBody) *SwitchMainOrgResponse
	GetBody() *SwitchMainOrgResponseBody
}

type SwitchMainOrgResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchMainOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchMainOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgResponse) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchMainOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchMainOrgResponse) GetBody() *SwitchMainOrgResponseBody {
	return s.Body
}

func (s *SwitchMainOrgResponse) SetHeaders(v map[string]*string) *SwitchMainOrgResponse {
	s.Headers = v
	return s
}

func (s *SwitchMainOrgResponse) SetStatusCode(v int32) *SwitchMainOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchMainOrgResponse) SetBody(v *SwitchMainOrgResponseBody) *SwitchMainOrgResponse {
	s.Body = v
	return s
}

func (s *SwitchMainOrgResponse) Validate() error {
	return dara.Validate(s)
}
