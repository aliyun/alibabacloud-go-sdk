// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEgressProxyConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEgressProxyConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListEgressProxyConfigurationsResponseBody) *ListEgressProxyConfigurationsResponse
	GetBody() *ListEgressProxyConfigurationsResponseBody
}

type ListEgressProxyConfigurationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEgressProxyConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEgressProxyConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListEgressProxyConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEgressProxyConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEgressProxyConfigurationsResponse) GetBody() *ListEgressProxyConfigurationsResponseBody {
	return s.Body
}

func (s *ListEgressProxyConfigurationsResponse) SetHeaders(v map[string]*string) *ListEgressProxyConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListEgressProxyConfigurationsResponse) SetStatusCode(v int32) *ListEgressProxyConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponse) SetBody(v *ListEgressProxyConfigurationsResponseBody) *ListEgressProxyConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListEgressProxyConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
