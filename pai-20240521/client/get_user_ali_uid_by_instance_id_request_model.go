// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAliUidByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetUserAliUidByInstanceIdRequest
	GetRegionId() *string
	SetResourceOwnerUid(v string) *GetUserAliUidByInstanceIdRequest
	GetResourceOwnerUid() *string
	SetResourceType(v string) *GetUserAliUidByInstanceIdRequest
	GetResourceType() *string
}

type GetUserAliUidByInstanceIdRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// PodId
	ResourceOwnerUid *string `json:"ResourceOwnerUid,omitempty" xml:"ResourceOwnerUid,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetUserAliUidByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserAliUidByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserAliUidByInstanceIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserAliUidByInstanceIdRequest) GetResourceOwnerUid() *string {
	return s.ResourceOwnerUid
}

func (s *GetUserAliUidByInstanceIdRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetUserAliUidByInstanceIdRequest) SetRegionId(v string) *GetUserAliUidByInstanceIdRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserAliUidByInstanceIdRequest) SetResourceOwnerUid(v string) *GetUserAliUidByInstanceIdRequest {
	s.ResourceOwnerUid = &v
	return s
}

func (s *GetUserAliUidByInstanceIdRequest) SetResourceType(v string) *GetUserAliUidByInstanceIdRequest {
	s.ResourceType = &v
	return s
}

func (s *GetUserAliUidByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
