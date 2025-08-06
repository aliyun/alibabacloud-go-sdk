// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProduceOperateLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListUserProduceOperateLogsRequest
	GetBizId() *string
	SetBizType(v string) *ListUserProduceOperateLogsRequest
	GetBizType() *string
	SetPageNum(v int32) *ListUserProduceOperateLogsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserProduceOperateLogsRequest
	GetPageSize() *int32
}

type ListUserProduceOperateLogsRequest struct {
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
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
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

func (s ListUserProduceOperateLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserProduceOperateLogsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListUserProduceOperateLogsRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListUserProduceOperateLogsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserProduceOperateLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserProduceOperateLogsRequest) SetBizId(v string) *ListUserProduceOperateLogsRequest {
	s.BizId = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) SetBizType(v string) *ListUserProduceOperateLogsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) SetPageNum(v int32) *ListUserProduceOperateLogsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) SetPageSize(v int32) *ListUserProduceOperateLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) Validate() error {
	return dara.Validate(s)
}
