// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerScanConfigResponseBody) *DescribeContainerScanConfigResponse
	GetBody() *DescribeContainerScanConfigResponseBody
}

type DescribeContainerScanConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerScanConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerScanConfigResponse) GetBody() *DescribeContainerScanConfigResponseBody {
	return s.Body
}

func (s *DescribeContainerScanConfigResponse) SetHeaders(v map[string]*string) *DescribeContainerScanConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerScanConfigResponse) SetStatusCode(v int32) *DescribeContainerScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerScanConfigResponse) SetBody(v *DescribeContainerScanConfigResponseBody) *DescribeContainerScanConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerScanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
