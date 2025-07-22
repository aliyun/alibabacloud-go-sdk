// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeFaultDiagnosisUserListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeFaultDiagnosisUserListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFaultDiagnosisUserListResponseBody
	GetRequestId() *string
	SetTotalCnt(v int32) *DescribeFaultDiagnosisUserListResponseBody
	GetTotalCnt() *int32
	SetUserList(v []*DescribeFaultDiagnosisUserListResponseBodyUserList) *DescribeFaultDiagnosisUserListResponseBody
	GetUserList() []*DescribeFaultDiagnosisUserListResponseBodyUserList
}

type DescribeFaultDiagnosisUserListResponseBody struct {
	// example:
	//
	// 2
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCnt *int32                                                `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	UserList []*DescribeFaultDiagnosisUserListResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeFaultDiagnosisUserListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFaultDiagnosisUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaultDiagnosisUserListResponseBody) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribeFaultDiagnosisUserListResponseBody) GetUserList() []*DescribeFaultDiagnosisUserListResponseBodyUserList {
	return s.UserList
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetPageNo(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetPageSize(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetTotalCnt(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetUserList(v []*DescribeFaultDiagnosisUserListResponseBodyUserList) *DescribeFaultDiagnosisUserListResponseBody {
	s.UserList = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisUserListResponseBodyUserList struct {
	// example:
	//
	// 1614936817
	ChannelCreatedTs *int64 `json:"ChannelCreatedTs,omitempty" xml:"ChannelCreatedTs,omitempty"`
	// example:
	//
	// 904
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64                                                         `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	FaultList   []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList `json:"FaultList,omitempty" xml:"FaultList,omitempty" type:"Repeated"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) GetChannelCreatedTs() *int64 {
	return s.ChannelCreatedTs
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) GetFaultList() []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList {
	return s.FaultList
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) GetUserId() *string {
	return s.UserId
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetChannelCreatedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.ChannelCreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetChannelId(v string) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetFaultList(v []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.FaultList = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetUserId(v string) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.UserId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisUserListResponseBodyUserListFaultList struct {
	// example:
	//
	// JOIN_SLOW
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) GetFaultType() *string {
	return s.FaultType
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) SetFaultType(v string) *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList {
	s.FaultType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) Validate() error {
	return dara.Validate(s)
}
