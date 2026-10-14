// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSystemLogResponseBody
	GetRequestId() *string
	SetSystemLog(v []*DescribeSystemLogResponseBodySystemLog) *DescribeSystemLogResponseBody
	GetSystemLog() []*DescribeSystemLogResponseBodySystemLog
	SetTotal(v int64) *DescribeSystemLogResponseBody
	GetTotal() *int64
}

type DescribeSystemLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8BC3A33A-F832-58DB-952F-7682A25AD14C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the billing logs for burstable clean bandwidth.
	SystemLog []*DescribeSystemLogResponseBodySystemLog `json:"SystemLog,omitempty" xml:"SystemLog,omitempty" type:"Repeated"`
	// The total number of billing logs for burstable clean bandwidth that are queried.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeSystemLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemLogResponseBody) GetSystemLog() []*DescribeSystemLogResponseBodySystemLog {
	return s.SystemLog
}

func (s *DescribeSystemLogResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSystemLogResponseBody) SetRequestId(v string) *DescribeSystemLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemLogResponseBody) SetSystemLog(v []*DescribeSystemLogResponseBodySystemLog) *DescribeSystemLogResponseBody {
	s.SystemLog = v
	return s
}

func (s *DescribeSystemLogResponseBody) SetTotal(v int64) *DescribeSystemLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSystemLogResponseBody) Validate() error {
	if s.SystemLog != nil {
		for _, item := range s.SystemLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemLogResponseBodySystemLog struct {
	// The IP address of the Anti-DDoS instance.
	//
	// example:
	//
	// 203.107.XX.XX
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The type of the system log. The value is fixed at **20**, which indicates the billing log of burstable clean bandwidth.
	//
	// example:
	//
	// 20
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The time when the bill was created. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1631793531000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the bill was last modified. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1635425407000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Alibaba Cloud account to which the bill belongs.
	//
	// example:
	//
	// 171986973287****
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// The operation type. The value is fixed at **100**, which indicates that a billing record for burstable clean bandwidth is added.
	//
	// example:
	//
	// 100
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The details of the bill. This value is a string that is converted from a JSON struct. The JSON struct contains the following fields:
	//
	// - **newEntity**: Object type, which indicates the bill record and contains the following fields:
	//
	//   - **billValue**: Integer type | The usage of burstable clean bandwidth in the billing month. Unit: Mbit/s.
	//
	//   - **instanceId**: String type | The ID of the Anti-DDoS instance.
	//
	//   - **ip**: String type | The IP address of the Anti-DDoS instance.
	//
	//   - **maxBw**: String type | The actual peak value of service traffic in the billing month (monthly 95th percentile). Unit: Mbit/s.
	//
	//   - **month**: Long type | The month when the bill is generated (the billing month is the calendar month before the month when the bill is generated). This value is a timestamp in milliseconds.
	//
	//   - **overBandwidth**: Integer type | The bandwidth that exceeds the clean bandwidth specification of the instance. Unit: Mbit/s.
	//
	//   - **peakTime**: Long type | The time when the actual peak value of service traffic is measured. This value is a timestamp in seconds.
	//
	//   - **startTimestamp**: Long type | The start time of the period when the peak value of instance service traffic is measured. This value is a timestamp in seconds.
	//
	// example:
	//
	// {"newEntity":{"billValue":"60","instanceId":"ddoscoo-cn-zz121ogz****","ip":"203.107.XX.XX","maxBw":"300","month":1627747200000,"overBandwidth":"120","peakTime":1629871200,"startTimestamp":1629871200}}
	OpDesc *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
	// The status of the bill. Valid values:
	//
	// - **0**: indicates that the bill is pending.
	//
	// - **1**: indicates that the bill is generated.
	//
	// - **2**: indicates that the bill is terminated.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSystemLogResponseBodySystemLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLogResponseBodySystemLog) GoString() string {
	return s.String()
}

func (s *DescribeSystemLogResponseBodySystemLog) GetEntityObject() *string {
	return s.EntityObject
}

func (s *DescribeSystemLogResponseBodySystemLog) GetEntityType() *int32 {
	return s.EntityType
}

func (s *DescribeSystemLogResponseBodySystemLog) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSystemLogResponseBodySystemLog) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSystemLogResponseBodySystemLog) GetOpAccount() *string {
	return s.OpAccount
}

func (s *DescribeSystemLogResponseBodySystemLog) GetOpAction() *int32 {
	return s.OpAction
}

func (s *DescribeSystemLogResponseBodySystemLog) GetOpDesc() *string {
	return s.OpDesc
}

func (s *DescribeSystemLogResponseBodySystemLog) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeSystemLogResponseBodySystemLog) SetEntityObject(v string) *DescribeSystemLogResponseBodySystemLog {
	s.EntityObject = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetEntityType(v int32) *DescribeSystemLogResponseBodySystemLog {
	s.EntityType = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetGmtCreate(v int64) *DescribeSystemLogResponseBodySystemLog {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetGmtModified(v int64) *DescribeSystemLogResponseBodySystemLog {
	s.GmtModified = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetOpAccount(v string) *DescribeSystemLogResponseBodySystemLog {
	s.OpAccount = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetOpAction(v int32) *DescribeSystemLogResponseBodySystemLog {
	s.OpAction = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetOpDesc(v string) *DescribeSystemLogResponseBodySystemLog {
	s.OpDesc = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) SetStatus(v int32) *DescribeSystemLogResponseBodySystemLog {
	s.Status = &v
	return s
}

func (s *DescribeSystemLogResponseBodySystemLog) Validate() error {
	return dara.Validate(s)
}
