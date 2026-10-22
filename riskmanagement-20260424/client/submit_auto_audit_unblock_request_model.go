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
	// The list of assets.
	//
	// > The array supports up to 50 IP addresses.
	//
	// example:
	//
	// [{"domain":"","ip":"120.26.29.16","assetType":"EIP"}]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// The auditor.
	//
	// > You do not need to specify this parameter. It does not take effect.
	//
	// example:
	//
	// none
	Auditor *string `json:"Auditor,omitempty" xml:"Auditor,omitempty"`
	// The customer UID.
	//
	// example:
	//
	// 1317334647812936
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// The type of the email to send.
	//
	// example:
	//
	// xxx
	EmailType *string `json:"EmailType,omitempty" xml:"EmailType,omitempty"`
	// Specifies whether the email includes an unsubscribe option.
	//
	// - **true**: Included.
	//
	// - **false**: Not included.
	//
	// example:
	//
	// true
	HasUnsubscribe *bool `json:"HasUnsubscribe,omitempty" xml:"HasUnsubscribe,omitempty"`
	// The description of the number of emails to send.
	//
	// example:
	//
	// 10
	SendCount *string `json:"SendCount,omitempty" xml:"SendCount,omitempty"`
	// The description.
	//
	// example:
	//
	// Usage instructions
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
