// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotDetectPornConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveSnapshotDetectPornConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveSnapshotDetectPornConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveSnapshotDetectPornConfigResponseBody) *DescribeLiveSnapshotDetectPornConfigResponse
	GetBody() *DescribeLiveSnapshotDetectPornConfigResponseBody
}

type DescribeLiveSnapshotDetectPornConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveSnapshotDetectPornConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) GetBody() *DescribeLiveSnapshotDetectPornConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetStatusCode(v int32) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) SetBody(v *DescribeLiveSnapshotDetectPornConfigResponseBody) *DescribeLiveSnapshotDetectPornConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponse) Validate() error {
	return dara.Validate(s)
}
