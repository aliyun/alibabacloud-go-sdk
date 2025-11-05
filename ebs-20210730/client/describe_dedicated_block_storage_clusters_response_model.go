// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedBlockStorageClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedBlockStorageClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedBlockStorageClustersResponseBody) *DescribeDedicatedBlockStorageClustersResponse
	GetBody() *DescribeDedicatedBlockStorageClustersResponseBody
}

type DescribeDedicatedBlockStorageClustersResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedBlockStorageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedBlockStorageClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedBlockStorageClustersResponse) GetBody() *DescribeDedicatedBlockStorageClustersResponseBody {
	return s.Body
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetStatusCode(v int32) *DescribeDedicatedBlockStorageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetBody(v *DescribeDedicatedBlockStorageClustersResponseBody) *DescribeDedicatedBlockStorageClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
