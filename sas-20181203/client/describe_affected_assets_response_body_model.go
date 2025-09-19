// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAffectedAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*DescribeAffectedAssetsResponseBodyAssetList) *DescribeAffectedAssetsResponseBody
	GetAssetList() []*DescribeAffectedAssetsResponseBodyAssetList
	SetPageInfo(v *DescribeAffectedAssetsResponseBodyPageInfo) *DescribeAffectedAssetsResponseBody
	GetPageInfo() *DescribeAffectedAssetsResponseBodyPageInfo
	SetRequestId(v string) *DescribeAffectedAssetsResponseBody
	GetRequestId() *string
}

type DescribeAffectedAssetsResponseBody struct {
	// An array that consists of the affected servers.
	AssetList []*DescribeAffectedAssetsResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeAffectedAssetsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAffectedAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAffectedAssetsResponseBody) GetAssetList() []*DescribeAffectedAssetsResponseBodyAssetList {
	return s.AssetList
}

func (s *DescribeAffectedAssetsResponseBody) GetPageInfo() *DescribeAffectedAssetsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeAffectedAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAffectedAssetsResponseBody) SetAssetList(v []*DescribeAffectedAssetsResponseBodyAssetList) *DescribeAffectedAssetsResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeAffectedAssetsResponseBody) SetPageInfo(v *DescribeAffectedAssetsResponseBodyPageInfo) *DescribeAffectedAssetsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAffectedAssetsResponseBody) SetRequestId(v string) *DescribeAffectedAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAffectedAssetsResponseBodyAssetList struct {
	// The ID of the server.
	//
	// example:
	//
	// 11
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// 11
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 10.10.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.0.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The number of viruses detected on the server.
	//
	// example:
	//
	// 1
	RiskNum *int32 `json:"RiskNum,omitempty" xml:"RiskNum,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 947d7514-258a-4b47-9dde-9dxxxxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAffectedAssetsResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedAssetsResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) GetRiskNum() *int32 {
	return s.RiskNum
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) SetInstanceId(v string) *DescribeAffectedAssetsResponseBodyAssetList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) SetInstanceName(v string) *DescribeAffectedAssetsResponseBodyAssetList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) SetInternetIp(v string) *DescribeAffectedAssetsResponseBodyAssetList {
	s.InternetIp = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) SetIntranetIp(v string) *DescribeAffectedAssetsResponseBodyAssetList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) SetRiskNum(v int32) *DescribeAffectedAssetsResponseBodyAssetList {
	s.RiskNum = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) SetUuid(v string) *DescribeAffectedAssetsResponseBodyAssetList {
	s.Uuid = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyAssetList) Validate() error {
	return dara.Validate(s)
}

type DescribeAffectedAssetsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAffectedAssetsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedAssetsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) SetCount(v int32) *DescribeAffectedAssetsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAffectedAssetsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) SetPageSize(v int32) *DescribeAffectedAssetsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAffectedAssetsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAffectedAssetsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
