// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchLSQLV3MySQLServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchLSQLV3MySQLServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchLSQLV3MySQLServiceResponse
	GetStatusCode() *int32
	SetBody(v *SwitchLSQLV3MySQLServiceResponseBody) *SwitchLSQLV3MySQLServiceResponse
	GetBody() *SwitchLSQLV3MySQLServiceResponseBody
}

type SwitchLSQLV3MySQLServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchLSQLV3MySQLServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchLSQLV3MySQLServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchLSQLV3MySQLServiceResponse) GoString() string {
	return s.String()
}

func (s *SwitchLSQLV3MySQLServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchLSQLV3MySQLServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchLSQLV3MySQLServiceResponse) GetBody() *SwitchLSQLV3MySQLServiceResponseBody {
	return s.Body
}

func (s *SwitchLSQLV3MySQLServiceResponse) SetHeaders(v map[string]*string) *SwitchLSQLV3MySQLServiceResponse {
	s.Headers = v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponse) SetStatusCode(v int32) *SwitchLSQLV3MySQLServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponse) SetBody(v *SwitchLSQLV3MySQLServiceResponseBody) *SwitchLSQLV3MySQLServiceResponse {
	s.Body = v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
