// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelFeatureFGFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelFeatureFGFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelFeatureFGFeatureResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelFeatureFGFeatureResponseBody) *UpdateModelFeatureFGFeatureResponse
	GetBody() *UpdateModelFeatureFGFeatureResponseBody
}

type UpdateModelFeatureFGFeatureResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelFeatureFGFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelFeatureFGFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelFeatureFGFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelFeatureFGFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelFeatureFGFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelFeatureFGFeatureResponse) GetBody() *UpdateModelFeatureFGFeatureResponseBody {
	return s.Body
}

func (s *UpdateModelFeatureFGFeatureResponse) SetHeaders(v map[string]*string) *UpdateModelFeatureFGFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelFeatureFGFeatureResponse) SetStatusCode(v int32) *UpdateModelFeatureFGFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelFeatureFGFeatureResponse) SetBody(v *UpdateModelFeatureFGFeatureResponseBody) *UpdateModelFeatureFGFeatureResponse {
	s.Body = v
	return s
}

func (s *UpdateModelFeatureFGFeatureResponse) Validate() error {
	return dara.Validate(s)
}
