// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInlinePolicyForAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInlinePolicyForAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInlinePolicyForAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInlinePolicyForAccessConfigurationResponseBody) *UpdateInlinePolicyForAccessConfigurationResponse
	GetBody() *UpdateInlinePolicyForAccessConfigurationResponseBody
}

type UpdateInlinePolicyForAccessConfigurationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInlinePolicyForAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInlinePolicyForAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInlinePolicyForAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) GetBody() *UpdateInlinePolicyForAccessConfigurationResponseBody {
	return s.Body
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetHeaders(v map[string]*string) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetStatusCode(v int32) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) SetBody(v *UpdateInlinePolicyForAccessConfigurationResponseBody) *UpdateInlinePolicyForAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
