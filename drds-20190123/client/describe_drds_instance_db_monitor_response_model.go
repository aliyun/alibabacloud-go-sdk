// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceDbMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsInstanceDbMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsInstanceDbMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsInstanceDbMonitorResponseBody) *DescribeDrdsInstanceDbMonitorResponse
	GetBody() *DescribeDrdsInstanceDbMonitorResponseBody
}

type DescribeDrdsInstanceDbMonitorResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceDbMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsInstanceDbMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsInstanceDbMonitorResponse) GetBody() *DescribeDrdsInstanceDbMonitorResponseBody {
	return s.Body
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceDbMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetStatusCode(v int32) *DescribeDrdsInstanceDbMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponse) SetBody(v *DescribeDrdsInstanceDbMonitorResponseBody) *DescribeDrdsInstanceDbMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponse) Validate() error {
	return dara.Validate(s)
}
