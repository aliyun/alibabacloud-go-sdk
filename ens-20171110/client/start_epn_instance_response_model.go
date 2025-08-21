// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartEpnInstanceResponseBody) *StartEpnInstanceResponse
	GetBody() *StartEpnInstanceResponseBody
}

type StartEpnInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartEpnInstanceResponse) GetBody() *StartEpnInstanceResponseBody {
	return s.Body
}

func (s *StartEpnInstanceResponse) SetHeaders(v map[string]*string) *StartEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartEpnInstanceResponse) SetStatusCode(v int32) *StartEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEpnInstanceResponse) SetBody(v *StartEpnInstanceResponseBody) *StartEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *StartEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
