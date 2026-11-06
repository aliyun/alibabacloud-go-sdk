// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListInstanceRequest
	GetName() *string
	SetNum(v int64) *ListInstanceRequest
	GetNum() *int64
	SetPage(v int64) *ListInstanceRequest
	GetPage() *int64
	SetSaleId(v string) *ListInstanceRequest
	GetSaleId() *string
	SetStatus(v string) *ListInstanceRequest
	GetStatus() *string
}

type ListInstanceRequest struct {
	// The name of the instance.
	//
	// example:
	//
	// Test instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the diversion instance.
	//
	// example:
	//
	// ddos_diversion_public_cn-xxxxxxxxxxxxx
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The instance status. Valid values:
	//
	// - normal: The instance is running normally.
	//
	// - expired: The instance has expired.
	//
	// - deleting: The instance is being deleted.
	//
	// - stopped: The instance is stopped.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) GetName() *string {
	return s.Name
}

func (s *ListInstanceRequest) GetNum() *int64 {
	return s.Num
}

func (s *ListInstanceRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListInstanceRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *ListInstanceRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceRequest) SetName(v string) *ListInstanceRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceRequest) SetNum(v int64) *ListInstanceRequest {
	s.Num = &v
	return s
}

func (s *ListInstanceRequest) SetPage(v int64) *ListInstanceRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceRequest) SetSaleId(v string) *ListInstanceRequest {
	s.SaleId = &v
	return s
}

func (s *ListInstanceRequest) SetStatus(v string) *ListInstanceRequest {
	s.Status = &v
	return s
}

func (s *ListInstanceRequest) Validate() error {
	return dara.Validate(s)
}
