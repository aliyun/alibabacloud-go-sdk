// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListVpcResponseBody
	GetCode() *int32
	SetMessage(v string) *ListVpcResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVpcResponseBody
	GetRequestId() *string
	SetVpcList(v *ListVpcResponseBodyVpcList) *ListVpcResponseBody
	GetVpcList() *ListVpcResponseBodyVpcList
}

type ListVpcResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about VPCs.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcList   *ListVpcResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Struct"`
}

func (s ListVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVpcResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcResponseBody) GetVpcList() *ListVpcResponseBodyVpcList {
	return s.VpcList
}

func (s *ListVpcResponseBody) SetCode(v int32) *ListVpcResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpcResponseBody) SetMessage(v string) *ListVpcResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpcResponseBody) SetRequestId(v string) *ListVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcResponseBody) SetVpcList(v *ListVpcResponseBodyVpcList) *ListVpcResponseBody {
	s.VpcList = v
	return s
}

func (s *ListVpcResponseBody) Validate() error {
	if s.VpcList != nil {
		if err := s.VpcList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVpcResponseBodyVpcList struct {
	VpcEntity []*ListVpcResponseBodyVpcListVpcEntity `json:"VpcEntity,omitempty" xml:"VpcEntity,omitempty" type:"Repeated"`
}

func (s ListVpcResponseBodyVpcList) String() string {
	return dara.Prettify(s)
}

func (s ListVpcResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *ListVpcResponseBodyVpcList) GetVpcEntity() []*ListVpcResponseBodyVpcListVpcEntity {
	return s.VpcEntity
}

func (s *ListVpcResponseBodyVpcList) SetVpcEntity(v []*ListVpcResponseBodyVpcListVpcEntity) *ListVpcResponseBodyVpcList {
	s.VpcEntity = v
	return s
}

func (s *ListVpcResponseBodyVpcList) Validate() error {
	if s.VpcEntity != nil {
		for _, item := range s.VpcEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcResponseBodyVpcListVpcEntity struct {
	EcsNum   *int32  `json:"EcsNum,omitempty" xml:"EcsNum,omitempty"`
	Expired  *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName  *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ListVpcResponseBodyVpcListVpcEntity) String() string {
	return dara.Prettify(s)
}

func (s ListVpcResponseBodyVpcListVpcEntity) GoString() string {
	return s.String()
}

func (s *ListVpcResponseBodyVpcListVpcEntity) GetEcsNum() *int32 {
	return s.EcsNum
}

func (s *ListVpcResponseBodyVpcListVpcEntity) GetExpired() *bool {
	return s.Expired
}

func (s *ListVpcResponseBodyVpcListVpcEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcResponseBodyVpcListVpcEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListVpcResponseBodyVpcListVpcEntity) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcResponseBodyVpcListVpcEntity) GetVpcName() *string {
	return s.VpcName
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetEcsNum(v int32) *ListVpcResponseBodyVpcListVpcEntity {
	s.EcsNum = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetExpired(v bool) *ListVpcResponseBodyVpcListVpcEntity {
	s.Expired = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetRegionId(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.RegionId = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetUserId(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.UserId = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetVpcId(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.VpcId = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetVpcName(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.VpcName = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) Validate() error {
	return dara.Validate(s)
}
