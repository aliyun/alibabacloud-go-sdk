// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChartReleases(v []*ListChartReleaseResponseBodyChartReleases) *ListChartReleaseResponseBody
	GetChartReleases() []*ListChartReleaseResponseBodyChartReleases
	SetCode(v string) *ListChartReleaseResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListChartReleaseResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListChartReleaseResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChartReleaseResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChartReleaseResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListChartReleaseResponseBody
	GetTotalCount() *string
}

type ListChartReleaseResponseBody struct {
	// The list of chart versions.
	ChartReleases []*ListChartReleaseResponseBodyChartReleases `json:"ChartReleases,omitempty" xml:"ChartReleases,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F68823F6-F1B5-4A4E-8421-A83CAB8F2963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChartReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChartReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartReleaseResponseBody) GetChartReleases() []*ListChartReleaseResponseBodyChartReleases {
	return s.ChartReleases
}

func (s *ListChartReleaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChartReleaseResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListChartReleaseResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChartReleaseResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChartReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChartReleaseResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListChartReleaseResponseBody) SetChartReleases(v []*ListChartReleaseResponseBodyChartReleases) *ListChartReleaseResponseBody {
	s.ChartReleases = v
	return s
}

func (s *ListChartReleaseResponseBody) SetCode(v string) *ListChartReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetIsSuccess(v bool) *ListChartReleaseResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetPageNo(v int32) *ListChartReleaseResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetPageSize(v int32) *ListChartReleaseResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetRequestId(v string) *ListChartReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetTotalCount(v string) *ListChartReleaseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChartReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListChartReleaseResponseBodyChartReleases struct {
	// The name of the chart version.
	//
	// example:
	//
	// chart1
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the chart was last modified.
	//
	// example:
	//
	// 1571930323000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The version number of the chart.
	//
	// example:
	//
	// 0.1.0
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The ID of the chart repository.
	//
	// example:
	//
	// crcr-gpsu7b8chmxk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The size of the chart of the version. Unit: bytes.
	//
	// example:
	//
	// 2826
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the chart.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListChartReleaseResponseBodyChartReleases) String() string {
	return dara.Prettify(s)
}

func (s ListChartReleaseResponseBodyChartReleases) GoString() string {
	return s.String()
}

func (s *ListChartReleaseResponseBodyChartReleases) GetChart() *string {
	return s.Chart
}

func (s *ListChartReleaseResponseBodyChartReleases) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChartReleaseResponseBodyChartReleases) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListChartReleaseResponseBodyChartReleases) GetRelease() *string {
	return s.Release
}

func (s *ListChartReleaseResponseBodyChartReleases) GetRepoId() *string {
	return s.RepoId
}

func (s *ListChartReleaseResponseBodyChartReleases) GetSize() *string {
	return s.Size
}

func (s *ListChartReleaseResponseBodyChartReleases) GetStatus() *string {
	return s.Status
}

func (s *ListChartReleaseResponseBodyChartReleases) SetChart(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Chart = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetInstanceId(v string) *ListChartReleaseResponseBodyChartReleases {
	s.InstanceId = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetModifiedTime(v int64) *ListChartReleaseResponseBodyChartReleases {
	s.ModifiedTime = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetRelease(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Release = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetRepoId(v string) *ListChartReleaseResponseBodyChartReleases {
	s.RepoId = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetSize(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Size = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetStatus(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Status = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) Validate() error {
	return dara.Validate(s)
}
