// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineBuildConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineBuildConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineBuildConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineBuildConfigurationsResponseBody) *ListRoutineBuildConfigurationsResponse
	GetBody() *ListRoutineBuildConfigurationsResponseBody
}

type ListRoutineBuildConfigurationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineBuildConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineBuildConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineBuildConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineBuildConfigurationsResponse) GetBody() *ListRoutineBuildConfigurationsResponseBody {
	return s.Body
}

func (s *ListRoutineBuildConfigurationsResponse) SetHeaders(v map[string]*string) *ListRoutineBuildConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineBuildConfigurationsResponse) SetStatusCode(v int32) *ListRoutineBuildConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineBuildConfigurationsResponse) SetBody(v *ListRoutineBuildConfigurationsResponseBody) *ListRoutineBuildConfigurationsResponse {
	s.Body = v
	return s
}

func (s *ListRoutineBuildConfigurationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
