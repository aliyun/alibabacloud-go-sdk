// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchInstanceProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchInstanceProxyResponse
	GetStatusCode() *int32
	SetBody(v *SwitchInstanceProxyResponseBody) *SwitchInstanceProxyResponse
	GetBody() *SwitchInstanceProxyResponseBody
}

type SwitchInstanceProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchInstanceProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchInstanceProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceProxyResponse) GoString() string {
	return s.String()
}

func (s *SwitchInstanceProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchInstanceProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchInstanceProxyResponse) GetBody() *SwitchInstanceProxyResponseBody {
	return s.Body
}

func (s *SwitchInstanceProxyResponse) SetHeaders(v map[string]*string) *SwitchInstanceProxyResponse {
	s.Headers = v
	return s
}

func (s *SwitchInstanceProxyResponse) SetStatusCode(v int32) *SwitchInstanceProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchInstanceProxyResponse) SetBody(v *SwitchInstanceProxyResponseBody) *SwitchInstanceProxyResponse {
	s.Body = v
	return s
}

func (s *SwitchInstanceProxyResponse) Validate() error {
	return dara.Validate(s)
}
