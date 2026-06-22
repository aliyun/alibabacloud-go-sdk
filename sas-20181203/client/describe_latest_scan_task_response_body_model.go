// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLatestScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLastCheckTime(v int64) *DescribeLatestScanTaskResponseBody
	GetLastCheckTime() *int64
	SetRequestId(v string) *DescribeLatestScanTaskResponseBody
	GetRequestId() *string
	SetRiskNum(v int32) *DescribeLatestScanTaskResponseBody
	GetRiskNum() *int32
	SetTargetInfo(v string) *DescribeLatestScanTaskResponseBody
	GetTargetInfo() *string
	SetUuids(v []*string) *DescribeLatestScanTaskResponseBody
	GetUuids() []*string
}

type DescribeLatestScanTaskResponseBody struct {
	// The timestamp of the most recent scan, in milliseconds.
	//
	// example:
	//
	// 1671610264000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of virus risks detected on the server.
	//
	// example:
	//
	// 1
	RiskNum *int32 `json:"RiskNum,omitempty" xml:"RiskNum,omitempty"`
	// The asset information scanned by the virus scan node. This parameter is expressed as a character string converted from a JSON array. The following fields are included:
	//
	// - **type**: The Asset Type on which the virus scan is executed. Valid values:
	//
	//     - **groupId**: server group.
	//
	//     - **uuid**: server.
	//
	// - **name**: The name of the server group or server.
	//
	// - **target**: The asset on which the virus scan is executed. The following describes the values of this field:
	//
	//     - If **type*	- is set to **groupId**, this field specifies the server group ID.
	//
	//     - If **type*	- is set to **uuid**, this field specifies the UUID of the server.
	//
	// example:
	//
	// [{"type":"uuid","name":"Host001","target":"503201a7-14c6-4280-801b-1169ed42****"}]
	TargetInfo *string `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty"`
	// The list of UUIDs of the assets.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s DescribeLatestScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLatestScanTaskResponseBody) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *DescribeLatestScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLatestScanTaskResponseBody) GetRiskNum() *int32 {
	return s.RiskNum
}

func (s *DescribeLatestScanTaskResponseBody) GetTargetInfo() *string {
	return s.TargetInfo
}

func (s *DescribeLatestScanTaskResponseBody) GetUuids() []*string {
	return s.Uuids
}

func (s *DescribeLatestScanTaskResponseBody) SetLastCheckTime(v int64) *DescribeLatestScanTaskResponseBody {
	s.LastCheckTime = &v
	return s
}

func (s *DescribeLatestScanTaskResponseBody) SetRequestId(v string) *DescribeLatestScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLatestScanTaskResponseBody) SetRiskNum(v int32) *DescribeLatestScanTaskResponseBody {
	s.RiskNum = &v
	return s
}

func (s *DescribeLatestScanTaskResponseBody) SetTargetInfo(v string) *DescribeLatestScanTaskResponseBody {
	s.TargetInfo = &v
	return s
}

func (s *DescribeLatestScanTaskResponseBody) SetUuids(v []*string) *DescribeLatestScanTaskResponseBody {
	s.Uuids = v
	return s
}

func (s *DescribeLatestScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
