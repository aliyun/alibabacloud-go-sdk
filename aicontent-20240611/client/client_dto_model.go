// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ClientDTO
	GetAddress() *string
	SetAllowedModelGroupConfig(v string) *ClientDTO
	GetAllowedModelGroupConfig() *string
	SetAllowedModels(v string) *ClientDTO
	GetAllowedModels() *string
	SetClientUuid(v string) *ClientDTO
	GetClientUuid() *string
	SetContact(v string) *ClientDTO
	GetContact() *string
	SetDeleteTag(v int32) *ClientDTO
	GetDeleteTag() *int32
	SetDiscount(v float64) *ClientDTO
	GetDiscount() *float64
	SetGmtCreate(v string) *ClientDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ClientDTO
	GetGmtModified() *string
	SetId(v int64) *ClientDTO
	GetId() *int64
	SetLevel(v int32) *ClientDTO
	GetLevel() *int32
	SetMain(v int32) *ClientDTO
	GetMain() *int32
	SetName(v string) *ClientDTO
	GetName() *string
	SetNodeType(v string) *ClientDTO
	GetNodeType() *string
	SetParentId(v int64) *ClientDTO
	GetParentId() *int64
	SetRemark(v string) *ClientDTO
	GetRemark() *string
	SetUserId(v int64) *ClientDTO
	GetUserId() *int64
}

type ClientDTO struct {
	// The company address.
	//
	// example:
	//
	// Hangzhou
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The authorization management configuration in the format of {"model_ids":[...],"group_ids":[...]}. An empty value indicates that group authorization is not used.
	//
	// example:
	//
	// {"model_ids":[1,2],"group_ids":["mg_xxx"]}
	AllowedModelGroupConfig *string `json:"allowedModelGroupConfig,omitempty" xml:"allowedModelGroupConfig,omitempty"`
	// The list of allowed model IDs, separated by commas. An empty value indicates all models.
	//
	// example:
	//
	// 1,2,3
	AllowedModels *string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty"`
	// The customer UUID.
	//
	// example:
	//
	// st-xxxx
	ClientUuid *string `json:"clientUuid,omitempty" xml:"clientUuid,omitempty"`
	// The contact information.
	//
	// example:
	//
	// 13800138000
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// The status. A value of 0 indicates enabled, and a non-zero value indicates disabled.
	//
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// The discount factor. A value of 1.0 indicates no discount.
	//
	// example:
	//
	// 1.0
	Discount *float64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// ID
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The hierarchy level. Valid values: 1 to 5.
	//
	// example:
	//
	// 1
	Level *int32 `json:"level,omitempty" xml:"level,omitempty"`
	// The account type. A value of 0 indicates a main account, and a value of 1 indicates a regular account.
	//
	// example:
	//
	// 0
	Main *int32 `json:"main,omitempty" xml:"main,omitempty"`
	// The customer name.
	//
	// example:
	//
	// My customer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The node type. Valid values: department: Department. user: User node, which is a leaf in the department tree.
	//
	// example:
	//
	// department
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// The parent department ID. An empty value indicates the top level.
	//
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Remarks
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The user ID. This field has a value only when the node type is user.
	//
	// example:
	//
	// 30001
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ClientDTO) String() string {
	return dara.Prettify(s)
}

func (s ClientDTO) GoString() string {
	return s.String()
}

func (s *ClientDTO) GetAddress() *string {
	return s.Address
}

func (s *ClientDTO) GetAllowedModelGroupConfig() *string {
	return s.AllowedModelGroupConfig
}

func (s *ClientDTO) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ClientDTO) GetClientUuid() *string {
	return s.ClientUuid
}

func (s *ClientDTO) GetContact() *string {
	return s.Contact
}

func (s *ClientDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ClientDTO) GetDiscount() *float64 {
	return s.Discount
}

func (s *ClientDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ClientDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ClientDTO) GetId() *int64 {
	return s.Id
}

func (s *ClientDTO) GetLevel() *int32 {
	return s.Level
}

func (s *ClientDTO) GetMain() *int32 {
	return s.Main
}

func (s *ClientDTO) GetName() *string {
	return s.Name
}

func (s *ClientDTO) GetNodeType() *string {
	return s.NodeType
}

func (s *ClientDTO) GetParentId() *int64 {
	return s.ParentId
}

func (s *ClientDTO) GetRemark() *string {
	return s.Remark
}

func (s *ClientDTO) GetUserId() *int64 {
	return s.UserId
}

func (s *ClientDTO) SetAddress(v string) *ClientDTO {
	s.Address = &v
	return s
}

func (s *ClientDTO) SetAllowedModelGroupConfig(v string) *ClientDTO {
	s.AllowedModelGroupConfig = &v
	return s
}

func (s *ClientDTO) SetAllowedModels(v string) *ClientDTO {
	s.AllowedModels = &v
	return s
}

func (s *ClientDTO) SetClientUuid(v string) *ClientDTO {
	s.ClientUuid = &v
	return s
}

func (s *ClientDTO) SetContact(v string) *ClientDTO {
	s.Contact = &v
	return s
}

func (s *ClientDTO) SetDeleteTag(v int32) *ClientDTO {
	s.DeleteTag = &v
	return s
}

func (s *ClientDTO) SetDiscount(v float64) *ClientDTO {
	s.Discount = &v
	return s
}

func (s *ClientDTO) SetGmtCreate(v string) *ClientDTO {
	s.GmtCreate = &v
	return s
}

func (s *ClientDTO) SetGmtModified(v string) *ClientDTO {
	s.GmtModified = &v
	return s
}

func (s *ClientDTO) SetId(v int64) *ClientDTO {
	s.Id = &v
	return s
}

func (s *ClientDTO) SetLevel(v int32) *ClientDTO {
	s.Level = &v
	return s
}

func (s *ClientDTO) SetMain(v int32) *ClientDTO {
	s.Main = &v
	return s
}

func (s *ClientDTO) SetName(v string) *ClientDTO {
	s.Name = &v
	return s
}

func (s *ClientDTO) SetNodeType(v string) *ClientDTO {
	s.NodeType = &v
	return s
}

func (s *ClientDTO) SetParentId(v int64) *ClientDTO {
	s.ParentId = &v
	return s
}

func (s *ClientDTO) SetRemark(v string) *ClientDTO {
	s.Remark = &v
	return s
}

func (s *ClientDTO) SetUserId(v int64) *ClientDTO {
	s.UserId = &v
	return s
}

func (s *ClientDTO) Validate() error {
	return dara.Validate(s)
}
