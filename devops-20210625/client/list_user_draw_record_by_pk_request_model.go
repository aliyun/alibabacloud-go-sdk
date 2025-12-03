// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDrawRecordByPkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunPk(v string) *ListUserDrawRecordByPkRequest
	GetAliyunPk() *string
	SetDrawGroup(v string) *ListUserDrawRecordByPkRequest
	GetDrawGroup() *string
	SetDrawPoolName(v string) *ListUserDrawRecordByPkRequest
	GetDrawPoolName() *string
}

type ListUserDrawRecordByPkRequest struct {
	// example:
	//
	// 1789095186553536
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// example:
	//
	// task_group_0000
	DrawGroup *string `json:"drawGroup,omitempty" xml:"drawGroup,omitempty"`
	// example:
	//
	// developer-award-draw-pool_123
	DrawPoolName *string `json:"drawPoolName,omitempty" xml:"drawPoolName,omitempty"`
}

func (s ListUserDrawRecordByPkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDrawRecordByPkRequest) GoString() string {
	return s.String()
}

func (s *ListUserDrawRecordByPkRequest) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListUserDrawRecordByPkRequest) GetDrawGroup() *string {
	return s.DrawGroup
}

func (s *ListUserDrawRecordByPkRequest) GetDrawPoolName() *string {
	return s.DrawPoolName
}

func (s *ListUserDrawRecordByPkRequest) SetAliyunPk(v string) *ListUserDrawRecordByPkRequest {
	s.AliyunPk = &v
	return s
}

func (s *ListUserDrawRecordByPkRequest) SetDrawGroup(v string) *ListUserDrawRecordByPkRequest {
	s.DrawGroup = &v
	return s
}

func (s *ListUserDrawRecordByPkRequest) SetDrawPoolName(v string) *ListUserDrawRecordByPkRequest {
	s.DrawPoolName = &v
	return s
}

func (s *ListUserDrawRecordByPkRequest) Validate() error {
	return dara.Validate(s)
}
