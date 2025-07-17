// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentFeaturesResponseBody) *ListEnvironmentFeaturesResponse
	GetBody() *ListEnvironmentFeaturesResponseBody
}

type ListEnvironmentFeaturesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentFeaturesResponse) GetBody() *ListEnvironmentFeaturesResponseBody {
	return s.Body
}

func (s *ListEnvironmentFeaturesResponse) SetHeaders(v map[string]*string) *ListEnvironmentFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentFeaturesResponse) SetStatusCode(v int32) *ListEnvironmentFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentFeaturesResponse) SetBody(v *ListEnvironmentFeaturesResponseBody) *ListEnvironmentFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentFeaturesResponse) Validate() error {
	return dara.Validate(s)
}
