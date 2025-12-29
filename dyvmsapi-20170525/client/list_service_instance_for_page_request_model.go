// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceForPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServiceInstanceForPageRequest
	GetCode() *string
	SetOwnerId(v int64) *ListServiceInstanceForPageRequest
	GetOwnerId() *int64
	SetPager(v *ListServiceInstanceForPageRequestPager) *ListServiceInstanceForPageRequest
	GetPager() *ListServiceInstanceForPageRequestPager
	SetRelationNumber(v string) *ListServiceInstanceForPageRequest
	GetRelationNumber() *string
	SetResourceOwnerAccount(v string) *ListServiceInstanceForPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListServiceInstanceForPageRequest
	GetResourceOwnerId() *int64
	SetSceneId(v int64) *ListServiceInstanceForPageRequest
	GetSceneId() *int64
	SetUsageId(v int64) *ListServiceInstanceForPageRequest
	GetUsageId() *int64
}

type ListServiceInstanceForPageRequest struct {
	// 服务实例号
	//
	// example:
	//
	// 0571****1234
	Code    *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	OwnerId *int64                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Pager   *ListServiceInstanceForPageRequestPager `json:"Pager,omitempty" xml:"Pager,omitempty" type:"Struct"`
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

func (s ListServiceInstanceForPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageRequest) GetCode() *string {
	return s.Code
}

func (s *ListServiceInstanceForPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListServiceInstanceForPageRequest) GetPager() *ListServiceInstanceForPageRequestPager {
	return s.Pager
}

func (s *ListServiceInstanceForPageRequest) GetRelationNumber() *string {
	return s.RelationNumber
}

func (s *ListServiceInstanceForPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListServiceInstanceForPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListServiceInstanceForPageRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *ListServiceInstanceForPageRequest) GetUsageId() *int64 {
	return s.UsageId
}

func (s *ListServiceInstanceForPageRequest) SetCode(v string) *ListServiceInstanceForPageRequest {
	s.Code = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetOwnerId(v int64) *ListServiceInstanceForPageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetPager(v *ListServiceInstanceForPageRequestPager) *ListServiceInstanceForPageRequest {
	s.Pager = v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetRelationNumber(v string) *ListServiceInstanceForPageRequest {
	s.RelationNumber = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetResourceOwnerAccount(v string) *ListServiceInstanceForPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetResourceOwnerId(v int64) *ListServiceInstanceForPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetSceneId(v int64) *ListServiceInstanceForPageRequest {
	s.SceneId = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) SetUsageId(v int64) *ListServiceInstanceForPageRequest {
	s.UsageId = &v
	return s
}

func (s *ListServiceInstanceForPageRequest) Validate() error {
	if s.Pager != nil {
		if err := s.Pager.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceInstanceForPageRequestPager struct {
	// 当前页码
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// 每页数量，默认10，最大100
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServiceInstanceForPageRequestPager) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageRequestPager) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageRequestPager) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListServiceInstanceForPageRequestPager) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListServiceInstanceForPageRequestPager) SetPageIndex(v int64) *ListServiceInstanceForPageRequestPager {
	s.PageIndex = &v
	return s
}

func (s *ListServiceInstanceForPageRequestPager) SetPageSize(v int64) *ListServiceInstanceForPageRequestPager {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstanceForPageRequestPager) Validate() error {
	return dara.Validate(s)
}
