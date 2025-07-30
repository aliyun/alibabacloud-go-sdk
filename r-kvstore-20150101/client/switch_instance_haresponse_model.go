// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceHAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchInstanceHAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchInstanceHAResponse
	GetStatusCode() *int32
	SetBody(v *SwitchInstanceHAResponseBody) *SwitchInstanceHAResponse
	GetBody() *SwitchInstanceHAResponseBody
}

type SwitchInstanceHAResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchInstanceHAResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceHAResponse) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchInstanceHAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchInstanceHAResponse) GetBody() *SwitchInstanceHAResponseBody {
	return s.Body
}

func (s *SwitchInstanceHAResponse) SetHeaders(v map[string]*string) *SwitchInstanceHAResponse {
	s.Headers = v
	return s
}

func (s *SwitchInstanceHAResponse) SetStatusCode(v int32) *SwitchInstanceHAResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchInstanceHAResponse) SetBody(v *SwitchInstanceHAResponseBody) *SwitchInstanceHAResponse {
	s.Body = v
	return s
}

func (s *SwitchInstanceHAResponse) Validate() error {
	return dara.Validate(s)
}
