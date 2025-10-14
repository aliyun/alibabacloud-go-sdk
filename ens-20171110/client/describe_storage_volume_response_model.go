// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageVolumeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageVolumeResponseBody) *DescribeStorageVolumeResponse
	GetBody() *DescribeStorageVolumeResponseBody
}

type DescribeStorageVolumeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageVolumeResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageVolumeResponse) GetBody() *DescribeStorageVolumeResponseBody {
	return s.Body
}

func (s *DescribeStorageVolumeResponse) SetHeaders(v map[string]*string) *DescribeStorageVolumeResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageVolumeResponse) SetStatusCode(v int32) *DescribeStorageVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageVolumeResponse) SetBody(v *DescribeStorageVolumeResponseBody) *DescribeStorageVolumeResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageVolumeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
