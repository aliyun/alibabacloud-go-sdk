// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhysicalConnectionFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPhysicalConnectionFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPhysicalConnectionFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListPhysicalConnectionFeaturesResponseBody) *ListPhysicalConnectionFeaturesResponse
	GetBody() *ListPhysicalConnectionFeaturesResponseBody
}

type ListPhysicalConnectionFeaturesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPhysicalConnectionFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPhysicalConnectionFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPhysicalConnectionFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListPhysicalConnectionFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPhysicalConnectionFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPhysicalConnectionFeaturesResponse) GetBody() *ListPhysicalConnectionFeaturesResponseBody {
	return s.Body
}

func (s *ListPhysicalConnectionFeaturesResponse) SetHeaders(v map[string]*string) *ListPhysicalConnectionFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponse) SetStatusCode(v int32) *ListPhysicalConnectionFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponse) SetBody(v *ListPhysicalConnectionFeaturesResponseBody) *ListPhysicalConnectionFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
