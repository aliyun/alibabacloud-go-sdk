// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNASAvailableResourceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeNASAvailableResourceInfoResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeNASAvailableResourceInfoResponseBody
	GetMessage() *string
	SetNasAvailableResourceInfo(v []*DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) *DescribeNASAvailableResourceInfoResponseBody
	GetNasAvailableResourceInfo() []*DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo
	SetRequestId(v string) *DescribeNASAvailableResourceInfoResponseBody
	GetRequestId() *string
}

type DescribeNASAvailableResourceInfoResponseBody struct {
	// The business status code. A value of 0 indicates success.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// you are not authorized to this workspace, or workspace not exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The available NAS resource information.
	NasAvailableResourceInfo []*DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo `json:"NasAvailableResourceInfo,omitempty" xml:"NasAvailableResourceInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNASAvailableResourceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASAvailableResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNASAvailableResourceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNASAvailableResourceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNASAvailableResourceInfoResponseBody) GetNasAvailableResourceInfo() []*DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	return s.NasAvailableResourceInfo
}

func (s *DescribeNASAvailableResourceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNASAvailableResourceInfoResponseBody) SetCode(v string) *DescribeNASAvailableResourceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBody) SetMessage(v string) *DescribeNASAvailableResourceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBody) SetNasAvailableResourceInfo(v []*DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) *DescribeNASAvailableResourceInfoResponseBody {
	s.NasAvailableResourceInfo = v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBody) SetRequestId(v string) *DescribeNASAvailableResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBody) Validate() error {
	if s.NasAvailableResourceInfo != nil {
		for _, item := range s.NasAvailableResourceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo struct {
	// The product capabilities of the edge zone.
	Ability []*string `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Repeated"`
	// The area to which the ENS edge zone belongs.
	//
	// example:
	//
	// SouthWestChina
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The English name.
	//
	// example:
	//
	// cn-chenzhou-telecom_unicom_cmcc
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The ENS edge zone ID.
	//
	// example:
	//
	// cn-chenzhou-telecom_unicom_cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the ENS edge zone.
	//
	// example:
	//
	// China Hangzhou36
	EnsRegionName *string `json:"EnsRegionName,omitempty" xml:"EnsRegionName,omitempty"`
	// The resource count of active NAS resources.
	//
	// example:
	//
	// 1
	NasAvailableAmount *int32 `json:"NasAvailableAmount,omitempty" xml:"NasAvailableAmount,omitempty"`
	// The type of available NAS resources.
	//
	// example:
	//
	// capacity
	NasAvailableStorgeType *string `json:"NasAvailableStorgeType,omitempty" xml:"NasAvailableStorgeType,omitempty"`
	// The province to which the ENS edge zone belongs.
	//
	// example:
	//
	// Hubei Province
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GoString() string {
	return s.String()
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetAbility() []*string {
	return s.Ability
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetArea() *string {
	return s.Area
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetEnName() *string {
	return s.EnName
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetEnsRegionName() *string {
	return s.EnsRegionName
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetNasAvailableAmount() *int32 {
	return s.NasAvailableAmount
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetNasAvailableStorgeType() *string {
	return s.NasAvailableStorgeType
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) GetProvince() *string {
	return s.Province
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetAbility(v []*string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.Ability = v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetArea(v string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.Area = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetEnName(v string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.EnName = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetEnsRegionId(v string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetEnsRegionName(v string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.EnsRegionName = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetNasAvailableAmount(v int32) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.NasAvailableAmount = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetNasAvailableStorgeType(v string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.NasAvailableStorgeType = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) SetProvince(v string) *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo {
	s.Province = &v
	return s
}

func (s *DescribeNASAvailableResourceInfoResponseBodyNasAvailableResourceInfo) Validate() error {
	return dara.Validate(s)
}
