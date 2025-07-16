// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewOnlineFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureViewOnlineFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureViewOnlineFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureViewOnlineFeaturesResponseBody) *ListFeatureViewOnlineFeaturesResponse
	GetBody() *ListFeatureViewOnlineFeaturesResponseBody
}

type ListFeatureViewOnlineFeaturesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewOnlineFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewOnlineFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureViewOnlineFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureViewOnlineFeaturesResponse) GetBody() *ListFeatureViewOnlineFeaturesResponseBody {
	return s.Body
}

func (s *ListFeatureViewOnlineFeaturesResponse) SetHeaders(v map[string]*string) *ListFeatureViewOnlineFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponse) SetStatusCode(v int32) *ListFeatureViewOnlineFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponse) SetBody(v *ListFeatureViewOnlineFeaturesResponseBody) *ListFeatureViewOnlineFeaturesResponse {
	s.Body = v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponse) Validate() error {
	return dara.Validate(s)
}
