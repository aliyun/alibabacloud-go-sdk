// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStateConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStateConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStateConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListStateConfigurationsResponseBody) *ListStateConfigurationsResponse
	GetBody() *ListStateConfigurationsResponseBody
}

type ListStateConfigurationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStateConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStateConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStateConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStateConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStateConfigurationsResponse) GetBody() *ListStateConfigurationsResponseBody {
	return s.Body
}

func (s *ListStateConfigurationsResponse) SetHeaders(v map[string]*string) *ListStateConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListStateConfigurationsResponse) SetStatusCode(v int32) *ListStateConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStateConfigurationsResponse) SetBody(v *ListStateConfigurationsResponseBody) *ListStateConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListStateConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
