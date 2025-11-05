// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDedicatedBlockStorageClusterDiskThroughputStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse
	GetBody() *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody
}

type QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) GetBody() *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody {
	return s.Body
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) SetHeaders(v map[string]*string) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) SetStatusCode(v int32) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) SetBody(v *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponseBody) *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse {
	s.Body = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterDiskThroughputStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
