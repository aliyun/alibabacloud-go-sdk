// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFreeResourceDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *FreeResourceDetail
	GetAmount() *int32
	SetResourceType(v string) *FreeResourceDetail
	GetResourceType() *string
}

type FreeResourceDetail struct {
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// CPU
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s FreeResourceDetail) String() string {
	return dara.Prettify(s)
}

func (s FreeResourceDetail) GoString() string {
	return s.String()
}

func (s *FreeResourceDetail) GetAmount() *int32 {
	return s.Amount
}

func (s *FreeResourceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *FreeResourceDetail) SetAmount(v int32) *FreeResourceDetail {
	s.Amount = &v
	return s
}

func (s *FreeResourceDetail) SetResourceType(v string) *FreeResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *FreeResourceDetail) Validate() error {
	return dara.Validate(s)
}
