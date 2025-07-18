// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVectorConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVectorConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVectorConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVectorConfigurationResponseBody) *ModifyVectorConfigurationResponse
	GetBody() *ModifyVectorConfigurationResponseBody
}

type ModifyVectorConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVectorConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVectorConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVectorConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyVectorConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVectorConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVectorConfigurationResponse) GetBody() *ModifyVectorConfigurationResponseBody {
	return s.Body
}

func (s *ModifyVectorConfigurationResponse) SetHeaders(v map[string]*string) *ModifyVectorConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyVectorConfigurationResponse) SetStatusCode(v int32) *ModifyVectorConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVectorConfigurationResponse) SetBody(v *ModifyVectorConfigurationResponseBody) *ModifyVectorConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyVectorConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
