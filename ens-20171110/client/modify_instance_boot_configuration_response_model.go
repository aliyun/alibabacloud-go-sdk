// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceBootConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceBootConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceBootConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceBootConfigurationResponseBody) *ModifyInstanceBootConfigurationResponse
	GetBody() *ModifyInstanceBootConfigurationResponseBody
}

type ModifyInstanceBootConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceBootConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceBootConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceBootConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceBootConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceBootConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceBootConfigurationResponse) GetBody() *ModifyInstanceBootConfigurationResponseBody {
	return s.Body
}

func (s *ModifyInstanceBootConfigurationResponse) SetHeaders(v map[string]*string) *ModifyInstanceBootConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceBootConfigurationResponse) SetStatusCode(v int32) *ModifyInstanceBootConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceBootConfigurationResponse) SetBody(v *ModifyInstanceBootConfigurationResponseBody) *ModifyInstanceBootConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceBootConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
