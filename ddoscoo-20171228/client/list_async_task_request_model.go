// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListAsyncTaskRequest
	GetLang() *string
	SetPageNo(v int32) *ListAsyncTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAsyncTaskRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListAsyncTaskRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *ListAsyncTaskRequest
	GetSourceIp() *string
}

type ListAsyncTaskRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAsyncTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAsyncTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAsyncTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAsyncTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListAsyncTaskRequest) SetLang(v string) *ListAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *ListAsyncTaskRequest) SetPageNo(v int32) *ListAsyncTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListAsyncTaskRequest) SetPageSize(v int32) *ListAsyncTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListAsyncTaskRequest) SetResourceGroupId(v string) *ListAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAsyncTaskRequest) SetSourceIp(v string) *ListAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *ListAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
