// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAutoAuditUnblockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v string) *SubmitAutoAuditUnblockRequest
	GetAssetList() *string
	SetAuditor(v string) *SubmitAutoAuditUnblockRequest
	GetAuditor() *string
	SetCustomerUid(v string) *SubmitAutoAuditUnblockRequest
	GetCustomerUid() *string
	SetEmailType(v string) *SubmitAutoAuditUnblockRequest
	GetEmailType() *string
	SetHasUnsubscribe(v bool) *SubmitAutoAuditUnblockRequest
	GetHasUnsubscribe() *bool
	SetSendCount(v string) *SubmitAutoAuditUnblockRequest
	GetSendCount() *string
	SetUseDesc(v string) *SubmitAutoAuditUnblockRequest
	GetUseDesc() *string
}

type SubmitAutoAuditUnblockRequest struct {
	// example:
	//
	// [{"domain":"","ip":"120.26.29.16","assetType":"EIP"}]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// example:
	//
	// none
	Auditor *string `json:"Auditor,omitempty" xml:"Auditor,omitempty"`
	// example:
	//
	// 1317334647812936
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// xxx
	EmailType *string `json:"EmailType,omitempty" xml:"EmailType,omitempty"`
	// example:
	//
	// true
	HasUnsubscribe *bool `json:"HasUnsubscribe,omitempty" xml:"HasUnsubscribe,omitempty"`
	// example:
	//
	// 10
	SendCount *string `json:"SendCount,omitempty" xml:"SendCount,omitempty"`
	// example:
	//
	// 使用说明。
	UseDesc *string `json:"UseDesc,omitempty" xml:"UseDesc,omitempty"`
}

func (s SubmitAutoAuditUnblockRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAutoAuditUnblockRequest) GoString() string {
	return s.String()
}

func (s *SubmitAutoAuditUnblockRequest) GetAssetList() *string {
	return s.AssetList
}

func (s *SubmitAutoAuditUnblockRequest) GetAuditor() *string {
	return s.Auditor
}

func (s *SubmitAutoAuditUnblockRequest) GetCustomerUid() *string {
	return s.CustomerUid
}

func (s *SubmitAutoAuditUnblockRequest) GetEmailType() *string {
	return s.EmailType
}

func (s *SubmitAutoAuditUnblockRequest) GetHasUnsubscribe() *bool {
	return s.HasUnsubscribe
}

func (s *SubmitAutoAuditUnblockRequest) GetSendCount() *string {
	return s.SendCount
}

func (s *SubmitAutoAuditUnblockRequest) GetUseDesc() *string {
	return s.UseDesc
}

func (s *SubmitAutoAuditUnblockRequest) SetAssetList(v string) *SubmitAutoAuditUnblockRequest {
	s.AssetList = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) SetAuditor(v string) *SubmitAutoAuditUnblockRequest {
	s.Auditor = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) SetCustomerUid(v string) *SubmitAutoAuditUnblockRequest {
	s.CustomerUid = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) SetEmailType(v string) *SubmitAutoAuditUnblockRequest {
	s.EmailType = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) SetHasUnsubscribe(v bool) *SubmitAutoAuditUnblockRequest {
	s.HasUnsubscribe = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) SetSendCount(v string) *SubmitAutoAuditUnblockRequest {
	s.SendCount = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) SetUseDesc(v string) *SubmitAutoAuditUnblockRequest {
	s.UseDesc = &v
	return s
}

func (s *SubmitAutoAuditUnblockRequest) Validate() error {
	return dara.Validate(s)
}
