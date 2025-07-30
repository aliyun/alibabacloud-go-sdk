// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchNetworkResponse
	GetStatusCode() *int32
	SetBody(v *SwitchNetworkResponseBody) *SwitchNetworkResponse
	GetBody() *SwitchNetworkResponseBody
}

type SwitchNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchNetworkResponse) GoString() string {
	return s.String()
}

func (s *SwitchNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchNetworkResponse) GetBody() *SwitchNetworkResponseBody {
	return s.Body
}

func (s *SwitchNetworkResponse) SetHeaders(v map[string]*string) *SwitchNetworkResponse {
	s.Headers = v
	return s
}

func (s *SwitchNetworkResponse) SetStatusCode(v int32) *SwitchNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchNetworkResponse) SetBody(v *SwitchNetworkResponseBody) *SwitchNetworkResponse {
	s.Body = v
	return s
}

func (s *SwitchNetworkResponse) Validate() error {
	return dara.Validate(s)
}
