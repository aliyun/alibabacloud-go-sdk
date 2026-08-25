// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessConfigurationProvisioningsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessConfigurationProvisioningsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessConfigurationProvisioningsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessConfigurationProvisioningsResponseBody) *ListAccessConfigurationProvisioningsResponse
	GetBody() *ListAccessConfigurationProvisioningsResponseBody
}

type ListAccessConfigurationProvisioningsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessConfigurationProvisioningsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessConfigurationProvisioningsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessConfigurationProvisioningsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessConfigurationProvisioningsResponse) GetBody() *ListAccessConfigurationProvisioningsResponseBody {
	return s.Body
}

func (s *ListAccessConfigurationProvisioningsResponse) SetHeaders(v map[string]*string) *ListAccessConfigurationProvisioningsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponse) SetStatusCode(v int32) *ListAccessConfigurationProvisioningsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponse) SetBody(v *ListAccessConfigurationProvisioningsResponseBody) *ListAccessConfigurationProvisioningsResponse {
	s.Body = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
