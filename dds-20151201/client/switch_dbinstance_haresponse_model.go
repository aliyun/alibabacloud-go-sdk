// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceHAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchDBInstanceHAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchDBInstanceHAResponse
	GetStatusCode() *int32
	SetBody(v *SwitchDBInstanceHAResponseBody) *SwitchDBInstanceHAResponse
	GetBody() *SwitchDBInstanceHAResponseBody
}

type SwitchDBInstanceHAResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchDBInstanceHAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchDBInstanceHAResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceHAResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchDBInstanceHAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchDBInstanceHAResponse) GetBody() *SwitchDBInstanceHAResponseBody {
	return s.Body
}

func (s *SwitchDBInstanceHAResponse) SetHeaders(v map[string]*string) *SwitchDBInstanceHAResponse {
	s.Headers = v
	return s
}

func (s *SwitchDBInstanceHAResponse) SetStatusCode(v int32) *SwitchDBInstanceHAResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDBInstanceHAResponse) SetBody(v *SwitchDBInstanceHAResponseBody) *SwitchDBInstanceHAResponse {
	s.Body = v
	return s
}

func (s *SwitchDBInstanceHAResponse) Validate() error {
	return dara.Validate(s)
}
