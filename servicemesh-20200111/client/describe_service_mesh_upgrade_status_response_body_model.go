// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshUpgradeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceMeshUpgradeStatusResponseBody
	GetRequestId() *string
	SetUpgradeDetail(v *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) *DescribeServiceMeshUpgradeStatusResponseBody
	GetUpgradeDetail() *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail
}

type DescribeServiceMeshUpgradeStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upgrade results.
	UpgradeDetail *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail `json:"UpgradeDetail,omitempty" xml:"UpgradeDetail,omitempty" type:"Struct"`
}

func (s DescribeServiceMeshUpgradeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) GetUpgradeDetail() *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	return s.UpgradeDetail
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) SetRequestId(v string) *DescribeServiceMeshUpgradeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) SetUpgradeDetail(v *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) *DescribeServiceMeshUpgradeStatusResponseBody {
	s.UpgradeDetail = v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBody) Validate() error {
	if s.UpgradeDetail != nil {
		if err := s.UpgradeDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail struct {
	// The number of ingress gateways that are upgraded.
	//
	// example:
	//
	// 1
	FinishedGatewaysNum *int64 `json:"FinishedGatewaysNum,omitempty" xml:"FinishedGatewaysNum,omitempty"`
	// The information about the status of the ingress gateways.
	GatewayStatusRecord map[string]*UpgradeDetailGatewayStatusRecordValue `json:"GatewayStatusRecord,omitempty" xml:"GatewayStatusRecord,omitempty"`
	// The status of the ASM instance. Valid values:
	//
	// 	- running: The instance is running.
	//
	// 	- `upgrading`: The instance is being upgraded.
	//
	// 	- `upgrading_failed`: The upgrade of the instance fails.
	//
	// example:
	//
	// running
	MeshStatus *string `json:"MeshStatus,omitempty" xml:"MeshStatus,omitempty"`
	// The total number of ingress gateways in the ASM instance.
	//
	// example:
	//
	// 2
	TotalGatewaysNum *int64 `json:"TotalGatewaysNum,omitempty" xml:"TotalGatewaysNum,omitempty"`
}

func (s DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) GetFinishedGatewaysNum() *int64 {
	return s.FinishedGatewaysNum
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) GetGatewayStatusRecord() map[string]*UpgradeDetailGatewayStatusRecordValue {
	return s.GatewayStatusRecord
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) GetMeshStatus() *string {
	return s.MeshStatus
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) GetTotalGatewaysNum() *int64 {
	return s.TotalGatewaysNum
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetFinishedGatewaysNum(v int64) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.FinishedGatewaysNum = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetGatewayStatusRecord(v map[string]*UpgradeDetailGatewayStatusRecordValue) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.GatewayStatusRecord = v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetMeshStatus(v string) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.MeshStatus = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) SetTotalGatewaysNum(v int64) *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail {
	s.TotalGatewaysNum = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponseBodyUpgradeDetail) Validate() error {
	return dara.Validate(s)
}
