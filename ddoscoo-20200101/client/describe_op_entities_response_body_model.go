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
	// The operation log records.
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// FB24D70C-71F5-4000-8CD8-22CDA0C53CD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of operation records.
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
	// - **1**: IP Address of Anti-DDoS Proxy instance
	//
	// - **2**: Anti-DDoS Pro and Anti-DDoS Premium mitigation plan
	//
	// - **3**: ECS instance
	//
	// - **4**: full logs
	//
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The time when the operation was performed. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1584451769000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The Alibaba Cloud account ID that performed the operation.
	//
	// example:
	//
	// 128965410602****
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// The type of the operation. Valid values:
	//
	// - **1**: Settings for the burstable protection bandwidth.
	//
	// - **5**: Offset with an anti-DDoS mitigation plan.
	//
	// - **8**: Change an ECS IP address.
	//
	// - **9**: Unblock (deactivation of blackhole filtering).
	//
	// - **10**: Settings for near-origin traffic diversion.
	//
	// - **11**: Clear full logs.
	//
	// - **12**: Downgrade instance specifications. This indicates that the burstable protection bandwidth is reduced when the instance expires or the account has an overdue payment.
	//
	// - **13**: Recover instance specifications. This indicates that the burstable protection bandwidth is recovered when the instance renewal is completed or the overdue payment is settled.
	//
	// example:
	//
	// 9
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The description of the operation, expressed as a JSON format character string with the following structure:
	//
	// - **newEntity**: String. The parameters after the operation.
	//
	// - **oldEntity**: String. The parameters before the operation.
	//
	// Both **newEntity*	- and **oldEntity*	- are expressed as JSON format character strings. The operation parameters vary based on the operation type (**OpAction**).
	//
	// When **OpAction*	- is **1**, **12**, or **13**, the operation parameters have the following structure:
	//
	// - **elasticBandwidth**: Integer. The burstable protection bandwidth. Unit: Gbit/s.
	//
	//     Example: `{"newEntity":{"elasticBandwidth":300},"oldEntity":{"elasticBandwidth":300}}`
	//
	//
	// When **OpAction*	- is **5**, the operation parameters have the following structure:
	//
	// - **bandwidth**: Integer. The burstable protection bandwidth. Unit: Gbit/s.
	//
	// - **count**: Integer. The number of anti-DDoS mitigation plans.
	//
	// - **deductCount**: Integer. The number of offset anti-DDoS mitigation plans.
	//
	// - **expireTime**: Long. The expiration time of the anti-DDoS mitigation plan. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// - **instanceId**: String. The instance ID of the Anti-DDoS Pro and Anti-DDoS Premium instance.
	//
	// - **peakFlow**: Integer. The peak traffic. Unit: bit/s.
	//
	//     Example: `{"newEntity":{"bandwidth":100,"count":4,"deductCount":1,"expireTime":1616299196000,"instanceId":"ddoscoo-cn-v641kpmq****","peakFlow":751427000}}`
	//
	//
	// When **OpAction*	- is **8**, the operation parameters have the following structure:
	//
	// - **instanceId**: String. The instance ID of the ECS instance whose IP address is changed.
	//
	//     Example: `{"newEntity":{"instanceId":"i-wz9h6nc313zptbqn****"}}`
	//
	//
	//
	// When **OpAction*	- is **9**, the operation parameters have the following structure:
	//
	// - **actionMethod**: String. The operation method. Value: **undo**, which indicates Unblock (deactivation of blackhole filtering).
	//
	//     Example: `{"newEntity":{"actionMethod":"undo"}}`
	//
	//
	//
	// When **OpAction*	- is **10**, the operation parameters have the following structure:
	//
	// - **actionMethod**: String. The operation method. Valid values:
	//
	//     - **do**: Enable near-origin traffic diversion.
	//
	//     - **undo**: Disable near-origin traffic diversion.
	//
	// - **lines**: Array. The line types. Valid values:
	//
	//     - **ct**: China Telecom outside China line
	//
	//     - **cut**: China Unicom outside China line
	//
	//     Example: `{"newEntity":{"actionMethod":"undo","lines":["ct"]}}`
	//
	//
	// When **OpAction*	- is **11**, there are no operation parameters and the operation description is empty.
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
