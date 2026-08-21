// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryShelfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketQueryShelfResponseBodyData) *TicketQueryShelfResponseBody
	GetData() *TicketQueryShelfResponseBodyData
	SetErrorCode(v string) *TicketQueryShelfResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketQueryShelfResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketQueryShelfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketQueryShelfResponseBody
	GetSuccess() *bool
}

type TicketQueryShelfResponseBody struct {
	Data *TicketQueryShelfResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ScenicIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ScenicId不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketQueryShelfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfResponseBody) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfResponseBody) GetData() *TicketQueryShelfResponseBodyData {
	return s.Data
}

func (s *TicketQueryShelfResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketQueryShelfResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketQueryShelfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketQueryShelfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketQueryShelfResponseBody) SetData(v *TicketQueryShelfResponseBodyData) *TicketQueryShelfResponseBody {
	s.Data = v
	return s
}

func (s *TicketQueryShelfResponseBody) SetErrorCode(v string) *TicketQueryShelfResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryShelfResponseBody) SetErrorMsg(v string) *TicketQueryShelfResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketQueryShelfResponseBody) SetRequestId(v string) *TicketQueryShelfResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketQueryShelfResponseBody) SetSuccess(v bool) *TicketQueryShelfResponseBody {
	s.Success = &v
	return s
}

func (s *TicketQueryShelfResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryShelfResponseBodyData struct {
	Shelves []*TicketQueryShelfResponseBodyDataShelves `json:"Shelves,omitempty" xml:"Shelves,omitempty" type:"Repeated"`
}

func (s TicketQueryShelfResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfResponseBodyData) GetShelves() []*TicketQueryShelfResponseBodyDataShelves {
	return s.Shelves
}

func (s *TicketQueryShelfResponseBodyData) SetShelves(v []*TicketQueryShelfResponseBodyDataShelves) *TicketQueryShelfResponseBodyData {
	s.Shelves = v
	return s
}

func (s *TicketQueryShelfResponseBodyData) Validate() error {
	if s.Shelves != nil {
		for _, item := range s.Shelves {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketQueryShelfResponseBodyDataShelves struct {
	// example:
	//
	// 1951890
	ShelfId *int64 `json:"ShelfId,omitempty" xml:"ShelfId,omitempty"`
	// example:
	//
	// 0
	ShelfIndex *int32 `json:"ShelfIndex,omitempty" xml:"ShelfIndex,omitempty"`
	// example:
	//
	// 1日门票
	ShelfName *string                                        `json:"ShelfName,omitempty" xml:"ShelfName,omitempty"`
	Tabs      []*TicketQueryShelfResponseBodyDataShelvesTabs `json:"Tabs,omitempty" xml:"Tabs,omitempty" type:"Repeated"`
}

func (s TicketQueryShelfResponseBodyDataShelves) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfResponseBodyDataShelves) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfResponseBodyDataShelves) GetShelfId() *int64 {
	return s.ShelfId
}

func (s *TicketQueryShelfResponseBodyDataShelves) GetShelfIndex() *int32 {
	return s.ShelfIndex
}

func (s *TicketQueryShelfResponseBodyDataShelves) GetShelfName() *string {
	return s.ShelfName
}

func (s *TicketQueryShelfResponseBodyDataShelves) GetTabs() []*TicketQueryShelfResponseBodyDataShelvesTabs {
	return s.Tabs
}

func (s *TicketQueryShelfResponseBodyDataShelves) SetShelfId(v int64) *TicketQueryShelfResponseBodyDataShelves {
	s.ShelfId = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelves) SetShelfIndex(v int32) *TicketQueryShelfResponseBodyDataShelves {
	s.ShelfIndex = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelves) SetShelfName(v string) *TicketQueryShelfResponseBodyDataShelves {
	s.ShelfName = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelves) SetTabs(v []*TicketQueryShelfResponseBodyDataShelvesTabs) *TicketQueryShelfResponseBodyDataShelves {
	s.Tabs = v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelves) Validate() error {
	if s.Tabs != nil {
		for _, item := range s.Tabs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketQueryShelfResponseBodyDataShelvesTabs struct {
	Cells []*TicketQueryShelfResponseBodyDataShelvesTabsCells `json:"Cells,omitempty" xml:"Cells,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TabIndex *int32 `json:"TabIndex,omitempty" xml:"TabIndex,omitempty"`
	// example:
	//
	// 景点门票
	TabName *string `json:"TabName,omitempty" xml:"TabName,omitempty"`
}

func (s TicketQueryShelfResponseBodyDataShelvesTabs) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfResponseBodyDataShelvesTabs) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) GetCells() []*TicketQueryShelfResponseBodyDataShelvesTabsCells {
	return s.Cells
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) GetTabIndex() *int32 {
	return s.TabIndex
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) GetTabName() *string {
	return s.TabName
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) SetCells(v []*TicketQueryShelfResponseBodyDataShelvesTabsCells) *TicketQueryShelfResponseBodyDataShelvesTabs {
	s.Cells = v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) SetTabIndex(v int32) *TicketQueryShelfResponseBodyDataShelvesTabs {
	s.TabIndex = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) SetTabName(v string) *TicketQueryShelfResponseBodyDataShelvesTabs {
	s.TabName = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabs) Validate() error {
	if s.Cells != nil {
		for _, item := range s.Cells {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketQueryShelfResponseBodyDataShelvesTabsCells struct {
	// example:
	//
	// 3507
	SpuId *int64 `json:"SpuId,omitempty" xml:"SpuId,omitempty"`
	// example:
	//
	// 60484007
	TicketKindId *int64 `json:"TicketKindId,omitempty" xml:"TicketKindId,omitempty"`
}

func (s TicketQueryShelfResponseBodyDataShelvesTabsCells) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfResponseBodyDataShelvesTabsCells) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabsCells) GetSpuId() *int64 {
	return s.SpuId
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabsCells) GetTicketKindId() *int64 {
	return s.TicketKindId
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabsCells) SetSpuId(v int64) *TicketQueryShelfResponseBodyDataShelvesTabsCells {
	s.SpuId = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabsCells) SetTicketKindId(v int64) *TicketQueryShelfResponseBodyDataShelvesTabsCells {
	s.TicketKindId = &v
	return s
}

func (s *TicketQueryShelfResponseBodyDataShelvesTabsCells) Validate() error {
	return dara.Validate(s)
}
