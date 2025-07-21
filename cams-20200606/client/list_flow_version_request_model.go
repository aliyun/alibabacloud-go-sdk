// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *ListFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *ListFlowVersionRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *ListFlowVersionRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListFlowVersionRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListFlowVersionRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFlowVersionRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListFlowVersionRequest
	GetStatus() *string
}

type ListFlowVersionRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information, default is “{}”.
	//
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// Flow code.
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf****
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size.
	//
	// example:
	//
	// 5
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Flow version status.
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *ListFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ListFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ListFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFlowVersionRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListFlowVersionRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFlowVersionRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFlowVersionRequest) SetBizCode(v string) *ListFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *ListFlowVersionRequest) SetBizExtend(v map[string]interface{}) *ListFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *ListFlowVersionRequest) SetFlowCode(v string) *ListFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *ListFlowVersionRequest) SetOwnerId(v int64) *ListFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFlowVersionRequest) SetPageNo(v int64) *ListFlowVersionRequest {
	s.PageNo = &v
	return s
}

func (s *ListFlowVersionRequest) SetPageSize(v int64) *ListFlowVersionRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowVersionRequest) SetResourceOwnerAccount(v string) *ListFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFlowVersionRequest) SetResourceOwnerId(v int64) *ListFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowVersionRequest) SetStatus(v string) *ListFlowVersionRequest {
	s.Status = &v
	return s
}

func (s *ListFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
