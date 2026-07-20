// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v int32) *ApplyQueryRequest
	GetApplyId() *int32
	SetApplyShowId(v string) *ApplyQueryRequest
	GetApplyShowId() *string
	SetSubCorpId(v string) *ApplyQueryRequest
	GetSubCorpId() *string
	SetThirdpartApplyId(v string) *ApplyQueryRequest
	GetThirdpartApplyId() *string
	SetType(v int32) *ApplyQueryRequest
	GetType() *int32
}

type ApplyQueryRequest struct {
	// example:
	//
	// 123
	ApplyId *int32 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 201710111505000464651
	ApplyShowId *string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	// example:
	//
	// btrip123
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// example:
	//
	// adczd
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *ApplyQueryRequest) GetApplyId() *int32 {
	return s.ApplyId
}

func (s *ApplyQueryRequest) GetApplyShowId() *string {
	return s.ApplyShowId
}

func (s *ApplyQueryRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyQueryRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyQueryRequest) GetType() *int32 {
	return s.Type
}

func (s *ApplyQueryRequest) SetApplyId(v int32) *ApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *ApplyQueryRequest) SetApplyShowId(v string) *ApplyQueryRequest {
	s.ApplyShowId = &v
	return s
}

func (s *ApplyQueryRequest) SetSubCorpId(v string) *ApplyQueryRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyQueryRequest) SetThirdpartApplyId(v string) *ApplyQueryRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyQueryRequest) SetType(v int32) *ApplyQueryRequest {
	s.Type = &v
	return s
}

func (s *ApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
