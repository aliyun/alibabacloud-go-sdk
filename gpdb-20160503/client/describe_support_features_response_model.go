// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportFeaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportFeaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportFeaturesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportFeaturesResponseBody) *DescribeSupportFeaturesResponse
	GetBody() *DescribeSupportFeaturesResponseBody
}

type DescribeSupportFeaturesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportFeaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportFeaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportFeaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportFeaturesResponse) GetBody() *DescribeSupportFeaturesResponseBody {
	return s.Body
}

func (s *DescribeSupportFeaturesResponse) SetHeaders(v map[string]*string) *DescribeSupportFeaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportFeaturesResponse) SetStatusCode(v int32) *DescribeSupportFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportFeaturesResponse) SetBody(v *DescribeSupportFeaturesResponseBody) *DescribeSupportFeaturesResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportFeaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
