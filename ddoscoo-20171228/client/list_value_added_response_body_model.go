// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListValueAddedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListValueAddedResponseBody
	GetRequestId() *string
	SetValueAddedList(v []*ListValueAddedResponseBodyValueAddedList) *ListValueAddedResponseBody
	GetValueAddedList() []*ListValueAddedResponseBodyValueAddedList
}

type ListValueAddedResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ValueAddedList []*ListValueAddedResponseBodyValueAddedList `json:"ValueAddedList,omitempty" xml:"ValueAddedList,omitempty" type:"Repeated"`
}

func (s ListValueAddedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListValueAddedResponseBody) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListValueAddedResponseBody) GetValueAddedList() []*ListValueAddedResponseBodyValueAddedList {
	return s.ValueAddedList
}

func (s *ListValueAddedResponseBody) SetRequestId(v string) *ListValueAddedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListValueAddedResponseBody) SetValueAddedList(v []*ListValueAddedResponseBodyValueAddedList) *ListValueAddedResponseBody {
	s.ValueAddedList = v
	return s
}

func (s *ListValueAddedResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListValueAddedResponseBodyValueAddedList struct {
	// example:
	//
	// 1580918400000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1575527305000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// ddos_fl_pre-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5497558138880
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// example:
	//
	// 1
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StoreRegion *string `json:"StoreRegion,omitempty" xml:"StoreRegion,omitempty"`
}

func (s ListValueAddedResponseBodyValueAddedList) String() string {
	return dara.Prettify(s)
}

func (s ListValueAddedResponseBodyValueAddedList) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponseBodyValueAddedList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListValueAddedResponseBodyValueAddedList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListValueAddedResponseBodyValueAddedList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListValueAddedResponseBodyValueAddedList) GetLogSize() *int64 {
	return s.LogSize
}

func (s *ListValueAddedResponseBodyValueAddedList) GetStatus() *int32 {
	return s.Status
}

func (s *ListValueAddedResponseBodyValueAddedList) GetStoreRegion() *string {
	return s.StoreRegion
}

func (s *ListValueAddedResponseBodyValueAddedList) SetExpireTime(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.ExpireTime = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetGmtCreate(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.GmtCreate = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetInstanceId(v string) *ListValueAddedResponseBodyValueAddedList {
	s.InstanceId = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetLogSize(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.LogSize = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetStatus(v int32) *ListValueAddedResponseBodyValueAddedList {
	s.Status = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetStoreRegion(v string) *ListValueAddedResponseBodyValueAddedList {
	s.StoreRegion = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) Validate() error {
	return dara.Validate(s)
}
