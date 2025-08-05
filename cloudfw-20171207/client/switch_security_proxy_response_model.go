// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSecurityProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchSecurityProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchSecurityProxyResponse
	GetStatusCode() *int32
	SetBody(v *SwitchSecurityProxyResponseBody) *SwitchSecurityProxyResponse
	GetBody() *SwitchSecurityProxyResponseBody
}

type SwitchSecurityProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchSecurityProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchSecurityProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchSecurityProxyResponse) GoString() string {
	return s.String()
}

func (s *SwitchSecurityProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchSecurityProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchSecurityProxyResponse) GetBody() *SwitchSecurityProxyResponseBody {
	return s.Body
}

func (s *SwitchSecurityProxyResponse) SetHeaders(v map[string]*string) *SwitchSecurityProxyResponse {
	s.Headers = v
	return s
}

func (s *SwitchSecurityProxyResponse) SetStatusCode(v int32) *SwitchSecurityProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchSecurityProxyResponse) SetBody(v *SwitchSecurityProxyResponseBody) *SwitchSecurityProxyResponse {
	s.Body = v
	return s
}

func (s *SwitchSecurityProxyResponse) Validate() error {
	return dara.Validate(s)
}
