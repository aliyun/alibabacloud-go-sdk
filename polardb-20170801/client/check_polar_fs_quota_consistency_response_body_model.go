// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckPolarFsQuotaConsistencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolarFsInstanceId(v string) *CheckPolarFsQuotaConsistencyResponseBody
	GetPolarFsInstanceId() *string
	SetQuotaItem(v *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) *CheckPolarFsQuotaConsistencyResponseBody
	GetQuotaItem() *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem
	SetRequestId(v string) *CheckPolarFsQuotaConsistencyResponseBody
	GetRequestId() *string
}

type CheckPolarFsQuotaConsistencyResponseBody struct {
	// The ID of the PolarFS instance.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// An object that contains details about the quota.
	QuotaItem *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem `json:"QuotaItem,omitempty" xml:"QuotaItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckPolarFsQuotaConsistencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarFsQuotaConsistencyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) GetQuotaItem() *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem {
	return s.QuotaItem
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) SetPolarFsInstanceId(v string) *CheckPolarFsQuotaConsistencyResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) SetQuotaItem(v *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) *CheckPolarFsQuotaConsistencyResponseBody {
	s.QuotaItem = v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) SetRequestId(v string) *CheckPolarFsQuotaConsistencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBody) Validate() error {
	if s.QuotaItem != nil {
		if err := s.QuotaItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckPolarFsQuotaConsistencyResponseBodyQuotaItem struct {
	// The storage capacity quota.
	//
	// example:
	//
	// 104857600
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The inode quota.
	//
	// example:
	//
	// 50000
	Inodes *int64 `json:"Inodes,omitempty" xml:"Inodes,omitempty"`
	// The directory path.
	//
	// example:
	//
	// /zookeepertest718
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The used capacity, in bytes.
	//
	// example:
	//
	// 104857600
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	// The number of used inodes.
	//
	// example:
	//
	// 1
	UsedInodes *int64 `json:"UsedInodes,omitempty" xml:"UsedInodes,omitempty"`
}

func (s CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) GoString() string {
	return s.String()
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) GetInodes() *int64 {
	return s.Inodes
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) GetPath() *string {
	return s.Path
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) GetUsedInodes() *int64 {
	return s.UsedInodes
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) SetCapacity(v int64) *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem {
	s.Capacity = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) SetInodes(v int64) *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem {
	s.Inodes = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) SetPath(v string) *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem {
	s.Path = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) SetUsedCapacity(v int64) *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem {
	s.UsedCapacity = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) SetUsedInodes(v int64) *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem {
	s.UsedInodes = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyResponseBodyQuotaItem) Validate() error {
	return dara.Validate(s)
}
