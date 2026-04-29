// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v string) *CloudGetAsrRequest
	GetAll() *string
	SetCallType(v int64) *CloudGetAsrRequest
	GetCallType() *int64
	SetEnterpriseId(v int64) *CloudGetAsrRequest
	GetEnterpriseId() *int64
	SetMainUniqueId(v string) *CloudGetAsrRequest
	GetMainUniqueId() *string
}

type CloudGetAsrRequest struct {
	// 当all=true时按照beignTime的顺序返回两侧对话文本
	//
	// example:
	//
	// true
	All *string `json:"All,omitempty" xml:"All,omitempty"`
	// 通话类型；1:呼入,2:webcall,4:预览外呼,5:预测外呼,6:主叫外呼,9:内部呼叫
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 主通道的唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// sip-2-1490862368.123
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
}

func (s CloudGetAsrRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrRequest) GoString() string {
	return s.String()
}

func (s *CloudGetAsrRequest) GetAll() *string {
	return s.All
}

func (s *CloudGetAsrRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *CloudGetAsrRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetAsrRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *CloudGetAsrRequest) SetAll(v string) *CloudGetAsrRequest {
	s.All = &v
	return s
}

func (s *CloudGetAsrRequest) SetCallType(v int64) *CloudGetAsrRequest {
	s.CallType = &v
	return s
}

func (s *CloudGetAsrRequest) SetEnterpriseId(v int64) *CloudGetAsrRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetAsrRequest) SetMainUniqueId(v string) *CloudGetAsrRequest {
	s.MainUniqueId = &v
	return s
}

func (s *CloudGetAsrRequest) Validate() error {
	return dara.Validate(s)
}
