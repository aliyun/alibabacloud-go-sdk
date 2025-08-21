// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCCGlobalSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebCCGlobalSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebCCGlobalSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebCCGlobalSwitchResponseBody) *ModifyWebCCGlobalSwitchResponse
	GetBody() *ModifyWebCCGlobalSwitchResponseBody
}

type ModifyWebCCGlobalSwitchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebCCGlobalSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebCCGlobalSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCCGlobalSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCCGlobalSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebCCGlobalSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebCCGlobalSwitchResponse) GetBody() *ModifyWebCCGlobalSwitchResponseBody {
	return s.Body
}

func (s *ModifyWebCCGlobalSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebCCGlobalSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCCGlobalSwitchResponse) SetStatusCode(v int32) *ModifyWebCCGlobalSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebCCGlobalSwitchResponse) SetBody(v *ModifyWebCCGlobalSwitchResponseBody) *ModifyWebCCGlobalSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyWebCCGlobalSwitchResponse) Validate() error {
	return dara.Validate(s)
}
