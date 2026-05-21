// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonthlySlaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryMonthlySlaListRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*QueryMonthlySlaListRequestEcIdAccountIds) *QueryMonthlySlaListRequest
	GetEcIdAccountIds() []*QueryMonthlySlaListRequestEcIdAccountIds
	SetInstanceIds(v []*string) *QueryMonthlySlaListRequest
	GetInstanceIds() []*string
	SetMonths(v []*int32) *QueryMonthlySlaListRequest
	GetMonths() []*int32
	SetNbid(v string) *QueryMonthlySlaListRequest
	GetNbid() *string
	SetPageSize(v int32) *QueryMonthlySlaListRequest
	GetPageSize() *int32
	SetPayStatuses(v []*int32) *QueryMonthlySlaListRequest
	GetPayStatuses() []*int32
	SetProductCodes(v []*string) *QueryMonthlySlaListRequest
	GetProductCodes() []*string
}

type QueryMonthlySlaListRequest struct {
	// example:
	//
	// 1
	CurrentPage    *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*QueryMonthlySlaListRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// ["instance_1","instance_2"]
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// [202602,202603]
	Months []*int32 `json:"Months,omitempty" xml:"Months,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PayStatuses []*int32 `json:"PayStatuses,omitempty" xml:"PayStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// ["ecs","oss"]
	ProductCodes []*string `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty" type:"Repeated"`
}

func (s QueryMonthlySlaListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthlySlaListRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthlySlaListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMonthlySlaListRequest) GetEcIdAccountIds() []*QueryMonthlySlaListRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *QueryMonthlySlaListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *QueryMonthlySlaListRequest) GetMonths() []*int32 {
	return s.Months
}

func (s *QueryMonthlySlaListRequest) GetNbid() *string {
	return s.Nbid
}

func (s *QueryMonthlySlaListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMonthlySlaListRequest) GetPayStatuses() []*int32 {
	return s.PayStatuses
}

func (s *QueryMonthlySlaListRequest) GetProductCodes() []*string {
	return s.ProductCodes
}

func (s *QueryMonthlySlaListRequest) SetCurrentPage(v int32) *QueryMonthlySlaListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMonthlySlaListRequest) SetEcIdAccountIds(v []*QueryMonthlySlaListRequestEcIdAccountIds) *QueryMonthlySlaListRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryMonthlySlaListRequest) SetInstanceIds(v []*string) *QueryMonthlySlaListRequest {
	s.InstanceIds = v
	return s
}

func (s *QueryMonthlySlaListRequest) SetMonths(v []*int32) *QueryMonthlySlaListRequest {
	s.Months = v
	return s
}

func (s *QueryMonthlySlaListRequest) SetNbid(v string) *QueryMonthlySlaListRequest {
	s.Nbid = &v
	return s
}

func (s *QueryMonthlySlaListRequest) SetPageSize(v int32) *QueryMonthlySlaListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMonthlySlaListRequest) SetPayStatuses(v []*int32) *QueryMonthlySlaListRequest {
	s.PayStatuses = v
	return s
}

func (s *QueryMonthlySlaListRequest) SetProductCodes(v []*string) *QueryMonthlySlaListRequest {
	s.ProductCodes = v
	return s
}

func (s *QueryMonthlySlaListRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMonthlySlaListRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryMonthlySlaListRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthlySlaListRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryMonthlySlaListRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *QueryMonthlySlaListRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *QueryMonthlySlaListRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryMonthlySlaListRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryMonthlySlaListRequestEcIdAccountIds) SetEcId(v string) *QueryMonthlySlaListRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *QueryMonthlySlaListRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
