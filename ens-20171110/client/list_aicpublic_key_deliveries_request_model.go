// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICPublicKeyDeliveriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAICPublicKeyDeliveriesRequest
	GetInstanceId() *string
	SetKeyGroup(v string) *ListAICPublicKeyDeliveriesRequest
	GetKeyGroup() *string
	SetKeyName(v string) *ListAICPublicKeyDeliveriesRequest
	GetKeyName() *string
	SetKeyType(v string) *ListAICPublicKeyDeliveriesRequest
	GetKeyType() *string
	SetPageNumber(v int32) *ListAICPublicKeyDeliveriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAICPublicKeyDeliveriesRequest
	GetPageSize() *int32
}

type ListAICPublicKeyDeliveriesRequest struct {
	// example:
	//
	// aic-xxxx-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// adb
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAICPublicKeyDeliveriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeyDeliveriesRequest) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeyDeliveriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAICPublicKeyDeliveriesRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ListAICPublicKeyDeliveriesRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ListAICPublicKeyDeliveriesRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *ListAICPublicKeyDeliveriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAICPublicKeyDeliveriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAICPublicKeyDeliveriesRequest) SetInstanceId(v string) *ListAICPublicKeyDeliveriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesRequest) SetKeyGroup(v string) *ListAICPublicKeyDeliveriesRequest {
	s.KeyGroup = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesRequest) SetKeyName(v string) *ListAICPublicKeyDeliveriesRequest {
	s.KeyName = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesRequest) SetKeyType(v string) *ListAICPublicKeyDeliveriesRequest {
	s.KeyType = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesRequest) SetPageNumber(v int32) *ListAICPublicKeyDeliveriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesRequest) SetPageSize(v int32) *ListAICPublicKeyDeliveriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesRequest) Validate() error {
	return dara.Validate(s)
}
