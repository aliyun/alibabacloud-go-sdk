// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelFeatureFGFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelFeatureFGFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelFeatureFGFeatureResponse
	GetStatusCode() *int32
	SetBody(v *GetModelFeatureFGFeatureResponseBody) *GetModelFeatureFGFeatureResponse
	GetBody() *GetModelFeatureFGFeatureResponseBody
}

type GetModelFeatureFGFeatureResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelFeatureFGFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelFeatureFGFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGFeatureResponse) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelFeatureFGFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelFeatureFGFeatureResponse) GetBody() *GetModelFeatureFGFeatureResponseBody {
	return s.Body
}

func (s *GetModelFeatureFGFeatureResponse) SetHeaders(v map[string]*string) *GetModelFeatureFGFeatureResponse {
	s.Headers = v
	return s
}

func (s *GetModelFeatureFGFeatureResponse) SetStatusCode(v int32) *GetModelFeatureFGFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelFeatureFGFeatureResponse) SetBody(v *GetModelFeatureFGFeatureResponseBody) *GetModelFeatureFGFeatureResponse {
	s.Body = v
	return s
}

func (s *GetModelFeatureFGFeatureResponse) Validate() error {
	return dara.Validate(s)
}
