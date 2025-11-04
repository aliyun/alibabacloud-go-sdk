// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMediaConvertUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeterImsMediaConvertUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeterImsMediaConvertUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeterImsMediaConvertUsageResponseBody) *DescribeMeterImsMediaConvertUsageResponse
	GetBody() *DescribeMeterImsMediaConvertUsageResponseBody
}

type DescribeMeterImsMediaConvertUsageResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterImsMediaConvertUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterImsMediaConvertUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeterImsMediaConvertUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeterImsMediaConvertUsageResponse) GetBody() *DescribeMeterImsMediaConvertUsageResponseBody {
	return s.Body
}

func (s *DescribeMeterImsMediaConvertUsageResponse) SetHeaders(v map[string]*string) *DescribeMeterImsMediaConvertUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponse) SetStatusCode(v int32) *DescribeMeterImsMediaConvertUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponse) SetBody(v *DescribeMeterImsMediaConvertUsageResponseBody) *DescribeMeterImsMediaConvertUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeMeterImsMediaConvertUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
