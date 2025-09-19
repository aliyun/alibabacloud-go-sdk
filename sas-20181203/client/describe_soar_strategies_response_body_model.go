// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSoarStrategiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarStrategiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSoarStrategiesResponseBody
	GetRequestId() *string
	SetSoarStrategies(v []*DescribeSoarStrategiesResponseBodySoarStrategies) *DescribeSoarStrategiesResponseBody
	GetSoarStrategies() []*DescribeSoarStrategiesResponseBodySoarStrategies
	SetTotalCount(v int32) *DescribeSoarStrategiesResponseBody
	GetTotalCount() *int32
}

type DescribeSoarStrategiesResponseBody struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4BB99533-4FDC-5B9C-A5E4-5AE3E9BE5C78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The policies.
	SoarStrategies []*DescribeSoarStrategiesResponseBodySoarStrategies `json:"SoarStrategies,omitempty" xml:"SoarStrategies,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarStrategiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarStrategiesResponseBody) GetSoarStrategies() []*DescribeSoarStrategiesResponseBodySoarStrategies {
	return s.SoarStrategies
}

func (s *DescribeSoarStrategiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSoarStrategiesResponseBody) SetPageNumber(v int32) *DescribeSoarStrategiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBody) SetPageSize(v int32) *DescribeSoarStrategiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBody) SetRequestId(v string) *DescribeSoarStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBody) SetSoarStrategies(v []*DescribeSoarStrategiesResponseBodySoarStrategies) *DescribeSoarStrategiesResponseBody {
	s.SoarStrategies = v
	return s
}

func (s *DescribeSoarStrategiesResponseBody) SetTotalCount(v int32) *DescribeSoarStrategiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSoarStrategiesResponseBodySoarStrategies struct {
	// The Alibaba Cloud account ID of the creator. Default value: 0. The value indicates a system user.
	//
	// example:
	//
	// 0
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// strategy_description_01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the policy was created. Unit: milliseconds.
	//
	// example:
	//
	// 1703556715000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the policy was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1698114242000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 16064025
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// strategy_name01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- runmode_TRIGGER_BY_USER: manually executed
	//
	// example:
	//
	// runmode_TRIGGER_BY_USER
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- type_vulfix: vulnerability operations
	//
	// example:
	//
	// type_vulfix
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSoarStrategiesResponseBodySoarStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategiesResponseBodySoarStrategies) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetDescription() *string {
	return s.Description
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetId() *int64 {
	return s.Id
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetName() *string {
	return s.Name
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetRunMode() *string {
	return s.RunMode
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) GetType() *string {
	return s.Type
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetCreator(v string) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.Creator = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetDescription(v string) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.Description = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetGmtCreate(v int64) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetGmtModified(v int64) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.GmtModified = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetId(v int64) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.Id = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetName(v string) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.Name = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetRunMode(v string) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.RunMode = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) SetType(v string) *DescribeSoarStrategiesResponseBodySoarStrategies {
	s.Type = &v
	return s
}

func (s *DescribeSoarStrategiesResponseBodySoarStrategies) Validate() error {
	return dara.Validate(s)
}
