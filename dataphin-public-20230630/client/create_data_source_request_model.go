// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDataSourceRequestCreateCommand) *CreateDataSourceRequest
	GetCreateCommand() *CreateDataSourceRequestCreateCommand
	SetOpTenantId(v int64) *CreateDataSourceRequest
	GetOpTenantId() *int64
}

type CreateDataSourceRequest struct {
	CreateCommand *CreateDataSourceRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) GetCreateCommand() *CreateDataSourceRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDataSourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataSourceRequest) SetCreateCommand(v *CreateDataSourceRequestCreateCommand) *CreateDataSourceRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDataSourceRequest) SetOpTenantId(v int64) *CreateDataSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataSourceRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataSourceRequestCreateCommand struct {
	DevDataSourceCreate *CreateDataSourceRequestCreateCommandDevDataSourceCreate `json:"DevDataSourceCreate,omitempty" xml:"DevDataSourceCreate,omitempty" type:"Struct"`
	// 数据源创建结构体
	ProdDataSourceCreate *CreateDataSourceRequestCreateCommandProdDataSourceCreate `json:"ProdDataSourceCreate,omitempty" xml:"ProdDataSourceCreate,omitempty" type:"Struct"`
}

func (s CreateDataSourceRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommand) GetDevDataSourceCreate() *CreateDataSourceRequestCreateCommandDevDataSourceCreate {
	return s.DevDataSourceCreate
}

func (s *CreateDataSourceRequestCreateCommand) GetProdDataSourceCreate() *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	return s.ProdDataSourceCreate
}

func (s *CreateDataSourceRequestCreateCommand) SetDevDataSourceCreate(v *CreateDataSourceRequestCreateCommandDevDataSourceCreate) *CreateDataSourceRequestCreateCommand {
	s.DevDataSourceCreate = v
	return s
}

func (s *CreateDataSourceRequestCreateCommand) SetProdDataSourceCreate(v *CreateDataSourceRequestCreateCommandProdDataSourceCreate) *CreateDataSourceRequestCreateCommand {
	s.ProdDataSourceCreate = v
	return s
}

func (s *CreateDataSourceRequestCreateCommand) Validate() error {
	if s.DevDataSourceCreate != nil {
		if err := s.DevDataSourceCreate.Validate(); err != nil {
			return err
		}
	}
	if s.ProdDataSourceCreate != nil {
		if err := s.ProdDataSourceCreate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataSourceRequestCreateCommandDevDataSourceCreate struct {
	// 数据源创建结构体
	DataSourceCreate *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate `json:"DataSourceCreate,omitempty" xml:"DataSourceCreate,omitempty" type:"Struct"`
	// example:
	//
	// 1011
	ProdDataSourceId *int64 `json:"ProdDataSourceId,omitempty" xml:"ProdDataSourceId,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreate) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreate) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) GetDataSourceCreate() *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	return s.DataSourceCreate
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) GetProdDataSourceId() *int64 {
	return s.ProdDataSourceId
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) SetDataSourceCreate(v *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) *CreateDataSourceRequestCreateCommandDevDataSourceCreate {
	s.DataSourceCreate = v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) SetProdDataSourceId(v int64) *CreateDataSourceRequestCreateCommandDevDataSourceCreate {
	s.ProdDataSourceId = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) Validate() error {
	if s.DataSourceCreate != nil {
		if err := s.DataSourceCreate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate struct {
	// example:
	//
	// true
	CheckActivity *bool `json:"CheckActivity,omitempty" xml:"CheckActivity,omitempty"`
	// This parameter is required.
	ConfigItemList []*CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// datasource for xxx in dev
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dp_test_dev
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GetCheckActivity() *bool {
	return s.CheckActivity
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GetConfigItemList() []*CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList {
	return s.ConfigItemList
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GetDescription() *string {
	return s.Description
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GetName() *string {
	return s.Name
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GetType() *string {
	return s.Type
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetCheckActivity(v bool) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.CheckActivity = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetConfigItemList(v []*CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.ConfigItemList = v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetDescription(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetName(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetType(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) Validate() error {
	if s.ConfigItemList != nil {
		for _, item := range s.ConfigItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) GetKey() *string {
	return s.Key
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) GetValue() *string {
	return s.Value
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) SetKey(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList {
	s.Key = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) SetValue(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList {
	s.Value = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) Validate() error {
	return dara.Validate(s)
}

type CreateDataSourceRequestCreateCommandProdDataSourceCreate struct {
	// example:
	//
	// true
	CheckActivity *bool `json:"CheckActivity,omitempty" xml:"CheckActivity,omitempty"`
	// This parameter is required.
	ConfigItemList []*CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// datasource for xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dp_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreate) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreate) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) GetCheckActivity() *bool {
	return s.CheckActivity
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) GetConfigItemList() []*CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList {
	return s.ConfigItemList
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) GetDescription() *string {
	return s.Description
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) GetName() *string {
	return s.Name
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) GetType() *string {
	return s.Type
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetCheckActivity(v bool) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.CheckActivity = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetConfigItemList(v []*CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.ConfigItemList = v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetDescription(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetName(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetType(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) Validate() error {
	if s.ConfigItemList != nil {
		for _, item := range s.ConfigItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) GetKey() *string {
	return s.Key
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) GetValue() *string {
	return s.Value
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) SetKey(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList {
	s.Key = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) SetValue(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList {
	s.Value = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) Validate() error {
	return dara.Validate(s)
}
