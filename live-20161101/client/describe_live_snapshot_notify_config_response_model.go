// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveSnapshotNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveSnapshotNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveSnapshotNotifyConfigResponseBody) *DescribeLiveSnapshotNotifyConfigResponse
	GetBody() *DescribeLiveSnapshotNotifyConfigResponseBody
}

type DescribeLiveSnapshotNotifyConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveSnapshotNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveSnapshotNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) GetBody() *DescribeLiveSnapshotNotifyConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveSnapshotNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) SetStatusCode(v int32) *DescribeLiveSnapshotNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) SetBody(v *DescribeLiveSnapshotNotifyConfigResponseBody) *DescribeLiveSnapshotNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
