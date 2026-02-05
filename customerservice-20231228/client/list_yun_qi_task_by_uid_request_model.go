// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYunQiTaskByUidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v int64) *ListYunQiTaskByUidRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *ListYunQiTaskByUidRequest
	GetCreateTimeStart() *int64
	SetFreeOrderApplyCodes(v []*string) *ListYunQiTaskByUidRequest
	GetFreeOrderApplyCodes() []*string
	SetFreeOrderApplyIds(v []*int64) *ListYunQiTaskByUidRequest
	GetFreeOrderApplyIds() []*int64
	SetOrderIds(v []*int64) *ListYunQiTaskByUidRequest
	GetOrderIds() []*int64
	SetPageNum(v int32) *ListYunQiTaskByUidRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListYunQiTaskByUidRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *ListYunQiTaskByUidRequest
	GetStatusList() []*string
}

type ListYunQiTaskByUidRequest struct {
	CreateTimeEnd       *int64    `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	CreateTimeStart     *int64    `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	FreeOrderApplyCodes []*string `json:"freeOrderApplyCodes,omitempty" xml:"freeOrderApplyCodes,omitempty" type:"Repeated"`
	FreeOrderApplyIds   []*int64  `json:"freeOrderApplyIds,omitempty" xml:"freeOrderApplyIds,omitempty" type:"Repeated"`
	OrderIds            []*int64  `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	PageNum             *int32    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize            *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatusList          []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
}

func (s ListYunQiTaskByUidRequest) String() string {
	return dara.Prettify(s)
}

func (s ListYunQiTaskByUidRequest) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *ListYunQiTaskByUidRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *ListYunQiTaskByUidRequest) GetFreeOrderApplyCodes() []*string {
	return s.FreeOrderApplyCodes
}

func (s *ListYunQiTaskByUidRequest) GetFreeOrderApplyIds() []*int64 {
	return s.FreeOrderApplyIds
}

func (s *ListYunQiTaskByUidRequest) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *ListYunQiTaskByUidRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListYunQiTaskByUidRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListYunQiTaskByUidRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListYunQiTaskByUidRequest) SetCreateTimeEnd(v int64) *ListYunQiTaskByUidRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetCreateTimeStart(v int64) *ListYunQiTaskByUidRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetFreeOrderApplyCodes(v []*string) *ListYunQiTaskByUidRequest {
	s.FreeOrderApplyCodes = v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetFreeOrderApplyIds(v []*int64) *ListYunQiTaskByUidRequest {
	s.FreeOrderApplyIds = v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetOrderIds(v []*int64) *ListYunQiTaskByUidRequest {
	s.OrderIds = v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetPageNum(v int32) *ListYunQiTaskByUidRequest {
	s.PageNum = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetPageSize(v int32) *ListYunQiTaskByUidRequest {
	s.PageSize = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetStatusList(v []*string) *ListYunQiTaskByUidRequest {
	s.StatusList = v
	return s
}

func (s *ListYunQiTaskByUidRequest) Validate() error {
	return dara.Validate(s)
}
