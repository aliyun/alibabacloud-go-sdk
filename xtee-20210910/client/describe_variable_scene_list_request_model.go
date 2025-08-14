// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableSceneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableSceneListRequest
	GetLang() *string
	SetBizType(v string) *DescribeVariableSceneListRequest
	GetBizType() *string
	SetConfigKey(v string) *DescribeVariableSceneListRequest
	GetConfigKey() *string
	SetCurrentPage(v string) *DescribeVariableSceneListRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeVariableSceneListRequest
	GetPageSize() *string
	SetPaging(v bool) *DescribeVariableSceneListRequest
	GetPaging() *bool
	SetRegId(v string) *DescribeVariableSceneListRequest
	GetRegId() *string
}

type DescribeVariableSceneListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// variable_scene
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// Configuration key.
	//
	// example:
	//
	// account_abuse_detection
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Paging flag, default is true.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Paging *bool `json:"paging,omitempty" xml:"paging,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeVariableSceneListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableSceneListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableSceneListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableSceneListRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeVariableSceneListRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeVariableSceneListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVariableSceneListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVariableSceneListRequest) GetPaging() *bool {
	return s.Paging
}

func (s *DescribeVariableSceneListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableSceneListRequest) SetLang(v string) *DescribeVariableSceneListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableSceneListRequest) SetBizType(v string) *DescribeVariableSceneListRequest {
	s.BizType = &v
	return s
}

func (s *DescribeVariableSceneListRequest) SetConfigKey(v string) *DescribeVariableSceneListRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeVariableSceneListRequest) SetCurrentPage(v string) *DescribeVariableSceneListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVariableSceneListRequest) SetPageSize(v string) *DescribeVariableSceneListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVariableSceneListRequest) SetPaging(v bool) *DescribeVariableSceneListRequest {
	s.Paging = &v
	return s
}

func (s *DescribeVariableSceneListRequest) SetRegId(v string) *DescribeVariableSceneListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableSceneListRequest) Validate() error {
	return dara.Validate(s)
}
