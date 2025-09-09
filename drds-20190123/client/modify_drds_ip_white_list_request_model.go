// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDrdsIpWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ModifyDrdsIpWhiteListRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *ModifyDrdsIpWhiteListRequest
	GetDrdsInstanceId() *string
	SetGroupAttribute(v string) *ModifyDrdsIpWhiteListRequest
	GetGroupAttribute() *string
	SetGroupName(v string) *ModifyDrdsIpWhiteListRequest
	GetGroupName() *string
	SetIpWhiteList(v string) *ModifyDrdsIpWhiteListRequest
	GetIpWhiteList() *string
	SetMode(v bool) *ModifyDrdsIpWhiteListRequest
	GetMode() *bool
}

type ModifyDrdsIpWhiteListRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds****c6vxxyzd
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The attribute of the IP address whitelist group.
	GroupAttribute *string `json:"GroupAttribute,omitempty" xml:"GroupAttribute,omitempty"`
	// The name of the IP address whitelist group.
	//
	// example:
	//
	// drds_******
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modified whitelist. Separate multiple IP addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.***.***.***,10.***.***.***
	IpWhiteList *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	// Specifies the mode. Valid values:
	//
	// 	- `True`: append modifications
	//
	// 	- `False`: overwrite modification
	//
	// example:
	//
	// false
	Mode *bool `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyDrdsIpWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDrdsIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyDrdsIpWhiteListRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ModifyDrdsIpWhiteListRequest) GetGroupAttribute() *string {
	return s.GroupAttribute
}

func (s *ModifyDrdsIpWhiteListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDrdsIpWhiteListRequest) GetIpWhiteList() *string {
	return s.IpWhiteList
}

func (s *ModifyDrdsIpWhiteListRequest) GetMode() *bool {
	return s.Mode
}

func (s *ModifyDrdsIpWhiteListRequest) SetDbName(v string) *ModifyDrdsIpWhiteListRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetDrdsInstanceId(v string) *ModifyDrdsIpWhiteListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupAttribute(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupAttribute = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetGroupName(v string) *ModifyDrdsIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetIpWhiteList(v string) *ModifyDrdsIpWhiteListRequest {
	s.IpWhiteList = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) SetMode(v bool) *ModifyDrdsIpWhiteListRequest {
	s.Mode = &v
	return s
}

func (s *ModifyDrdsIpWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
