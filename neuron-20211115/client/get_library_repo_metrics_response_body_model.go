// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryRepoMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLibraryRepoMetricsResponseBody
	GetRequestId() *string
	SetResult(v *LibraryRepoMetricResult) *GetLibraryRepoMetricsResponseBody
	GetResult() *LibraryRepoMetricResult
}

type GetLibraryRepoMetricsResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LibraryRepoMetricResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetLibraryRepoMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryRepoMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryRepoMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLibraryRepoMetricsResponseBody) GetResult() *LibraryRepoMetricResult {
	return s.Result
}

func (s *GetLibraryRepoMetricsResponseBody) SetRequestId(v string) *GetLibraryRepoMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryRepoMetricsResponseBody) SetResult(v *LibraryRepoMetricResult) *GetLibraryRepoMetricsResponseBody {
	s.Result = v
	return s
}

func (s *GetLibraryRepoMetricsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
