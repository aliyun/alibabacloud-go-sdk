// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskMonitorDataListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskMonitorDataListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskMonitorDataListResponseBody) *DescribeDiskMonitorDataListResponse
	GetBody() *DescribeDiskMonitorDataListResponseBody
}

type DescribeDiskMonitorDataListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskMonitorDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskMonitorDataListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskMonitorDataListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskMonitorDataListResponse) GetBody() *DescribeDiskMonitorDataListResponseBody {
	return s.Body
}

func (s *DescribeDiskMonitorDataListResponse) SetHeaders(v map[string]*string) *DescribeDiskMonitorDataListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskMonitorDataListResponse) SetStatusCode(v int32) *DescribeDiskMonitorDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskMonitorDataListResponse) SetBody(v *DescribeDiskMonitorDataListResponseBody) *DescribeDiskMonitorDataListResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskMonitorDataListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
