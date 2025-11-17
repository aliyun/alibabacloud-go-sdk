// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeatureAvailableFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelFeatureAvailableFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelFeatureAvailableFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListModelFeatureAvailableFeaturesResponseBody) *ListModelFeatureAvailableFeaturesResponse
	GetBody() *ListModelFeatureAvailableFeaturesResponseBody
}

type ListModelFeatureAvailableFeaturesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelFeatureAvailableFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelFeatureAvailableFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeatureAvailableFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListModelFeatureAvailableFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelFeatureAvailableFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelFeatureAvailableFeaturesResponse) GetBody() *ListModelFeatureAvailableFeaturesResponseBody {
	return s.Body
}

func (s *ListModelFeatureAvailableFeaturesResponse) SetHeaders(v map[string]*string) *ListModelFeatureAvailableFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponse) SetStatusCode(v int32) *ListModelFeatureAvailableFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponse) SetBody(v *ListModelFeatureAvailableFeaturesResponseBody) *ListModelFeatureAvailableFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListModelFeatureAvailableFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
