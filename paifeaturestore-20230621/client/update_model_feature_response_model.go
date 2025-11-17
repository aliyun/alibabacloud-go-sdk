// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelFeatureResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelFeatureResponseBody) *UpdateModelFeatureResponse
	GetBody() *UpdateModelFeatureResponseBody
}

type UpdateModelFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelFeatureResponse) GetBody() *UpdateModelFeatureResponseBody {
	return s.Body
}

func (s *UpdateModelFeatureResponse) SetHeaders(v map[string]*string) *UpdateModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelFeatureResponse) SetStatusCode(v int32) *UpdateModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelFeatureResponse) SetBody(v *UpdateModelFeatureResponseBody) *UpdateModelFeatureResponse {
	s.Body = v
	return s
}

func (s *UpdateModelFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
