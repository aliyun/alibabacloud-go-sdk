// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryDeveloperRepoMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLibraryDeveloperRepoMetricsResponseBody
	GetRequestId() *string
	SetResult(v *LibraryDeveloperRepoMetricResult) *GetLibraryDeveloperRepoMetricsResponseBody
	GetResult() *LibraryDeveloperRepoMetricResult
}

type GetLibraryDeveloperRepoMetricsResponseBody struct {
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *LibraryDeveloperRepoMetricResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetLibraryDeveloperRepoMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryDeveloperRepoMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryDeveloperRepoMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLibraryDeveloperRepoMetricsResponseBody) GetResult() *LibraryDeveloperRepoMetricResult {
	return s.Result
}

func (s *GetLibraryDeveloperRepoMetricsResponseBody) SetRequestId(v string) *GetLibraryDeveloperRepoMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryDeveloperRepoMetricsResponseBody) SetResult(v *LibraryDeveloperRepoMetricResult) *GetLibraryDeveloperRepoMetricsResponseBody {
	s.Result = v
	return s
}

func (s *GetLibraryDeveloperRepoMetricsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
