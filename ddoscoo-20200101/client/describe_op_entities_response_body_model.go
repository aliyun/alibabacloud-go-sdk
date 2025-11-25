// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody
	GetOpEntities() []*DescribeOpEntitiesResponseBodyOpEntities
	SetRequestId(v string) *DescribeOpEntitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeOpEntitiesResponseBody
	GetTotalCount() *int64
}

type DescribeOpEntitiesResponseBody struct {
	// The operation records.
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FB24D70C-71F5-4000-8CD8-22CDA0C53CD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned operation records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) GetOpEntities() []*DescribeOpEntitiesResponseBodyOpEntities {
	return s.OpEntities
}

func (s *DescribeOpEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpEntitiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int64) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) Validate() error {
	if s.OpEntities != nil {
		for _, item := range s.OpEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	// The operation object.
	//
	// example:
	//
	// 203.***.***.132
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The type of the operation object. Valid values:
	//
	// 	- **1**: the IP address of the Anti-DDoS Proxy instance.
	//
	// 	- **2**: Anti-DDoS plans.
	//
	// 	- **3**: ECS instances.
	//
	// 	- **4**: all logs.
	//
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The time when the operation was performed. The value is a UNIX timestamp. Units: milliseconds.
	//
	// example:
	//
	// 1584451769000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the Alibaba Cloud account that is used to perform the operation.
	//
	// example:
	//
	// 128965410602****
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- **1**: configuring burstable protection bandwidth.
	//
	// 	- **5**: using Anti-DDoS plans.
	//
	// 	- **8**: changing the IP addresses of ECS instances.
	//
	// 	- **9**: deactivating blackhole filtering.
	//
	// 	- **10**: configuring the near-origin traffic diversion feature.
	//
	// 	- **11**: clearing all logs.
	//
	// 	- **12**: downgrading the specifications of the Anti-DDoS Proxy instance. If the instance expires or the account has overdue payments, this operation is performed to downgrade the burstable protection bandwidth.
	//
	// 	- **13**: restoring the specifications of the Anti-DDoS Proxy instance. If the instance is renewed or you have paid the overdue payments within your account, this operation is performed to restore the burstable protection bandwidth.
	//
	// example:
	//
	// 9
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The details of the operation. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **newEntity**: the values of the parameters after the operation. This field is of the string type.
	//
	// 	- **oldEntity**: the values of the parameters before the operation. This field is of the string type.
	//
	// Both **newEntity*	- and **oldEntity*	- are JSON strings. The returned parameters vary based on **OpAction**.
	//
	// If the value of **OpAction*	- is **1**, **12**, or **13**, the following parameter is returned:
	//
	// 	- **elasticBandwidth**: the burstable protection bandwidth. This parameter is of the integer type. Unit: Gbit/s.
	//
	//     Example: `{"newEntity":{"elasticBandwidth":300},"oldEntity":{"elasticBandwidth":300}}`
	//
	// If the value of **OpAction*	- is **5**, the following parameters are returned:
	//
	// 	- **bandwidth**: the burstable protection bandwidth. The parameter is of the integer type. Unit: Gbit/s.
	//
	// 	- **count**: the number of Anti-DDoS plans. This parameter is of the integer type.
	//
	// 	- **deductCount**: the number of used Anti-DDoS plans. This parameter is of the integer type.
	//
	// 	- **expireTime**: the expiration time of the Anti-DDoS plans. This parameter is of the long type. The value is a UNIX timestamp. Units: milliseconds.
	//
	// 	- **instanceId**: the ID of the Anti-DDoS Proxy instance. This parameter is of the string type.
	//
	// 	- **peakFlow**: the peak throughput on the Anti-DDoS Proxy instance. This parameter is of the integer type. Unit: bit/s.
	//
	//     Example: `{"newEntity":{"bandwidth":100,"count":4,"deductCount":1,"expireTime":1616299196000,"instanceId":"ddoscoo-cn-v641kpmq****","peakFlow":751427000}}`
	//
	// If the value of **OpAction*	- is **8**, the following parameter is returned:
	//
	// 	- **instanceId**: the IDs of the ECS instances whose IP addresses are changed. This parameter is of the string type.
	//
	//     Example: `{"newEntity":{"instanceId":"i-wz9h6nc313zptbqn****"}}`
	//
	// If the value of **OpAction*	- is **9**, the following parameter is returned:
	//
	// 	- **actionMethod**: the operation method. This parameter is of the string type. Valid value: **undo**, which indicates that you deactivated blackhole filtering.
	//
	//     Example: `{"newEntity":{"actionMethod":"undo"}}`
	//
	// If the value of **OpAction*	- is **10**, the following parameters are returned:
	//
	// 	- **actionMethod**: the operation method. This parameter is of the string type. Valid values:
	//
	//     	- **do**: The near-origin traffic diversion feature is enabled.
	//
	//     	- **undo**: The near-origin traffic diversion feature is disabled.
	//
	// 	- **lines**: the Internet service provider (ISP) line from which the traffic is blocked. This parameter is of the array type. Valid values:
	//
	//     	- **ct**: China Telecom (International).
	//
	//     	- **cut**: China Unicom (International).
	//
	//     Example: `{"newEntity":{"actionMethod":"undo","lines":["ct"]}}`
	//
	// If the value of **OpAction*	- is **11**, no parameter is returned, and the description is empty.
	//
	// example:
	//
	// {"newEntity":{"actionMethod":"undo"}}
	OpDesc *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetEntityObject() *string {
	return s.EntityObject
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetEntityType() *int32 {
	return s.EntityType
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetOpAccount() *string {
	return s.OpAccount
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetOpAction() *int32 {
	return s.OpAction
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetOpDesc() *string {
	return s.OpDesc
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) Validate() error {
	return dara.Validate(s)
}
