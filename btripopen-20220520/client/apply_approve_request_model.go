// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApproveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v string) *ApplyApproveRequest
	GetApplyId() *string
	SetNote(v string) *ApplyApproveRequest
	GetNote() *string
	SetOperateTime(v string) *ApplyApproveRequest
	GetOperateTime() *string
	SetStatus(v int32) *ApplyApproveRequest
	GetStatus() *int32
	SetSubCorpId(v string) *ApplyApproveRequest
	GetSubCorpId() *string
	SetUserId(v string) *ApplyApproveRequest
	GetUserId() *string
	SetUserName(v string) *ApplyApproveRequest
	GetUserName() *string
}

type ApplyApproveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sdfg
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	Note    *string `json:"note,omitempty" xml:"note,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-12 16:12:53
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// btrip123
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thirdpart12138
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyApproveRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyApproveRequest) GoString() string {
	return s.String()
}

func (s *ApplyApproveRequest) GetApplyId() *string {
	return s.ApplyId
}

func (s *ApplyApproveRequest) GetNote() *string {
	return s.Note
}

func (s *ApplyApproveRequest) GetOperateTime() *string {
	return s.OperateTime
}

func (s *ApplyApproveRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyApproveRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyApproveRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyApproveRequest) GetUserName() *string {
	return s.UserName
}

func (s *ApplyApproveRequest) SetApplyId(v string) *ApplyApproveRequest {
	s.ApplyId = &v
	return s
}

func (s *ApplyApproveRequest) SetNote(v string) *ApplyApproveRequest {
	s.Note = &v
	return s
}

func (s *ApplyApproveRequest) SetOperateTime(v string) *ApplyApproveRequest {
	s.OperateTime = &v
	return s
}

func (s *ApplyApproveRequest) SetStatus(v int32) *ApplyApproveRequest {
	s.Status = &v
	return s
}

func (s *ApplyApproveRequest) SetSubCorpId(v string) *ApplyApproveRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyApproveRequest) SetUserId(v string) *ApplyApproveRequest {
	s.UserId = &v
	return s
}

func (s *ApplyApproveRequest) SetUserName(v string) *ApplyApproveRequest {
	s.UserName = &v
	return s
}

func (s *ApplyApproveRequest) Validate() error {
	return dara.Validate(s)
}
