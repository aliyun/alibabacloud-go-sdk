// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListEffectiveAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*DescribeWhiteListEffectiveAssetsResponseBodyAssets) *DescribeWhiteListEffectiveAssetsResponseBody
	GetAssets() []*DescribeWhiteListEffectiveAssetsResponseBodyAssets
	SetCount(v int32) *DescribeWhiteListEffectiveAssetsResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeWhiteListEffectiveAssetsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWhiteListEffectiveAssetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWhiteListEffectiveAssetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWhiteListEffectiveAssetsResponseBody
	GetTotalCount() *int32
}

type DescribeWhiteListEffectiveAssetsResponseBody struct {
	// The servers on which the policy takes effect.
	Assets []*DescribeWhiteListEffectiveAssetsResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB39****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the servers on which the policy takes effect.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhiteListEffectiveAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListEffectiveAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) GetAssets() []*DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	return s.Assets
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) SetAssets(v []*DescribeWhiteListEffectiveAssetsResponseBodyAssets) *DescribeWhiteListEffectiveAssetsResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) SetCount(v int32) *DescribeWhiteListEffectiveAssetsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) SetCurrentPage(v int32) *DescribeWhiteListEffectiveAssetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) SetPageSize(v int32) *DescribeWhiteListEffectiveAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) SetRequestId(v string) *DescribeWhiteListEffectiveAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) SetTotalCount(v int32) *DescribeWhiteListEffectiveAssetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteListEffectiveAssetsResponseBodyAssets struct {
	// The public IP address of the server.
	//
	// example:
	//
	// 60.205.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// test
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The exception handling mode. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: alerted
	//
	// 	- **2**: blocked
	//
	// example:
	//
	// 1
	ProcessMethod *int32 `json:"ProcessMethod,omitempty" xml:"ProcessMethod,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 35815387
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The number of **untrusted programs*	- on the server.
	//
	// example:
	//
	// 5
	SuspiciousEventCount *int32 `json:"SuspiciousEventCount,omitempty" xml:"SuspiciousEventCount,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 544900ff-1be7-4655-9719-6311cecb3****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWhiteListEffectiveAssetsResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListEffectiveAssetsResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetMachineName() *string {
	return s.MachineName
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetProcessMethod() *int32 {
	return s.ProcessMethod
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetSuspiciousEventCount() *int32 {
	return s.SuspiciousEventCount
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetInternetIp(v string) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.InternetIp = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetIntranetIp(v string) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetMachineName(v string) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.MachineName = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetProcessMethod(v int32) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.ProcessMethod = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetStrategyId(v int64) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetStrategyName(v string) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.StrategyName = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetSuspiciousEventCount(v int32) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.SuspiciousEventCount = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) SetUuid(v string) *DescribeWhiteListEffectiveAssetsResponseBodyAssets {
	s.Uuid = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}
