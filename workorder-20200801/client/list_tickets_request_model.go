// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAfterTime(v string) *ListTicketsRequest
	GetCreatedAfterTime() *string
	SetCreatedBeforeTime(v string) *ListTicketsRequest
	GetCreatedBeforeTime() *string
	SetExtraConditionList(v []*ListTicketsRequestExtraConditionList) *ListTicketsRequest
	GetExtraConditionList() []*ListTicketsRequestExtraConditionList
	SetIds(v string) *ListTicketsRequest
	GetIds() *string
	SetPageSize(v int32) *ListTicketsRequest
	GetPageSize() *int32
	SetPageStart(v int32) *ListTicketsRequest
	GetPageStart() *int32
	SetProductCode(v string) *ListTicketsRequest
	GetProductCode() *string
	SetTicketStatus(v string) *ListTicketsRequest
	GetTicketStatus() *string
}

type ListTicketsRequest struct {
	CreatedAfterTime   *string                                 `json:"CreatedAfterTime,omitempty" xml:"CreatedAfterTime,omitempty"`
	CreatedBeforeTime  *string                                 `json:"CreatedBeforeTime,omitempty" xml:"CreatedBeforeTime,omitempty"`
	ExtraConditionList []*ListTicketsRequestExtraConditionList `json:"ExtraConditionList,omitempty" xml:"ExtraConditionList,omitempty" type:"Repeated"`
	Ids                *string                                 `json:"Ids,omitempty" xml:"Ids,omitempty"`
	PageSize           *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart          *int32                                  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	ProductCode        *string                                 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	TicketStatus       *string                                 `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
}

func (s ListTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) GetCreatedAfterTime() *string {
	return s.CreatedAfterTime
}

func (s *ListTicketsRequest) GetCreatedBeforeTime() *string {
	return s.CreatedBeforeTime
}

func (s *ListTicketsRequest) GetExtraConditionList() []*ListTicketsRequestExtraConditionList {
	return s.ExtraConditionList
}

func (s *ListTicketsRequest) GetIds() *string {
	return s.Ids
}

func (s *ListTicketsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsRequest) GetPageStart() *int32 {
	return s.PageStart
}

func (s *ListTicketsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListTicketsRequest) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *ListTicketsRequest) SetCreatedAfterTime(v string) *ListTicketsRequest {
	s.CreatedAfterTime = &v
	return s
}

func (s *ListTicketsRequest) SetCreatedBeforeTime(v string) *ListTicketsRequest {
	s.CreatedBeforeTime = &v
	return s
}

func (s *ListTicketsRequest) SetExtraConditionList(v []*ListTicketsRequestExtraConditionList) *ListTicketsRequest {
	s.ExtraConditionList = v
	return s
}

func (s *ListTicketsRequest) SetIds(v string) *ListTicketsRequest {
	s.Ids = &v
	return s
}

func (s *ListTicketsRequest) SetPageSize(v int32) *ListTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequest) SetPageStart(v int32) *ListTicketsRequest {
	s.PageStart = &v
	return s
}

func (s *ListTicketsRequest) SetProductCode(v string) *ListTicketsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListTicketsRequest) SetTicketStatus(v string) *ListTicketsRequest {
	s.TicketStatus = &v
	return s
}

func (s *ListTicketsRequest) Validate() error {
	if s.ExtraConditionList != nil {
		for _, item := range s.ExtraConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketsRequestExtraConditionList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTicketsRequestExtraConditionList) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequestExtraConditionList) GoString() string {
	return s.String()
}

func (s *ListTicketsRequestExtraConditionList) GetName() *string {
	return s.Name
}

func (s *ListTicketsRequestExtraConditionList) GetValue() *string {
	return s.Value
}

func (s *ListTicketsRequestExtraConditionList) SetName(v string) *ListTicketsRequestExtraConditionList {
	s.Name = &v
	return s
}

func (s *ListTicketsRequestExtraConditionList) SetValue(v string) *ListTicketsRequestExtraConditionList {
	s.Value = &v
	return s
}

func (s *ListTicketsRequestExtraConditionList) Validate() error {
	return dara.Validate(s)
}
