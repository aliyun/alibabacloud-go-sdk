// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRiskNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRiskNum(v []*ListInstanceRiskNumResponseBodyInstanceRiskNum) *ListInstanceRiskNumResponseBody
	GetInstanceRiskNum() []*ListInstanceRiskNumResponseBodyInstanceRiskNum
	SetRequestId(v string) *ListInstanceRiskNumResponseBody
	GetRequestId() *string
}

type ListInstanceRiskNumResponseBody struct {
	// The information about the risks in the instance.
	InstanceRiskNum []*ListInstanceRiskNumResponseBodyInstanceRiskNum `json:"InstanceRiskNum,omitempty" xml:"InstanceRiskNum,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 291B49F9-1685-4005-9D34-606B6F78****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceRiskNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumResponseBody) GetInstanceRiskNum() []*ListInstanceRiskNumResponseBodyInstanceRiskNum {
	return s.InstanceRiskNum
}

func (s *ListInstanceRiskNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceRiskNumResponseBody) SetInstanceRiskNum(v []*ListInstanceRiskNumResponseBodyInstanceRiskNum) *ListInstanceRiskNumResponseBody {
	s.InstanceRiskNum = v
	return s
}

func (s *ListInstanceRiskNumResponseBody) SetRequestId(v string) *ListInstanceRiskNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceRiskNumResponseBody) Validate() error {
	if s.InstanceRiskNum != nil {
		for _, item := range s.InstanceRiskNum {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceRiskNumResponseBodyInstanceRiskNum struct {
	// The information about the instance.
	InstanceItem *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem `json:"InstanceItem,omitempty" xml:"InstanceItem,omitempty" type:"Struct"`
	// The statistics about the risks.
	RiskNumEntity *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity `json:"RiskNumEntity,omitempty" xml:"RiskNumEntity,omitempty" type:"Struct"`
}

func (s ListInstanceRiskNumResponseBodyInstanceRiskNum) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumResponseBodyInstanceRiskNum) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNum) GetInstanceItem() *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem {
	return s.InstanceItem
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNum) GetRiskNumEntity() *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	return s.RiskNumEntity
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNum) SetInstanceItem(v *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) *ListInstanceRiskNumResponseBodyInstanceRiskNum {
	s.InstanceItem = v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNum) SetRiskNumEntity(v *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) *ListInstanceRiskNumResponseBodyInstanceRiskNum {
	s.RiskNumEntity = v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNum) Validate() error {
	if s.InstanceItem != nil {
		if err := s.InstanceItem.Validate(); err != nil {
			return err
		}
	}
	if s.RiskNumEntity != nil {
		if err := s.RiskNumEntity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-wz9fdluqx20mp2x7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// f2d6e901-1004-4ca8-9dae-53ec04a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) GetUuid() *string {
	return s.Uuid
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) SetInstanceId(v string) *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) SetUuid(v string) *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem {
	s.Uuid = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumInstanceItem) Validate() error {
	return dara.Validate(s)
}

type ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity struct {
	// The number of high-risk alerts.
	//
	// example:
	//
	// 5
	SuspiciousHighCount *int32 `json:"SuspiciousHighCount,omitempty" xml:"SuspiciousHighCount,omitempty"`
	// The number of low-risk alerts.
	//
	// example:
	//
	// 7
	SuspiciousLowCount *int32 `json:"SuspiciousLowCount,omitempty" xml:"SuspiciousLowCount,omitempty"`
	// The number of medium-risk alerts.
	//
	// example:
	//
	// 6
	SuspiciousMediumCount *int32 `json:"SuspiciousMediumCount,omitempty" xml:"SuspiciousMediumCount,omitempty"`
	// The number of high-risk vulnerabilities.
	//
	// example:
	//
	// 1
	VulHighCount *int32 `json:"VulHighCount,omitempty" xml:"VulHighCount,omitempty"`
	// The number of low-risk vulnerabilities.
	//
	// example:
	//
	// 3
	VulLowCount *int32 `json:"VulLowCount,omitempty" xml:"VulLowCount,omitempty"`
	// The number of medium-risk vulnerabilities.
	//
	// example:
	//
	// 2
	VulMediumCount *int32 `json:"VulMediumCount,omitempty" xml:"VulMediumCount,omitempty"`
	// The number of weak passwords exposed on the Internet.
	//
	// example:
	//
	// 4
	WeakPassWordCount *int32 `json:"WeakPassWordCount,omitempty" xml:"WeakPassWordCount,omitempty"`
}

func (s ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GoString() string {
	return s.String()
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetSuspiciousHighCount() *int32 {
	return s.SuspiciousHighCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetSuspiciousLowCount() *int32 {
	return s.SuspiciousLowCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetSuspiciousMediumCount() *int32 {
	return s.SuspiciousMediumCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetVulHighCount() *int32 {
	return s.VulHighCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetVulLowCount() *int32 {
	return s.VulLowCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetVulMediumCount() *int32 {
	return s.VulMediumCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) GetWeakPassWordCount() *int32 {
	return s.WeakPassWordCount
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetSuspiciousHighCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.SuspiciousHighCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetSuspiciousLowCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.SuspiciousLowCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetSuspiciousMediumCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.SuspiciousMediumCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetVulHighCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.VulHighCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetVulLowCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.VulLowCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetVulMediumCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.VulMediumCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) SetWeakPassWordCount(v int32) *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity {
	s.WeakPassWordCount = &v
	return s
}

func (s *ListInstanceRiskNumResponseBodyInstanceRiskNumRiskNumEntity) Validate() error {
	return dara.Validate(s)
}
