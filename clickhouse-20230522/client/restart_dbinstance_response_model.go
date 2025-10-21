// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse
	GetBody() *RestartDBInstanceResponseBody
}

type RestartDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDBInstanceResponse) GetBody() *RestartDBInstanceResponseBody {
	return s.Body
}

func (s *RestartDBInstanceResponse) SetHeaders(v map[string]*string) *RestartDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDBInstanceResponse) SetStatusCode(v int32) *RestartDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBInstanceResponse) SetBody(v *RestartDBInstanceResponseBody) *RestartDBInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
