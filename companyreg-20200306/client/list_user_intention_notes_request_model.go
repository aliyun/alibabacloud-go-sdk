// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserIntentionNotesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ListUserIntentionNotesRequest
	GetBizType() *string
	SetIntentionBizId(v string) *ListUserIntentionNotesRequest
	GetIntentionBizId() *string
	SetPageNum(v int32) *ListUserIntentionNotesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserIntentionNotesRequest
	GetPageSize() *int32
}

type ListUserIntentionNotesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// esp.beian_assist
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// I20210912102942000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserIntentionNotesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionNotesRequest) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListUserIntentionNotesRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserIntentionNotesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserIntentionNotesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserIntentionNotesRequest) SetBizType(v string) *ListUserIntentionNotesRequest {
	s.BizType = &v
	return s
}

func (s *ListUserIntentionNotesRequest) SetIntentionBizId(v string) *ListUserIntentionNotesRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserIntentionNotesRequest) SetPageNum(v int32) *ListUserIntentionNotesRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserIntentionNotesRequest) SetPageSize(v int32) *ListUserIntentionNotesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionNotesRequest) Validate() error {
	return dara.Validate(s)
}
