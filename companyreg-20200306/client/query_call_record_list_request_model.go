// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallRecordListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryCallRecordListRequest
	GetBizId() *string
	SetBizType(v string) *QueryCallRecordListRequest
	GetBizType() *string
	SetSkillType(v int64) *QueryCallRecordListRequest
	GetSkillType() *int64
}

type QueryCallRecordListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P20210928095324000002
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.bookkeeping
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SkillType *int64 `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
}

func (s QueryCallRecordListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallRecordListRequest) GoString() string {
	return s.String()
}

func (s *QueryCallRecordListRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryCallRecordListRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryCallRecordListRequest) GetSkillType() *int64 {
	return s.SkillType
}

func (s *QueryCallRecordListRequest) SetBizId(v string) *QueryCallRecordListRequest {
	s.BizId = &v
	return s
}

func (s *QueryCallRecordListRequest) SetBizType(v string) *QueryCallRecordListRequest {
	s.BizType = &v
	return s
}

func (s *QueryCallRecordListRequest) SetSkillType(v int64) *QueryCallRecordListRequest {
	s.SkillType = &v
	return s
}

func (s *QueryCallRecordListRequest) Validate() error {
	return dara.Validate(s)
}
