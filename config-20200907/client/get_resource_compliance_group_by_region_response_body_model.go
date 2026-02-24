// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceGroupByRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceResult(v *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) *GetResourceComplianceGroupByRegionResponseBody
	GetComplianceResult() *GetResourceComplianceGroupByRegionResponseBodyComplianceResult
	SetRequestId(v string) *GetResourceComplianceGroupByRegionResponseBody
	GetRequestId() *string
}

type GetResourceComplianceGroupByRegionResponseBody struct {
	ComplianceResult *GetResourceComplianceGroupByRegionResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBody) GetComplianceResult() *GetResourceComplianceGroupByRegionResponseBodyComplianceResult {
	return s.ComplianceResult
}

func (s *GetResourceComplianceGroupByRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceComplianceGroupByRegionResponseBody) SetComplianceResult(v *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) *GetResourceComplianceGroupByRegionResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBody) SetRequestId(v string) *GetResourceComplianceGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBody) Validate() error {
	if s.ComplianceResult != nil {
		if err := s.ComplianceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceComplianceGroupByRegionResponseBodyComplianceResult struct {
	ComplianceResultList []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResult) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) GetComplianceResultList() []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	return s.ComplianceResultList
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) SetComplianceResultList(v []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) *GetResourceComplianceGroupByRegionResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) Validate() error {
	if s.ComplianceResultList != nil {
		for _, item := range s.ComplianceResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList struct {
	Compliances []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	RegionId    *string                                                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GetCompliances() []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	return s.Compliances
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetRegionId(v string) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.RegionId = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) Validate() error {
	if s.Compliances != nil {
		for _, item := range s.Compliances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GetCount() *int64 {
	return s.Count
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) Validate() error {
	return dara.Validate(s)
}
