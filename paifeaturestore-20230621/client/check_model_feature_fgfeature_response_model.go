// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckModelFeatureFGFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckModelFeatureFGFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckModelFeatureFGFeatureResponse
	GetStatusCode() *int32
	SetBody(v *CheckModelFeatureFGFeatureResponseBody) *CheckModelFeatureFGFeatureResponse
	GetBody() *CheckModelFeatureFGFeatureResponseBody
}

type CheckModelFeatureFGFeatureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckModelFeatureFGFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckModelFeatureFGFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckModelFeatureFGFeatureResponse) GoString() string {
	return s.String()
}

func (s *CheckModelFeatureFGFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckModelFeatureFGFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckModelFeatureFGFeatureResponse) GetBody() *CheckModelFeatureFGFeatureResponseBody {
	return s.Body
}

func (s *CheckModelFeatureFGFeatureResponse) SetHeaders(v map[string]*string) *CheckModelFeatureFGFeatureResponse {
	s.Headers = v
	return s
}

func (s *CheckModelFeatureFGFeatureResponse) SetStatusCode(v int32) *CheckModelFeatureFGFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckModelFeatureFGFeatureResponse) SetBody(v *CheckModelFeatureFGFeatureResponseBody) *CheckModelFeatureFGFeatureResponse {
	s.Body = v
	return s
}

func (s *CheckModelFeatureFGFeatureResponse) Validate() error {
	return dara.Validate(s)
}
