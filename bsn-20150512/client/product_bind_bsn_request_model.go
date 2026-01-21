// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductBindBsnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliuid(v int64) *ProductBindBsnRequest
	GetAliuid() *int64
	SetNum(v int32) *ProductBindBsnRequest
	GetNum() *int32
	SetResourceId(v string) *ProductBindBsnRequest
	GetResourceId() *string
	SetResourceType(v int32) *ProductBindBsnRequest
	GetResourceType() *int32
}

type ProductBindBsnRequest struct {
	// 456*******163
	//
	// This parameter is required.
	Aliuid *int64 `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	Num *int32 `json:"num,omitempty" xml:"num,omitempty"`
	// 41****34
	//
	// This parameter is required.
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ProductBindBsnRequest) String() string {
	return dara.Prettify(s)
}

func (s ProductBindBsnRequest) GoString() string {
	return s.String()
}

func (s *ProductBindBsnRequest) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *ProductBindBsnRequest) GetNum() *int32 {
	return s.Num
}

func (s *ProductBindBsnRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ProductBindBsnRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *ProductBindBsnRequest) SetAliuid(v int64) *ProductBindBsnRequest {
	s.Aliuid = &v
	return s
}

func (s *ProductBindBsnRequest) SetNum(v int32) *ProductBindBsnRequest {
	s.Num = &v
	return s
}

func (s *ProductBindBsnRequest) SetResourceId(v string) *ProductBindBsnRequest {
	s.ResourceId = &v
	return s
}

func (s *ProductBindBsnRequest) SetResourceType(v int32) *ProductBindBsnRequest {
	s.ResourceType = &v
	return s
}

func (s *ProductBindBsnRequest) Validate() error {
	return dara.Validate(s)
}
