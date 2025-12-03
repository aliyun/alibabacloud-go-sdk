// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToMountPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachVscToMountPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachVscToMountPointsResponse
	GetStatusCode() *int32
	SetBody(v *AttachVscToMountPointsResponseBody) *AttachVscToMountPointsResponse
	GetBody() *AttachVscToMountPointsResponseBody
}

type AttachVscToMountPointsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachVscToMountPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachVscToMountPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToMountPointsResponse) GoString() string {
	return s.String()
}

func (s *AttachVscToMountPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachVscToMountPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachVscToMountPointsResponse) GetBody() *AttachVscToMountPointsResponseBody {
	return s.Body
}

func (s *AttachVscToMountPointsResponse) SetHeaders(v map[string]*string) *AttachVscToMountPointsResponse {
	s.Headers = v
	return s
}

func (s *AttachVscToMountPointsResponse) SetStatusCode(v int32) *AttachVscToMountPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVscToMountPointsResponse) SetBody(v *AttachVscToMountPointsResponseBody) *AttachVscToMountPointsResponse {
	s.Body = v
	return s
}

func (s *AttachVscToMountPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
