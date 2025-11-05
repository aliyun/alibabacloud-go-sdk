// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedBlockStorageClusterDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClusterDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedBlockStorageClusterDisksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedBlockStorageClusterDisksResponseBody) *DescribeDedicatedBlockStorageClusterDisksResponse
	GetBody() *DescribeDedicatedBlockStorageClusterDisksResponseBody
}

type DescribeDedicatedBlockStorageClusterDisksResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedBlockStorageClusterDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) GetBody() *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	return s.Body
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetStatusCode(v int32) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetBody(v *DescribeDedicatedBlockStorageClusterDisksResponseBody) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
