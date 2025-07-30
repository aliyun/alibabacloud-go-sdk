// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBusinessCategoryBasicInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBid(v int32) *BusinessCategoryBasicInfo
	GetBid() *int32
	SetName(v string) *BusinessCategoryBasicInfo
	GetName() *string
	SetOriginalId(v int64) *BusinessCategoryBasicInfo
	GetOriginalId() *int64
	SetServiceType(v int32) *BusinessCategoryBasicInfo
	GetServiceType() *int32
}

type BusinessCategoryBasicInfo struct {
	Bid         *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginalId  *int64  `json:"OriginalId,omitempty" xml:"OriginalId,omitempty"`
	ServiceType *int32  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s BusinessCategoryBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *BusinessCategoryBasicInfo) GetBid() *int32 {
	return s.Bid
}

func (s *BusinessCategoryBasicInfo) GetName() *string {
	return s.Name
}

func (s *BusinessCategoryBasicInfo) GetOriginalId() *int64 {
	return s.OriginalId
}

func (s *BusinessCategoryBasicInfo) GetServiceType() *int32 {
	return s.ServiceType
}

func (s *BusinessCategoryBasicInfo) SetBid(v int32) *BusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *BusinessCategoryBasicInfo) SetName(v string) *BusinessCategoryBasicInfo {
	s.Name = &v
	return s
}

func (s *BusinessCategoryBasicInfo) SetOriginalId(v int64) *BusinessCategoryBasicInfo {
	s.OriginalId = &v
	return s
}

func (s *BusinessCategoryBasicInfo) SetServiceType(v int32) *BusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

func (s *BusinessCategoryBasicInfo) Validate() error {
	return dara.Validate(s)
}
