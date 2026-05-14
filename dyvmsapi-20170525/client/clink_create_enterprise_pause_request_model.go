// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateEnterprisePauseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkCreateEnterprisePauseRequest
	GetEnterpriseId() *int64
	SetIsDefault(v int64) *ClinkCreateEnterprisePauseRequest
	GetIsDefault() *int64
	SetIsRest(v string) *ClinkCreateEnterprisePauseRequest
	GetIsRest() *string
	SetOwnerId(v int64) *ClinkCreateEnterprisePauseRequest
	GetOwnerId() *int64
	SetPauseStatus(v string) *ClinkCreateEnterprisePauseRequest
	GetPauseStatus() *string
	SetResourceOwnerAccount(v string) *ClinkCreateEnterprisePauseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkCreateEnterprisePauseRequest
	GetResourceOwnerId() *int64
}

type ClinkCreateEnterprisePauseRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 默认状态，0：不是；1：是
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsDefault *int64 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 休息状态，0：不是；1：是
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsRest  *string `json:"IsRest,omitempty" xml:"IsRest,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 置忙状态描述（不超4个字符）
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	PauseStatus          *string `json:"PauseStatus,omitempty" xml:"PauseStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkCreateEnterprisePauseRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateEnterprisePauseRequest) GoString() string {
	return s.String()
}

func (s *ClinkCreateEnterprisePauseRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkCreateEnterprisePauseRequest) GetIsDefault() *int64 {
	return s.IsDefault
}

func (s *ClinkCreateEnterprisePauseRequest) GetIsRest() *string {
	return s.IsRest
}

func (s *ClinkCreateEnterprisePauseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkCreateEnterprisePauseRequest) GetPauseStatus() *string {
	return s.PauseStatus
}

func (s *ClinkCreateEnterprisePauseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkCreateEnterprisePauseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkCreateEnterprisePauseRequest) SetEnterpriseId(v int64) *ClinkCreateEnterprisePauseRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) SetIsDefault(v int64) *ClinkCreateEnterprisePauseRequest {
	s.IsDefault = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) SetIsRest(v string) *ClinkCreateEnterprisePauseRequest {
	s.IsRest = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) SetOwnerId(v int64) *ClinkCreateEnterprisePauseRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) SetPauseStatus(v string) *ClinkCreateEnterprisePauseRequest {
	s.PauseStatus = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) SetResourceOwnerAccount(v string) *ClinkCreateEnterprisePauseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) SetResourceOwnerId(v int64) *ClinkCreateEnterprisePauseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkCreateEnterprisePauseRequest) Validate() error {
	return dara.Validate(s)
}
