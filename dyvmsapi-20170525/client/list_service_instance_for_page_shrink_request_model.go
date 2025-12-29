// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceForPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServiceInstanceForPageShrinkRequest
	GetCode() *string
	SetOwnerId(v int64) *ListServiceInstanceForPageShrinkRequest
	GetOwnerId() *int64
	SetPagerShrink(v string) *ListServiceInstanceForPageShrinkRequest
	GetPagerShrink() *string
	SetRelationNumber(v string) *ListServiceInstanceForPageShrinkRequest
	GetRelationNumber() *string
	SetResourceOwnerAccount(v string) *ListServiceInstanceForPageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListServiceInstanceForPageShrinkRequest
	GetResourceOwnerId() *int64
	SetSceneId(v int64) *ListServiceInstanceForPageShrinkRequest
	GetSceneId() *int64
	SetUsageId(v int64) *ListServiceInstanceForPageShrinkRequest
	GetUsageId() *int64
}

type ListServiceInstanceForPageShrinkRequest struct {
	// 服务实例号
	//
	// example:
	//
	// 0571****1234
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PagerShrink *string `json:"Pager,omitempty" xml:"Pager,omitempty"`
	// 关联实体号码
	//
	// example:
	//
	// 131****1111
	RelationNumber       *string `json:"RelationNumber,omitempty" xml:"RelationNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 场景ID
	//
	// example:
	//
	// 56
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 用途ID
	//
	// example:
	//
	// 89
	UsageId *int64 `json:"UsageId,omitempty" xml:"UsageId,omitempty"`
}

func (s ListServiceInstanceForPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *ListServiceInstanceForPageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListServiceInstanceForPageShrinkRequest) GetPagerShrink() *string {
	return s.PagerShrink
}

func (s *ListServiceInstanceForPageShrinkRequest) GetRelationNumber() *string {
	return s.RelationNumber
}

func (s *ListServiceInstanceForPageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListServiceInstanceForPageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListServiceInstanceForPageShrinkRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *ListServiceInstanceForPageShrinkRequest) GetUsageId() *int64 {
	return s.UsageId
}

func (s *ListServiceInstanceForPageShrinkRequest) SetCode(v string) *ListServiceInstanceForPageShrinkRequest {
	s.Code = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetOwnerId(v int64) *ListServiceInstanceForPageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetPagerShrink(v string) *ListServiceInstanceForPageShrinkRequest {
	s.PagerShrink = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetRelationNumber(v string) *ListServiceInstanceForPageShrinkRequest {
	s.RelationNumber = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetResourceOwnerAccount(v string) *ListServiceInstanceForPageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetResourceOwnerId(v int64) *ListServiceInstanceForPageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetSceneId(v int64) *ListServiceInstanceForPageShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) SetUsageId(v int64) *ListServiceInstanceForPageShrinkRequest {
	s.UsageId = &v
	return s
}

func (s *ListServiceInstanceForPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
