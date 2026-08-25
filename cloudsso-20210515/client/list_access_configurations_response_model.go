// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessConfigurationsResponseBody) *ListAccessConfigurationsResponse
	GetBody() *ListAccessConfigurationsResponseBody
}

type ListAccessConfigurationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessConfigurationsResponse) GetBody() *ListAccessConfigurationsResponseBody {
	return s.Body
}

func (s *ListAccessConfigurationsResponse) SetHeaders(v map[string]*string) *ListAccessConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessConfigurationsResponse) SetStatusCode(v int32) *ListAccessConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessConfigurationsResponse) SetBody(v *ListAccessConfigurationsResponseBody) *ListAccessConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListAccessConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
