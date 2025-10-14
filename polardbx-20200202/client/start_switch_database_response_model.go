// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSwitchDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSwitchDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSwitchDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *StartSwitchDatabaseResponseBody) *StartSwitchDatabaseResponse
	GetBody() *StartSwitchDatabaseResponseBody
}

type StartSwitchDatabaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSwitchDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSwitchDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSwitchDatabaseResponse) GoString() string {
	return s.String()
}

func (s *StartSwitchDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSwitchDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSwitchDatabaseResponse) GetBody() *StartSwitchDatabaseResponseBody {
	return s.Body
}

func (s *StartSwitchDatabaseResponse) SetHeaders(v map[string]*string) *StartSwitchDatabaseResponse {
	s.Headers = v
	return s
}

func (s *StartSwitchDatabaseResponse) SetStatusCode(v int32) *StartSwitchDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSwitchDatabaseResponse) SetBody(v *StartSwitchDatabaseResponseBody) *StartSwitchDatabaseResponse {
	s.Body = v
	return s
}

func (s *StartSwitchDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
