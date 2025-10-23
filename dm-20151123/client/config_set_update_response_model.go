// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetUpdateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetUpdateResponseBody) *ConfigSetUpdateResponse
	GetBody() *ConfigSetUpdateResponseBody
}

type ConfigSetUpdateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetUpdateResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetUpdateResponse) GetBody() *ConfigSetUpdateResponseBody {
	return s.Body
}

func (s *ConfigSetUpdateResponse) SetHeaders(v map[string]*string) *ConfigSetUpdateResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetUpdateResponse) SetStatusCode(v int32) *ConfigSetUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetUpdateResponse) SetBody(v *ConfigSetUpdateResponseBody) *ConfigSetUpdateResponse {
	s.Body = v
	return s
}

func (s *ConfigSetUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
