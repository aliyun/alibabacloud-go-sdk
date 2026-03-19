// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLFunction interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DLFunction
	GetCatalogName() *string
	SetClassName(v string) *DLFunction
	GetClassName() *string
	SetCreateTime(v int32) *DLFunction
	GetCreateTime() *int32
	SetCreatorId(v int64) *DLFunction
	GetCreatorId() *int64
	SetDbName(v string) *DLFunction
	GetDbName() *string
	SetFunctionName(v string) *DLFunction
	GetFunctionName() *string
	SetFunctionType(v string) *DLFunction
	GetFunctionType() *string
	SetModifierId(v int64) *DLFunction
	GetModifierId() *int64
	SetOwnerName(v string) *DLFunction
	GetOwnerName() *string
	SetOwnerType(v string) *DLFunction
	GetOwnerType() *string
	SetResourceUris(v []*DLResourceUri) *DLFunction
	GetResourceUris() []*DLResourceUri
}

type DLFunction struct {
	// The name of the data catalog.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The JAVA class that contains the function code.
	//
	// example:
	//
	// com.example.hive.MyCustomUDF
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The time when the function was created.
	//
	// example:
	//
	// 1731586286
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the function.
	//
	// example:
	//
	// 12****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the queried database.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the function.
	//
	// example:
	//
	// JAVA
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The ID of the user who modified the function.
	//
	// example:
	//
	// 26****
	ModifierId *int64 `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The owner of the function.
	//
	// example:
	//
	// zhangsan
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The type of the owner. Valid values:
	//
	// 	- ROLE
	//
	// 	- GROUP
	//
	// 	- USER
	//
	// example:
	//
	// USER
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The resource URIs of the function. You cannot modify the resource URIs after the function is created.
	ResourceUris []*DLResourceUri `json:"ResourceUris,omitempty" xml:"ResourceUris,omitempty" type:"Repeated"`
}

func (s DLFunction) String() string {
	return dara.Prettify(s)
}

func (s DLFunction) GoString() string {
	return s.String()
}

func (s *DLFunction) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DLFunction) GetClassName() *string {
	return s.ClassName
}

func (s *DLFunction) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLFunction) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *DLFunction) GetDbName() *string {
	return s.DbName
}

func (s *DLFunction) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DLFunction) GetFunctionType() *string {
	return s.FunctionType
}

func (s *DLFunction) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *DLFunction) GetOwnerName() *string {
	return s.OwnerName
}

func (s *DLFunction) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DLFunction) GetResourceUris() []*DLResourceUri {
	return s.ResourceUris
}

func (s *DLFunction) SetCatalogName(v string) *DLFunction {
	s.CatalogName = &v
	return s
}

func (s *DLFunction) SetClassName(v string) *DLFunction {
	s.ClassName = &v
	return s
}

func (s *DLFunction) SetCreateTime(v int32) *DLFunction {
	s.CreateTime = &v
	return s
}

func (s *DLFunction) SetCreatorId(v int64) *DLFunction {
	s.CreatorId = &v
	return s
}

func (s *DLFunction) SetDbName(v string) *DLFunction {
	s.DbName = &v
	return s
}

func (s *DLFunction) SetFunctionName(v string) *DLFunction {
	s.FunctionName = &v
	return s
}

func (s *DLFunction) SetFunctionType(v string) *DLFunction {
	s.FunctionType = &v
	return s
}

func (s *DLFunction) SetModifierId(v int64) *DLFunction {
	s.ModifierId = &v
	return s
}

func (s *DLFunction) SetOwnerName(v string) *DLFunction {
	s.OwnerName = &v
	return s
}

func (s *DLFunction) SetOwnerType(v string) *DLFunction {
	s.OwnerType = &v
	return s
}

func (s *DLFunction) SetResourceUris(v []*DLResourceUri) *DLFunction {
	s.ResourceUris = v
	return s
}

func (s *DLFunction) Validate() error {
	if s.ResourceUris != nil {
		for _, item := range s.ResourceUris {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
