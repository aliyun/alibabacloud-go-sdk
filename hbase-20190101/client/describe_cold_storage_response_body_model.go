// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeColdStorageResponseBody
	GetClusterId() *string
	SetColdStorageSize(v string) *DescribeColdStorageResponseBody
	GetColdStorageSize() *string
	SetColdStorageType(v string) *DescribeColdStorageResponseBody
	GetColdStorageType() *string
	SetColdStorageUseAmount(v string) *DescribeColdStorageResponseBody
	GetColdStorageUseAmount() *string
	SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody
	GetColdStorageUsePercent() *string
	SetOpenStatus(v string) *DescribeColdStorageResponseBody
	GetOpenStatus() *string
	SetPayType(v string) *DescribeColdStorageResponseBody
	GetPayType() *string
	SetRequestId(v string) *DescribeColdStorageResponseBody
	GetRequestId() *string
}

type DescribeColdStorageResponseBody struct {
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 800
	ColdStorageSize *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// example:
	//
	// BdsColdStorage
	ColdStorageType *string `json:"ColdStorageType,omitempty" xml:"ColdStorageType,omitempty"`
	// example:
	//
	// 20.00
	ColdStorageUseAmount *string `json:"ColdStorageUseAmount,omitempty" xml:"ColdStorageUseAmount,omitempty"`
	// example:
	//
	// 20.00
	ColdStorageUsePercent *string `json:"ColdStorageUsePercent,omitempty" xml:"ColdStorageUsePercent,omitempty"`
	// example:
	//
	// open
	OpenStatus *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// DCB9479E-F05F-4D1C-AFB7-C639B87764B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColdStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeColdStorageResponseBody) GetColdStorageSize() *string {
	return s.ColdStorageSize
}

func (s *DescribeColdStorageResponseBody) GetColdStorageType() *string {
	return s.ColdStorageType
}

func (s *DescribeColdStorageResponseBody) GetColdStorageUseAmount() *string {
	return s.ColdStorageUseAmount
}

func (s *DescribeColdStorageResponseBody) GetColdStorageUsePercent() *string {
	return s.ColdStorageUsePercent
}

func (s *DescribeColdStorageResponseBody) GetOpenStatus() *string {
	return s.OpenStatus
}

func (s *DescribeColdStorageResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeColdStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColdStorageResponseBody) SetClusterId(v string) *DescribeColdStorageResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageSize(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageType(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUseAmount(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUseAmount = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUsePercent = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetOpenStatus(v string) *DescribeColdStorageResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetPayType(v string) *DescribeColdStorageResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetRequestId(v string) *DescribeColdStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
