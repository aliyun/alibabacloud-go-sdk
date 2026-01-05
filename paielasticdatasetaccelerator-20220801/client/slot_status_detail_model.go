// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlotStatusDetail interface {
	dara.Model
	String() string
	GoString() string
	SetLoadedFileNum(v int64) *SlotStatusDetail
	GetLoadedFileNum() *int64
	SetLoadedFileSize(v string) *SlotStatusDetail
	GetLoadedFileSize() *string
	SetLoadingTimeCost(v int64) *SlotStatusDetail
	GetLoadingTimeCost() *int64
}

type SlotStatusDetail struct {
	// example:
	//
	// 1000000
	LoadedFileNum *int64 `json:"LoadedFileNum,omitempty" xml:"LoadedFileNum,omitempty"`
	// example:
	//
	// 200G
	LoadedFileSize *string `json:"LoadedFileSize,omitempty" xml:"LoadedFileSize,omitempty"`
	// example:
	//
	// 1800000
	LoadingTimeCost *int64 `json:"LoadingTimeCost,omitempty" xml:"LoadingTimeCost,omitempty"`
}

func (s SlotStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s SlotStatusDetail) GoString() string {
	return s.String()
}

func (s *SlotStatusDetail) GetLoadedFileNum() *int64 {
	return s.LoadedFileNum
}

func (s *SlotStatusDetail) GetLoadedFileSize() *string {
	return s.LoadedFileSize
}

func (s *SlotStatusDetail) GetLoadingTimeCost() *int64 {
	return s.LoadingTimeCost
}

func (s *SlotStatusDetail) SetLoadedFileNum(v int64) *SlotStatusDetail {
	s.LoadedFileNum = &v
	return s
}

func (s *SlotStatusDetail) SetLoadedFileSize(v string) *SlotStatusDetail {
	s.LoadedFileSize = &v
	return s
}

func (s *SlotStatusDetail) SetLoadingTimeCost(v int64) *SlotStatusDetail {
	s.LoadingTimeCost = &v
	return s
}

func (s *SlotStatusDetail) Validate() error {
	return dara.Validate(s)
}
