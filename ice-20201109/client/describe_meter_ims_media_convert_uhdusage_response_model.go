// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMediaConvertUHDUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeterImsMediaConvertUHDUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeterImsMediaConvertUHDUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeterImsMediaConvertUHDUsageResponseBody) *DescribeMeterImsMediaConvertUHDUsageResponse
	GetBody() *DescribeMeterImsMediaConvertUHDUsageResponseBody
}

type DescribeMeterImsMediaConvertUHDUsageResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterImsMediaConvertUHDUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterImsMediaConvertUHDUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMediaConvertUHDUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) GetBody() *DescribeMeterImsMediaConvertUHDUsageResponseBody {
	return s.Body
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) SetHeaders(v map[string]*string) *DescribeMeterImsMediaConvertUHDUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) SetStatusCode(v int32) *DescribeMeterImsMediaConvertUHDUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) SetBody(v *DescribeMeterImsMediaConvertUHDUsageResponseBody) *DescribeMeterImsMediaConvertUHDUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeMeterImsMediaConvertUHDUsageResponse) Validate() error {
	return dara.Validate(s)
}
