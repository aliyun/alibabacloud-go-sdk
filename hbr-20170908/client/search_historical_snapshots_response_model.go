// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchHistoricalSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchHistoricalSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchHistoricalSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *SearchHistoricalSnapshotsResponseBody) *SearchHistoricalSnapshotsResponse
	GetBody() *SearchHistoricalSnapshotsResponseBody
}

type SearchHistoricalSnapshotsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchHistoricalSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchHistoricalSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchHistoricalSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchHistoricalSnapshotsResponse) GetBody() *SearchHistoricalSnapshotsResponseBody {
	return s.Body
}

func (s *SearchHistoricalSnapshotsResponse) SetHeaders(v map[string]*string) *SearchHistoricalSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *SearchHistoricalSnapshotsResponse) SetStatusCode(v int32) *SearchHistoricalSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponse) SetBody(v *SearchHistoricalSnapshotsResponseBody) *SearchHistoricalSnapshotsResponse {
	s.Body = v
	return s
}

func (s *SearchHistoricalSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
