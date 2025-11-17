// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectFeaturesResponseBody) *ListProjectFeaturesResponse
	GetBody() *ListProjectFeaturesResponseBody
}

type ListProjectFeaturesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectFeaturesResponse) GetBody() *ListProjectFeaturesResponseBody {
	return s.Body
}

func (s *ListProjectFeaturesResponse) SetHeaders(v map[string]*string) *ListProjectFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectFeaturesResponse) SetStatusCode(v int32) *ListProjectFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectFeaturesResponse) SetBody(v *ListProjectFeaturesResponseBody) *ListProjectFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListProjectFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
