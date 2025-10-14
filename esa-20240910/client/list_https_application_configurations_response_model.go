// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsApplicationConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpsApplicationConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpsApplicationConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpsApplicationConfigurationsResponseBody) *ListHttpsApplicationConfigurationsResponse
	GetBody() *ListHttpsApplicationConfigurationsResponseBody
}

type ListHttpsApplicationConfigurationsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpsApplicationConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpsApplicationConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsApplicationConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListHttpsApplicationConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpsApplicationConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpsApplicationConfigurationsResponse) GetBody() *ListHttpsApplicationConfigurationsResponseBody {
	return s.Body
}

func (s *ListHttpsApplicationConfigurationsResponse) SetHeaders(v map[string]*string) *ListHttpsApplicationConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponse) SetStatusCode(v int32) *ListHttpsApplicationConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponse) SetBody(v *ListHttpsApplicationConfigurationsResponseBody) *ListHttpsApplicationConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListHttpsApplicationConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
