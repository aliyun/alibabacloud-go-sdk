// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactor(v string) *ListMemberRequest
	GetContactor() *string
	SetName(v string) *ListMemberRequest
	GetName() *string
	SetNum(v int32) *ListMemberRequest
	GetNum() *int32
	SetRegionId(v string) *ListMemberRequest
	GetRegionId() *string
	SetSize(v int32) *ListMemberRequest
	GetSize() *int32
	SetUid(v string) *ListMemberRequest
	GetUid() *string
}

type ListMemberRequest struct {
	Contactor *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Uid  *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemberRequest) GoString() string {
	return s.String()
}

func (s *ListMemberRequest) GetContactor() *string {
	return s.Contactor
}

func (s *ListMemberRequest) GetName() *string {
	return s.Name
}

func (s *ListMemberRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMemberRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListMemberRequest) GetUid() *string {
	return s.Uid
}

func (s *ListMemberRequest) SetContactor(v string) *ListMemberRequest {
	s.Contactor = &v
	return s
}

func (s *ListMemberRequest) SetName(v string) *ListMemberRequest {
	s.Name = &v
	return s
}

func (s *ListMemberRequest) SetNum(v int32) *ListMemberRequest {
	s.Num = &v
	return s
}

func (s *ListMemberRequest) SetRegionId(v string) *ListMemberRequest {
	s.RegionId = &v
	return s
}

func (s *ListMemberRequest) SetSize(v int32) *ListMemberRequest {
	s.Size = &v
	return s
}

func (s *ListMemberRequest) SetUid(v string) *ListMemberRequest {
	s.Uid = &v
	return s
}

func (s *ListMemberRequest) Validate() error {
	return dara.Validate(s)
}
