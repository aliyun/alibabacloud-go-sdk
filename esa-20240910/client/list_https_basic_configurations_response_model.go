// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsBasicConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpsBasicConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpsBasicConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpsBasicConfigurationsResponseBody) *ListHttpsBasicConfigurationsResponse
	GetBody() *ListHttpsBasicConfigurationsResponseBody
}

type ListHttpsBasicConfigurationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpsBasicConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpsBasicConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpsBasicConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpsBasicConfigurationsResponse) GetBody() *ListHttpsBasicConfigurationsResponseBody {
	return s.Body
}

func (s *ListHttpsBasicConfigurationsResponse) SetHeaders(v map[string]*string) *ListHttpsBasicConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListHttpsBasicConfigurationsResponse) SetStatusCode(v int32) *ListHttpsBasicConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponse) SetBody(v *ListHttpsBasicConfigurationsResponseBody) *ListHttpsBasicConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListHttpsBasicConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
