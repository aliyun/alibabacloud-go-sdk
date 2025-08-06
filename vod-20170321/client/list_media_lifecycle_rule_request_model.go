// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMediaLifecycleRuleRequest
	GetAppId() *string
	SetOwnerId(v int64) *ListMediaLifecycleRuleRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListMediaLifecycleRuleRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListMediaLifecycleRuleRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListMediaLifecycleRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMediaLifecycleRuleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *ListMediaLifecycleRuleRequest
	GetResourceRealOwnerId() *int64
	SetRuleType(v string) *ListMediaLifecycleRuleRequest
	GetRuleType() *string
}

type ListMediaLifecycleRuleRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListMediaLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *ListMediaLifecycleRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMediaLifecycleRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMediaLifecycleRuleRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListMediaLifecycleRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMediaLifecycleRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMediaLifecycleRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMediaLifecycleRuleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *ListMediaLifecycleRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListMediaLifecycleRuleRequest) SetAppId(v string) *ListMediaLifecycleRuleRequest {
	s.AppId = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetOwnerId(v int64) *ListMediaLifecycleRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetPageNo(v int32) *ListMediaLifecycleRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetPageSize(v int32) *ListMediaLifecycleRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetResourceOwnerAccount(v string) *ListMediaLifecycleRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetResourceOwnerId(v int64) *ListMediaLifecycleRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetResourceRealOwnerId(v int64) *ListMediaLifecycleRuleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) SetRuleType(v string) *ListMediaLifecycleRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ListMediaLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
