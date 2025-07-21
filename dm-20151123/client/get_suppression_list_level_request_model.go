// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuppressionListLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetSuppressionListLevelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSuppressionListLevelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSuppressionListLevelRequest
	GetResourceOwnerId() *int64
}

type GetSuppressionListLevelRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetSuppressionListLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSuppressionListLevelRequest) GoString() string {
	return s.String()
}

func (s *GetSuppressionListLevelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSuppressionListLevelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSuppressionListLevelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSuppressionListLevelRequest) SetOwnerId(v int64) *GetSuppressionListLevelRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSuppressionListLevelRequest) SetResourceOwnerAccount(v string) *GetSuppressionListLevelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSuppressionListLevelRequest) SetResourceOwnerId(v int64) *GetSuppressionListLevelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSuppressionListLevelRequest) Validate() error {
	return dara.Validate(s)
}
