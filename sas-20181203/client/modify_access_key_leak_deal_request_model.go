// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessKeyLeakDealRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyAccessKeyLeakDealRequest
	GetId() *int64
	SetIdList(v []*int64) *ModifyAccessKeyLeakDealRequest
	GetIdList() []*int64
	SetRemark(v string) *ModifyAccessKeyLeakDealRequest
	GetRemark() *string
	SetType(v string) *ModifyAccessKeyLeakDealRequest
	GetType() *string
}

type ModifyAccessKeyLeakDealRequest struct {
	// The ID of the AccessKey pair leak record.
	//
	// >Call the [DescribeAccesskeyLeakList](~~DescribeAccesskeyLeakList~~) operation to obtain this parameter. This parameter and the **IdList*	- parameter cannot both be empty.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of the AccessKey pair leak records.
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// The remarks for handling the AccessKey pair leak record.
	//
	// example:
	//
	// disabled.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The method to handle the AccessKey pair leak information. Valid values:
	//
	// - **manual**: Manual handling.
	//
	// - **disable**: Disable.
	//
	// - **add-whitelist**: Add to whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAccessKeyLeakDealRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessKeyLeakDealRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessKeyLeakDealRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyAccessKeyLeakDealRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *ModifyAccessKeyLeakDealRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyAccessKeyLeakDealRequest) GetType() *string {
	return s.Type
}

func (s *ModifyAccessKeyLeakDealRequest) SetId(v int64) *ModifyAccessKeyLeakDealRequest {
	s.Id = &v
	return s
}

func (s *ModifyAccessKeyLeakDealRequest) SetIdList(v []*int64) *ModifyAccessKeyLeakDealRequest {
	s.IdList = v
	return s
}

func (s *ModifyAccessKeyLeakDealRequest) SetRemark(v string) *ModifyAccessKeyLeakDealRequest {
	s.Remark = &v
	return s
}

func (s *ModifyAccessKeyLeakDealRequest) SetType(v string) *ModifyAccessKeyLeakDealRequest {
	s.Type = &v
	return s
}

func (s *ModifyAccessKeyLeakDealRequest) Validate() error {
	return dara.Validate(s)
}
