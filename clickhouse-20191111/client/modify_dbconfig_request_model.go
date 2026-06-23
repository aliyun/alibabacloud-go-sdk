// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyDBConfigRequest
	GetConfig() *string
	SetDBClusterId(v string) *ModifyDBConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyDBConfigRequest struct {
	// The dictionary configuration items.
	//
	// > The value of the Config parameter overwrites the existing configuration. To delete a dictionary configuration, remove it from the request.
	//
	// Call the [DescribeDBConfig](https://help.aliyun.com/document_detail/2360627.html) operation to query the dictionary configuration of the target instance.
	//
	// example:
	//
	// <dictionaries><dictionary><name>test_dictionary</name><source><clickhouse><host>10.0.0.0</host><port>3003</port><user>TestUser</user><password>testPassword</password><db>default</db><table>dictTestTable01</table><where>active=1</where><invalidate_query>SELECT max(value) FROM dictTestTable01</invalidate_query></clickhouse></source><lifetime><min>200</min><max>600</max></lifetime><layout><flat></flat></layout><structure><id><name>id</name><type>UInt64</type></id><attribute><name>value</name><type>String</type><null_value></null_value></attribute></structure></dictionary></dictionaries>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1r59y779o04****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. Call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the IDs of available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyDBConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBConfigRequest) SetConfig(v string) *ModifyDBConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyDBConfigRequest) SetDBClusterId(v string) *ModifyDBConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetOwnerAccount(v string) *ModifyDBConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBConfigRequest) SetOwnerId(v int64) *ModifyDBConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetRegionId(v string) *ModifyDBConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBConfigRequest) SetResourceOwnerId(v int64) *ModifyDBConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBConfigRequest) Validate() error {
	return dara.Validate(s)
}
