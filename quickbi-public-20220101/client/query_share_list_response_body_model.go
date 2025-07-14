// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryShareListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryShareListResponseBody
	GetRequestId() *string
	SetResult(v []*QueryShareListResponseBodyResult) *QueryShareListResponseBody
	GetResult() []*QueryShareListResponseBodyResult
	SetSuccess(v bool) *QueryShareListResponseBody
	GetSuccess() *bool
}

type QueryShareListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the list of objects to which the work has been shared.
	Result []*QueryShareListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryShareListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryShareListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryShareListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryShareListResponseBody) GetResult() []*QueryShareListResponseBodyResult {
	return s.Result
}

func (s *QueryShareListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryShareListResponseBody) SetRequestId(v string) *QueryShareListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryShareListResponseBody) SetResult(v []*QueryShareListResponseBodyResult) *QueryShareListResponseBody {
	s.Result = v
	return s
}

func (s *QueryShareListResponseBody) SetSuccess(v bool) *QueryShareListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryShareListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryShareListResponseBodyResult struct {
	// Sharing permissions. Possible values:
	//
	// - 1: View only
	//
	// - 3: View and export
	//
	// example:
	//
	// 3
	AuthPoint *int32 `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	// The timestamp in milliseconds indicating the expiration time of the authorization.
	//
	// example:
	//
	// 1640102400000
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the sharing configuration.
	//
	// example:
	//
	// 0ab9659e-29cf-47d7-a364-3a91553b****
	ShareId *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	// The ID of the sharing target, which could be a user ID or a group ID in Quick BI.
	//
	// - When ShareToType=2 (all members within an organization), ShareToId is the organization ID.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	ShareToId *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	// The name of the sharing target.
	//
	// example:
	//
	// test
	ShareToName *string `json:"ShareToName,omitempty" xml:"ShareToName,omitempty"`
	// The type of sharing. Possible values:
	//
	// - 0: User
	//
	// - 1: Group
	//
	// - 2: Organization
	//
	// - 3: Space
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	// The type of the shared work. The value range includes:
	//
	// - dataProduct: Data Portal
	//
	// - dashboard: Dashboard
	//
	// - report: Spreadsheet
	//
	// - dashboardOfflineQuery: Self-service Data Extraction
	//
	// - ANALYSIS: Ad-hoc Analysis
	//
	// - DATAFORM: Data Entry
	//
	// - SCREEN: Data Visualization Screen
	//
	// example:
	//
	// dashboard
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s QueryShareListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryShareListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryShareListResponseBodyResult) GetAuthPoint() *int32 {
	return s.AuthPoint
}

func (s *QueryShareListResponseBodyResult) GetExpireDate() *int64 {
	return s.ExpireDate
}

func (s *QueryShareListResponseBodyResult) GetReportId() *string {
	return s.ReportId
}

func (s *QueryShareListResponseBodyResult) GetShareId() *string {
	return s.ShareId
}

func (s *QueryShareListResponseBodyResult) GetShareToId() *string {
	return s.ShareToId
}

func (s *QueryShareListResponseBodyResult) GetShareToName() *string {
	return s.ShareToName
}

func (s *QueryShareListResponseBodyResult) GetShareToType() *int32 {
	return s.ShareToType
}

func (s *QueryShareListResponseBodyResult) GetShareType() *string {
	return s.ShareType
}

func (s *QueryShareListResponseBodyResult) SetAuthPoint(v int32) *QueryShareListResponseBodyResult {
	s.AuthPoint = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetExpireDate(v int64) *QueryShareListResponseBodyResult {
	s.ExpireDate = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetReportId(v string) *QueryShareListResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareId(v string) *QueryShareListResponseBodyResult {
	s.ShareId = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareToId(v string) *QueryShareListResponseBodyResult {
	s.ShareToId = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareToName(v string) *QueryShareListResponseBodyResult {
	s.ShareToName = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareToType(v int32) *QueryShareListResponseBodyResult {
	s.ShareToType = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareType(v string) *QueryShareListResponseBodyResult {
	s.ShareType = &v
	return s
}

func (s *QueryShareListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
