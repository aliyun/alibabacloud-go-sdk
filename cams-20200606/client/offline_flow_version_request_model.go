// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *OfflineFlowVersionRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *OfflineFlowVersionRequest
	GetBizExtend() map[string]interface{}
	SetFlowCode(v string) *OfflineFlowVersionRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *OfflineFlowVersionRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *OfflineFlowVersionRequest
	GetOwnerId() *int64
	SetRemark(v string) *OfflineFlowVersionRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *OfflineFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OfflineFlowVersionRequest
	GetResourceOwnerId() *int64
}

type OfflineFlowVersionRequest struct {
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
	// Flow version
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Flow remarks
	//
	// example:
	//
	// We don\\"t need this old version.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OfflineFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *OfflineFlowVersionRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *OfflineFlowVersionRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *OfflineFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *OfflineFlowVersionRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *OfflineFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OfflineFlowVersionRequest) GetRemark() *string {
	return s.Remark
}

func (s *OfflineFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OfflineFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OfflineFlowVersionRequest) SetBizCode(v string) *OfflineFlowVersionRequest {
	s.BizCode = &v
	return s
}

func (s *OfflineFlowVersionRequest) SetBizExtend(v map[string]interface{}) *OfflineFlowVersionRequest {
	s.BizExtend = v
	return s
}

func (s *OfflineFlowVersionRequest) SetFlowCode(v string) *OfflineFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *OfflineFlowVersionRequest) SetFlowVersion(v string) *OfflineFlowVersionRequest {
	s.FlowVersion = &v
	return s
}

func (s *OfflineFlowVersionRequest) SetOwnerId(v int64) *OfflineFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *OfflineFlowVersionRequest) SetRemark(v string) *OfflineFlowVersionRequest {
	s.Remark = &v
	return s
}

func (s *OfflineFlowVersionRequest) SetResourceOwnerAccount(v string) *OfflineFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OfflineFlowVersionRequest) SetResourceOwnerId(v int64) *OfflineFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OfflineFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
