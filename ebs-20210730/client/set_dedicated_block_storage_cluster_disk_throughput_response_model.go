// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDedicatedBlockStorageClusterDiskThroughputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDedicatedBlockStorageClusterDiskThroughputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDedicatedBlockStorageClusterDiskThroughputResponse
	GetStatusCode() *int32
	SetBody(v *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) *SetDedicatedBlockStorageClusterDiskThroughputResponse
	GetBody() *SetDedicatedBlockStorageClusterDiskThroughputResponseBody
}

type SetDedicatedBlockStorageClusterDiskThroughputResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDedicatedBlockStorageClusterDiskThroughputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDedicatedBlockStorageClusterDiskThroughputResponse) GoString() string {
	return s.String()
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) GetBody() *SetDedicatedBlockStorageClusterDiskThroughputResponseBody {
	return s.Body
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) SetHeaders(v map[string]*string) *SetDedicatedBlockStorageClusterDiskThroughputResponse {
	s.Headers = v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) SetStatusCode(v int32) *SetDedicatedBlockStorageClusterDiskThroughputResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) SetBody(v *SetDedicatedBlockStorageClusterDiskThroughputResponseBody) *SetDedicatedBlockStorageClusterDiskThroughputResponse {
	s.Body = v
	return s
}

func (s *SetDedicatedBlockStorageClusterDiskThroughputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
