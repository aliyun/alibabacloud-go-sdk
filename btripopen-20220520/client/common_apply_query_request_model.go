// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v int64) *CommonApplyQueryRequest
	GetApplyId() *int64
	SetBizCategory(v int32) *CommonApplyQueryRequest
	GetBizCategory() *int32
	SetBusinessInstanceId(v string) *CommonApplyQueryRequest
	GetBusinessInstanceId() *string
	SetUserId(v string) *CommonApplyQueryRequest
	GetUserId() *string
}

type CommonApplyQueryRequest struct {
	// example:
	//
	// 1003366164
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	BizCategory        *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	BusinessInstanceId *string `json:"business_instance_id,omitempty" xml:"business_instance_id,omitempty"`
	UserId             *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CommonApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CommonApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *CommonApplyQueryRequest) GetBizCategory() *int32 {
	return s.BizCategory
}

func (s *CommonApplyQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *CommonApplyQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *CommonApplyQueryRequest) SetApplyId(v int64) *CommonApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *CommonApplyQueryRequest) SetBizCategory(v int32) *CommonApplyQueryRequest {
	s.BizCategory = &v
	return s
}

func (s *CommonApplyQueryRequest) SetBusinessInstanceId(v string) *CommonApplyQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *CommonApplyQueryRequest) SetUserId(v string) *CommonApplyQueryRequest {
	s.UserId = &v
	return s
}

func (s *CommonApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
