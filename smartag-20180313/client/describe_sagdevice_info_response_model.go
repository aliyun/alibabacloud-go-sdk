// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSAGDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSAGDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSAGDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSAGDeviceInfoResponseBody) *DescribeSAGDeviceInfoResponse
	GetBody() *DescribeSAGDeviceInfoResponseBody
}

type DescribeSAGDeviceInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSAGDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSAGDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSAGDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSAGDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSAGDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSAGDeviceInfoResponse) GetBody() *DescribeSAGDeviceInfoResponseBody {
	return s.Body
}

func (s *DescribeSAGDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeSAGDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSAGDeviceInfoResponse) SetStatusCode(v int32) *DescribeSAGDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponse) SetBody(v *DescribeSAGDeviceInfoResponseBody) *DescribeSAGDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSAGDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
