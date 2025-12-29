// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagInfoBySelectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndustryId(v int64) *QueryTagInfoBySelectionRequest
	GetIndustryId() *int64
	SetOwnerId(v int64) *QueryTagInfoBySelectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryTagInfoBySelectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTagInfoBySelectionRequest
	GetResourceOwnerId() *int64
	SetSceneId(v int64) *QueryTagInfoBySelectionRequest
	GetSceneId() *int64
	SetTagId(v int64) *QueryTagInfoBySelectionRequest
	GetTagId() *int64
}

type QueryTagInfoBySelectionRequest struct {
	// The industry ID.
	//
	// example:
	//
	// 58
	IndustryId           *int64  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 83
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 71
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryTagInfoBySelectionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTagInfoBySelectionRequest) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionRequest) GetIndustryId() *int64 {
	return s.IndustryId
}

func (s *QueryTagInfoBySelectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTagInfoBySelectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTagInfoBySelectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTagInfoBySelectionRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *QueryTagInfoBySelectionRequest) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryTagInfoBySelectionRequest) SetIndustryId(v int64) *QueryTagInfoBySelectionRequest {
	s.IndustryId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetOwnerId(v int64) *QueryTagInfoBySelectionRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetResourceOwnerAccount(v string) *QueryTagInfoBySelectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetResourceOwnerId(v int64) *QueryTagInfoBySelectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetSceneId(v int64) *QueryTagInfoBySelectionRequest {
	s.SceneId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) SetTagId(v int64) *QueryTagInfoBySelectionRequest {
	s.TagId = &v
	return s
}

func (s *QueryTagInfoBySelectionRequest) Validate() error {
	return dara.Validate(s)
}
