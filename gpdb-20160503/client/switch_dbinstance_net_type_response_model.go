// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceNetTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchDBInstanceNetTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchDBInstanceNetTypeResponse
	GetStatusCode() *int32
	SetBody(v *SwitchDBInstanceNetTypeResponseBody) *SwitchDBInstanceNetTypeResponse
	GetBody() *SwitchDBInstanceNetTypeResponseBody
}

type SwitchDBInstanceNetTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchDBInstanceNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchDBInstanceNetTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchDBInstanceNetTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchDBInstanceNetTypeResponse) GetBody() *SwitchDBInstanceNetTypeResponseBody {
	return s.Body
}

func (s *SwitchDBInstanceNetTypeResponse) SetHeaders(v map[string]*string) *SwitchDBInstanceNetTypeResponse {
	s.Headers = v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) SetStatusCode(v int32) *SwitchDBInstanceNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) SetBody(v *SwitchDBInstanceNetTypeResponseBody) *SwitchDBInstanceNetTypeResponse {
	s.Body = v
	return s
}

func (s *SwitchDBInstanceNetTypeResponse) Validate() error {
	return dara.Validate(s)
}
