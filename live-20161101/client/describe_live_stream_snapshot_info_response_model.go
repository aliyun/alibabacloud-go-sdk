// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamSnapshotInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamSnapshotInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamSnapshotInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamSnapshotInfoResponseBody) *DescribeLiveStreamSnapshotInfoResponse
	GetBody() *DescribeLiveStreamSnapshotInfoResponseBody
}

type DescribeLiveStreamSnapshotInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamSnapshotInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamSnapshotInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamSnapshotInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamSnapshotInfoResponse) GetBody() *DescribeLiveStreamSnapshotInfoResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamSnapshotInfoResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamSnapshotInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponse) SetStatusCode(v int32) *DescribeLiveStreamSnapshotInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponse) SetBody(v *DescribeLiveStreamSnapshotInfoResponseBody) *DescribeLiveStreamSnapshotInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
