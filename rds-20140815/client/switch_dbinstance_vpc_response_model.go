// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchDBInstanceVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchDBInstanceVpcResponse
	GetStatusCode() *int32
	SetBody(v *SwitchDBInstanceVpcResponseBody) *SwitchDBInstanceVpcResponse
	GetBody() *SwitchDBInstanceVpcResponseBody
}

type SwitchDBInstanceVpcResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchDBInstanceVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchDBInstanceVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceVpcResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchDBInstanceVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchDBInstanceVpcResponse) GetBody() *SwitchDBInstanceVpcResponseBody {
	return s.Body
}

func (s *SwitchDBInstanceVpcResponse) SetHeaders(v map[string]*string) *SwitchDBInstanceVpcResponse {
	s.Headers = v
	return s
}

func (s *SwitchDBInstanceVpcResponse) SetStatusCode(v int32) *SwitchDBInstanceVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchDBInstanceVpcResponse) SetBody(v *SwitchDBInstanceVpcResponseBody) *SwitchDBInstanceVpcResponse {
	s.Body = v
	return s
}

func (s *SwitchDBInstanceVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
