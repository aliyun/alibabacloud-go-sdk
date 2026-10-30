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
	SetTotalCount(v int32) *DescribeOpEntitiesResponseBody
	GetTotalCount() *int32
}

type DescribeOpEntitiesResponseBody struct {
	// The list of Operation logs details.
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 52C8ECB0-0B1A-4E66-A31C-B6A855120E82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried operation logs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeOpEntitiesResponseBody) GetTotalCount() *int32 {
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

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int32) *DescribeOpEntitiesResponseBody {
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
	// The operation object, which is the ID of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The object type of the operation. The value is fixed as **1**, which indicates an Anti-DDoS Origin instance.
	//
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The time when the log was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1635818114000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The Alibaba Cloud account ID used to execute the operation.
	//
	// > If the value is **system**, the Anti-DDoS Origin service will automatically execute the operation.
	//
	// example:
	//
	// 171986973287****
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// The operation type. Valid values:
	//
	// - **3**: add protected IP addresses to the Anti-DDoS Origin instance
	//
	// - **4**: unbind protected IP addresses from the Anti-DDoS Origin instance
	//
	// - **5**: downgrade the Anti-DDoS Origin instance
	//
	// - **6**: deactivate blackhole filtering
	//
	// - **7**: reset the quota on blackhole filtering deactivation
	//
	// - **8**: resume burstable protection
	//
	// example:
	//
	// 8
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The operation details. Use a JSON struct to transform the data into a character string representation. The JSON structure contains the following fields:
	//
	// - **entity**: Object | The operation object. The fields contained in **entity*	- vary based on the operation type (**OpAction**). Details:
	//
	//     - If **OpAction*	- is **3*	- (add protected IP addresses), **entity*	- contains the following fields:
	//
	//         - **ips**: Array | The list of assets that are assigned public IP addresses that are attached to the Anti-DDoS Origin instance for protection.
	//
	//     - If **OpAction*	- is **4*	- (unbind protected IP addresses), **entity*	- contains the following fields:
	//
	//         - **ips**: Array | The list of assets that are assigned public IP addresses that are unbound from the Anti-DDoS Origin instance.
	//
	//     - If **OpAction*	- is **5*	- (decrease the quota of the instance), **entity*	- contains the following fields:
	//
	//         - **baseBandwidth**: Integer | The basic protection bandwidth. Unit: Gbit/s.
	//
	//         - **elasticBandwidth**: Integer | The burstable protection bandwidth. Unit: Gbit/s.
	//
	//         - **opSource**: Integer | The source of the operation. The value is fixed as **1**, which indicates that the Anti-DDoS Origin service will automatically execute the operation.
	//
	//     - If **OpAction*	- is **6*	- (deactivate blackhole filtering), **entity*	- contains the following fields:
	//
	//         - **ips**: Array | The list of assets that are assigned public IP addresses for which the blackhole filtering status is deactivated.
	//
	//     - If **OpAction*	- is **7*	- (Reset quota on blackhole filtering deactivation), **entity*	- is empty.
	//
	//     - If **OpAction*	- is **8*	- (recover burstable protection), **entity*	- contains the following fields:
	//
	//         - **baseBandwidth**: Integer | The basic protection bandwidth. Unit: Gbit/s.
	//
	//         - **elasticBandwidth**: Integer | The burstable protection bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// {"entity":{"baseBandwidth":20,"elasticBandwidth":20}}
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
