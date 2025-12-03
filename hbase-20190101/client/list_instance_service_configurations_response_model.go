// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceServiceConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceServiceConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceServiceConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceServiceConfigurationsResponseBody) *ListInstanceServiceConfigurationsResponse
	GetBody() *ListInstanceServiceConfigurationsResponseBody
}

type ListInstanceServiceConfigurationsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceServiceConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceServiceConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceServiceConfigurationsResponse) GetBody() *ListInstanceServiceConfigurationsResponseBody {
	return s.Body
}

func (s *ListInstanceServiceConfigurationsResponse) SetHeaders(v map[string]*string) *ListInstanceServiceConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponse) SetStatusCode(v int32) *ListInstanceServiceConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponse) SetBody(v *ListInstanceServiceConfigurationsResponseBody) *ListInstanceServiceConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
