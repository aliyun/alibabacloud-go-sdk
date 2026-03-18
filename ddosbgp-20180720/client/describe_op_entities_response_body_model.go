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
	// The details of the operation log.
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 52C8ECB0-0B1A-4E66-A31C-B6A855120E82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of operation logs.
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
	// The operation object, which is the ID of the instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The type of the operation object. The value is fixed as **1**, which indicates Anti-DDoS Origin instances.
	//
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The time when the log was generated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1635818114000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the Alibaba Cloud account that performs the operation.
	//
	// > If the value is **system**, the operation is performed by Anti-DDoS Origin.
	//
	// example:
	//
	// 171986973287****
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// The type of operation. Valid values:
	//
	// 	- **3**: indicates an operation to add an IP address to the Anti-DDoS Origin instance for protection.
	//
	// 	- **4**: indicates an operation to remove a protected IP address from the Anti-DDoS Origin instance.
	//
	// 	- **5**: indicates an operation to downgrade the Anti-DDoS Origin instance.
	//
	// 	- **6**: indicates an operation to deactivate blackhole filtering for an IP address.
	//
	// 	- **7**: indicates an operation to reset the number of times that you can deactivate blackhole filtering.
	//
	// 	- **8**: indicates an operation to enable burstable protection.
	//
	// example:
	//
	// 8
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The details of the operation. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **entity**: the operation object. Data type: object. The fields that are included in the value of the **entity*	- parameter vary based on the value of the **OpAction*	- parameter. Valid values:
	//
	//     	- If the value of the **OpAction*	- parameter is **3**, the value of the **entity*	- parameter consists of the following field:
	//
	//         	- **ips**: the public IP addresses that are protected by the Anti-DDoS Origin instance. Data type: array
	//
	//     	- If the value of the **OpAction*	- parameter is **4**, the value of the **entity*	- parameter consists of the following field:
	//
	//         	- **ips**: the public IP addresses that are no longer protected by the Anti-DDoS Origin instance. Data type: array.
	//
	//     	- If the value of the **OpAction*	- parameter is **5**, the value of the **entity*	- parameter consists of the following fields:
	//
	//         	- **baseBandwidth**: the basic protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	//         	- **elasticBandwidth**: the burstable protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	//         	- **opSource**: the source of the operation. The value is fixed as **1**, indicating that the operation is performed by Anti-DDoS Origin. Data type: integer.
	//
	//     	- If the value of the **OpAction*	- parameter is **6**, the value of the **entity*	- parameter consists of the following field:
	//
	//         	- **ips**: the public IP addresses for which to deactivate blackhole filtering. Data type: array.
	//
	//     	- If the value of the **OpAction*	- parameter is **7**, the **entity*	- parameter is not returned.
	//
	//     	- If the value of the **OpAction*	- parameter is **8**, the value of the **entity*	- parameter consists of the following fields:
	//
	//         	- **baseBandwidth**: the basic protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	//         	- **elasticBandwidth**: the burstable protection bandwidth. Unit: Gbit/s. Data type: integer.
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
