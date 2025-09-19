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
	// The timestamp when the last check was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1671610264000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of virus detection risks on the server.
	//
	// example:
	//
	// 1
	RiskNum *int32 `json:"RiskNum,omitempty" xml:"RiskNum,omitempty"`
	// The applicable scope of the whitelist. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: the type of the applicable scope. Valid values:
	//
	//     	- **GroupId**: the ID of a server group
	//
	//     	- **Uuid**: the UUID of a server
	//
	// 	- **uuids**: the UUIDs of servers
	//
	// 	- **groupIds**: the IDs of server groups
	//
	// >  If you leave this parameter empty, all servers are added to the whitelist. If you set the **type*	- field to **GroupId**, you must also specify the **groupIds*	- field. If you set the **type*	- field to **Uuid**, you must also specify the **uuids*	- field.
	//
	// example:
	//
	// [{"type":"uuid","name":"Host001","target":"503201a7-14c6-4280-801b-1169ed42****"}]
	TargetInfo *string `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty"`
	// The UUIDs of the assets.
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
