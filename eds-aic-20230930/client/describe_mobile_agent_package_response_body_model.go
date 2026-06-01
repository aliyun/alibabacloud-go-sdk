// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMobileAgentPackageResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMobileAgentPackageResponseBody
	GetMessage() *string
	SetPackageList(v []*DescribeMobileAgentPackageResponseBodyPackageList) *DescribeMobileAgentPackageResponseBody
	GetPackageList() []*DescribeMobileAgentPackageResponseBodyPackageList
	SetRequestId(v string) *DescribeMobileAgentPackageResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeMobileAgentPackageResponseBody
	GetTotalCount() *string
}

type DescribeMobileAgentPackageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success.
	Message     *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PackageList []*DescribeMobileAgentPackageResponseBodyPackageList `json:"PackageList,omitempty" xml:"PackageList,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMobileAgentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMobileAgentPackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMobileAgentPackageResponseBody) GetPackageList() []*DescribeMobileAgentPackageResponseBodyPackageList {
	return s.PackageList
}

func (s *DescribeMobileAgentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMobileAgentPackageResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeMobileAgentPackageResponseBody) SetCode(v string) *DescribeMobileAgentPackageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBody) SetMessage(v string) *DescribeMobileAgentPackageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBody) SetPackageList(v []*DescribeMobileAgentPackageResponseBodyPackageList) *DescribeMobileAgentPackageResponseBody {
	s.PackageList = v
	return s
}

func (s *DescribeMobileAgentPackageResponseBody) SetRequestId(v string) *DescribeMobileAgentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBody) SetTotalCount(v string) *DescribeMobileAgentPackageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBody) Validate() error {
	if s.PackageList != nil {
		for _, item := range s.PackageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMobileAgentPackageResponseBodyPackageList struct {
	// example:
	//
	// 2026-10-30 00:00:00
	ExpiredAt   *string   `json:"ExpiredAt,omitempty" xml:"ExpiredAt,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 8000
	PackageCredit *string `json:"PackageCredit,omitempty" xml:"PackageCredit,omitempty"`
	// example:
	//
	// cmag-bp19i1yxu60r7twy****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// advanced
	PackageSpec *string `json:"PackageSpec,omitempty" xml:"PackageSpec,omitempty"`
	// example:
	//
	// ACTIVE
	PackageStatus *string `json:"PackageStatus,omitempty" xml:"PackageStatus,omitempty"`
	// example:
	//
	// 1000
	UsedCredit *string `json:"UsedCredit,omitempty" xml:"UsedCredit,omitempty"`
}

func (s DescribeMobileAgentPackageResponseBodyPackageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackageResponseBodyPackageList) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetExpiredAt() *string {
	return s.ExpiredAt
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetPackageCredit() *string {
	return s.PackageCredit
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetPackageSpec() *string {
	return s.PackageSpec
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetPackageStatus() *string {
	return s.PackageStatus
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) GetUsedCredit() *string {
	return s.UsedCredit
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetExpiredAt(v string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.ExpiredAt = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetInstanceIds(v []*string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.InstanceIds = v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetPackageCredit(v string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.PackageCredit = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetPackageId(v string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.PackageId = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetPackageSpec(v string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.PackageSpec = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetPackageStatus(v string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.PackageStatus = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) SetUsedCredit(v string) *DescribeMobileAgentPackageResponseBodyPackageList {
	s.UsedCredit = &v
	return s
}

func (s *DescribeMobileAgentPackageResponseBodyPackageList) Validate() error {
	return dara.Validate(s)
}
