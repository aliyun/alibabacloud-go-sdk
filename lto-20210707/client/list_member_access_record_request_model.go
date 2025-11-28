// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberAccessRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessStatus(v string) *ListMemberAccessRecordRequest
	GetAccessStatus() *string
	SetContactor(v string) *ListMemberAccessRecordRequest
	GetContactor() *string
	SetName(v string) *ListMemberAccessRecordRequest
	GetName() *string
	SetNum(v int64) *ListMemberAccessRecordRequest
	GetNum() *int64
	SetRegionId(v string) *ListMemberAccessRecordRequest
	GetRegionId() *string
	SetSize(v int64) *ListMemberAccessRecordRequest
	GetSize() *int64
	SetUid(v string) *ListMemberAccessRecordRequest
	GetUid() *string
}

type ListMemberAccessRecordRequest struct {
	AccessStatus *string `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	Contactor    *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num      *int64  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Uid  *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListMemberAccessRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAccessRecordRequest) GoString() string {
	return s.String()
}

func (s *ListMemberAccessRecordRequest) GetAccessStatus() *string {
	return s.AccessStatus
}

func (s *ListMemberAccessRecordRequest) GetContactor() *string {
	return s.Contactor
}

func (s *ListMemberAccessRecordRequest) GetName() *string {
	return s.Name
}

func (s *ListMemberAccessRecordRequest) GetNum() *int64 {
	return s.Num
}

func (s *ListMemberAccessRecordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMemberAccessRecordRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListMemberAccessRecordRequest) GetUid() *string {
	return s.Uid
}

func (s *ListMemberAccessRecordRequest) SetAccessStatus(v string) *ListMemberAccessRecordRequest {
	s.AccessStatus = &v
	return s
}

func (s *ListMemberAccessRecordRequest) SetContactor(v string) *ListMemberAccessRecordRequest {
	s.Contactor = &v
	return s
}

func (s *ListMemberAccessRecordRequest) SetName(v string) *ListMemberAccessRecordRequest {
	s.Name = &v
	return s
}

func (s *ListMemberAccessRecordRequest) SetNum(v int64) *ListMemberAccessRecordRequest {
	s.Num = &v
	return s
}

func (s *ListMemberAccessRecordRequest) SetRegionId(v string) *ListMemberAccessRecordRequest {
	s.RegionId = &v
	return s
}

func (s *ListMemberAccessRecordRequest) SetSize(v int64) *ListMemberAccessRecordRequest {
	s.Size = &v
	return s
}

func (s *ListMemberAccessRecordRequest) SetUid(v string) *ListMemberAccessRecordRequest {
	s.Uid = &v
	return s
}

func (s *ListMemberAccessRecordRequest) Validate() error {
	return dara.Validate(s)
}
