// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsInstanceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsInstanceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsInstanceMonitorResponseBody) *DescribeDrdsInstanceMonitorResponse
	GetBody() *DescribeDrdsInstanceMonitorResponseBody
}

type DescribeDrdsInstanceMonitorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsInstanceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsInstanceMonitorResponse) GetBody() *DescribeDrdsInstanceMonitorResponseBody {
	return s.Body
}

func (s *DescribeDrdsInstanceMonitorResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) SetStatusCode(v int32) *DescribeDrdsInstanceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) SetBody(v *DescribeDrdsInstanceMonitorResponseBody) *DescribeDrdsInstanceMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponse) Validate() error {
	return dara.Validate(s)
}
