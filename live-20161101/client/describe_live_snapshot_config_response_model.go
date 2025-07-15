// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveSnapshotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveSnapshotConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveSnapshotConfigResponseBody) *DescribeLiveSnapshotConfigResponse
	GetBody() *DescribeLiveSnapshotConfigResponseBody
}

type DescribeLiveSnapshotConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveSnapshotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveSnapshotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveSnapshotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveSnapshotConfigResponse) GetBody() *DescribeLiveSnapshotConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveSnapshotConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveSnapshotConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetStatusCode(v int32) *DescribeLiveSnapshotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) SetBody(v *DescribeLiveSnapshotConfigResponseBody) *DescribeLiveSnapshotConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveSnapshotConfigResponse) Validate() error {
	return dara.Validate(s)
}
