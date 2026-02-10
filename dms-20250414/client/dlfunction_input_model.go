// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLFunctionInput interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *DLFunctionInput
	GetClassName() *string
	SetCreateTime(v int32) *DLFunctionInput
	GetCreateTime() *int32
	SetCreatorId(v int64) *DLFunctionInput
	GetCreatorId() *int64
	SetFunctionName(v string) *DLFunctionInput
	GetFunctionName() *string
	SetFunctionType(v string) *DLFunctionInput
	GetFunctionType() *string
	SetModifierId(v int64) *DLFunctionInput
	GetModifierId() *int64
	SetOwnerName(v string) *DLFunctionInput
	GetOwnerName() *string
	SetOwnerType(v string) *DLFunctionInput
	GetOwnerType() *string
	SetResourceUris(v []*DLResourceUri) *DLFunctionInput
	GetResourceUris() []*DLResourceUri
}

type DLFunctionInput struct {
	// example:
	//
	// com.example.hive.MyCustomUDF
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// example:
	//
	// 1731586286
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 12****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// my_funciton
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// JAVA
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// example:
	//
	// 26****
	ModifierId *int64 `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// zhangsan
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// USER
	OwnerType    *string          `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	ResourceUris []*DLResourceUri `json:"ResourceUris,omitempty" xml:"ResourceUris,omitempty" type:"Repeated"`
}

func (s DLFunctionInput) String() string {
	return dara.Prettify(s)
}

func (s DLFunctionInput) GoString() string {
	return s.String()
}

func (s *DLFunctionInput) GetClassName() *string {
	return s.ClassName
}

func (s *DLFunctionInput) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLFunctionInput) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *DLFunctionInput) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DLFunctionInput) GetFunctionType() *string {
	return s.FunctionType
}

func (s *DLFunctionInput) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *DLFunctionInput) GetOwnerName() *string {
	return s.OwnerName
}

func (s *DLFunctionInput) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DLFunctionInput) GetResourceUris() []*DLResourceUri {
	return s.ResourceUris
}

func (s *DLFunctionInput) SetClassName(v string) *DLFunctionInput {
	s.ClassName = &v
	return s
}

func (s *DLFunctionInput) SetCreateTime(v int32) *DLFunctionInput {
	s.CreateTime = &v
	return s
}

func (s *DLFunctionInput) SetCreatorId(v int64) *DLFunctionInput {
	s.CreatorId = &v
	return s
}

func (s *DLFunctionInput) SetFunctionName(v string) *DLFunctionInput {
	s.FunctionName = &v
	return s
}

func (s *DLFunctionInput) SetFunctionType(v string) *DLFunctionInput {
	s.FunctionType = &v
	return s
}

func (s *DLFunctionInput) SetModifierId(v int64) *DLFunctionInput {
	s.ModifierId = &v
	return s
}

func (s *DLFunctionInput) SetOwnerName(v string) *DLFunctionInput {
	s.OwnerName = &v
	return s
}

func (s *DLFunctionInput) SetOwnerType(v string) *DLFunctionInput {
	s.OwnerType = &v
	return s
}

func (s *DLFunctionInput) SetResourceUris(v []*DLResourceUri) *DLFunctionInput {
	s.ResourceUris = v
	return s
}

func (s *DLFunctionInput) Validate() error {
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
