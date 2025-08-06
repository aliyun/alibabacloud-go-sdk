// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarksConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetWatermarksConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetWatermarksConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetWatermarksConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *GetWatermarksConsoleRequest
	GetResourceRealOwnerId() *int64
}

type GetWatermarksConsoleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s GetWatermarksConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetWatermarksConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetWatermarksConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetWatermarksConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetWatermarksConsoleRequest) SetOwnerId(v int64) *GetWatermarksConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWatermarksConsoleRequest) SetResourceOwnerAccount(v string) *GetWatermarksConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWatermarksConsoleRequest) SetResourceOwnerId(v int64) *GetWatermarksConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWatermarksConsoleRequest) SetResourceRealOwnerId(v int64) *GetWatermarksConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetWatermarksConsoleRequest) Validate() error {
	return dara.Validate(s)
}
