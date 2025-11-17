// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmbeddedInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryEmbeddedInfoResponseBody
	GetRequestId() *string
	SetResult(v *QueryEmbeddedInfoResponseBodyResult) *QueryEmbeddedInfoResponseBody
	GetResult() *QueryEmbeddedInfoResponseBodyResult
	SetSuccess(v bool) *QueryEmbeddedInfoResponseBody
	GetSuccess() *bool
}

type QueryEmbeddedInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The embedded information of the reports under the organization.
	Result *QueryEmbeddedInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEmbeddedInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEmbeddedInfoResponseBody) GetResult() *QueryEmbeddedInfoResponseBodyResult {
	return s.Result
}

func (s *QueryEmbeddedInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEmbeddedInfoResponseBody) SetRequestId(v string) *QueryEmbeddedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBody) SetResult(v *QueryEmbeddedInfoResponseBodyResult) *QueryEmbeddedInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryEmbeddedInfoResponseBody) SetSuccess(v bool) *QueryEmbeddedInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEmbeddedInfoResponseBodyResult struct {
	// Embed the statistics of the work.
	Detail *QueryEmbeddedInfoResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The number of reports that are currently embedded.
	//
	// example:
	//
	// 3
	EmbeddedCount *int32 `json:"EmbeddedCount,omitempty" xml:"EmbeddedCount,omitempty"`
	// The maximum number of reports that can be embedded.
	//
	// example:
	//
	// 100
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s QueryEmbeddedInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponseBodyResult) GetDetail() *QueryEmbeddedInfoResponseBodyResultDetail {
	return s.Detail
}

func (s *QueryEmbeddedInfoResponseBodyResult) GetEmbeddedCount() *int32 {
	return s.EmbeddedCount
}

func (s *QueryEmbeddedInfoResponseBodyResult) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *QueryEmbeddedInfoResponseBodyResult) SetDetail(v *QueryEmbeddedInfoResponseBodyResultDetail) *QueryEmbeddedInfoResponseBodyResult {
	s.Detail = v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResult) SetEmbeddedCount(v int32) *QueryEmbeddedInfoResponseBodyResult {
	s.EmbeddedCount = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResult) SetMaxCount(v int32) *QueryEmbeddedInfoResponseBodyResult {
	s.MaxCount = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResult) Validate() error {
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEmbeddedInfoResponseBodyResultDetail struct {
	// The number of embedded self-service fetching.
	//
	// example:
	//
	// 1
	DashboardOfflineQuery *int32 `json:"DashboardOfflineQuery,omitempty" xml:"DashboardOfflineQuery,omitempty"`
	// The number of embedded dashboards.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of embedded spreadsheets.
	//
	// example:
	//
	// 1
	Report *int32 `json:"Report,omitempty" xml:"Report,omitempty"`
}

func (s QueryEmbeddedInfoResponseBodyResultDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedInfoResponseBodyResultDetail) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) GetDashboardOfflineQuery() *int32 {
	return s.DashboardOfflineQuery
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) GetPage() *int32 {
	return s.Page
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) GetReport() *int32 {
	return s.Report
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) SetDashboardOfflineQuery(v int32) *QueryEmbeddedInfoResponseBodyResultDetail {
	s.DashboardOfflineQuery = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) SetPage(v int32) *QueryEmbeddedInfoResponseBodyResultDetail {
	s.Page = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) SetReport(v int32) *QueryEmbeddedInfoResponseBodyResultDetail {
	s.Report = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) Validate() error {
	return dara.Validate(s)
}
