// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuppressionListLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *SetSuppressionListLevelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetSuppressionListLevelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetSuppressionListLevelRequest
	GetResourceOwnerId() *int64
	SetSuppressionListLevel(v string) *SetSuppressionListLevelRequest
	GetSuppressionListLevel() *string
}

type SetSuppressionListLevelRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SuppressionListLevel *string `json:"SuppressionListLevel,omitempty" xml:"SuppressionListLevel,omitempty"`
}

func (s SetSuppressionListLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSuppressionListLevelRequest) GoString() string {
	return s.String()
}

func (s *SetSuppressionListLevelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetSuppressionListLevelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetSuppressionListLevelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetSuppressionListLevelRequest) GetSuppressionListLevel() *string {
	return s.SuppressionListLevel
}

func (s *SetSuppressionListLevelRequest) SetOwnerId(v int64) *SetSuppressionListLevelRequest {
	s.OwnerId = &v
	return s
}

func (s *SetSuppressionListLevelRequest) SetResourceOwnerAccount(v string) *SetSuppressionListLevelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetSuppressionListLevelRequest) SetResourceOwnerId(v int64) *SetSuppressionListLevelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetSuppressionListLevelRequest) SetSuppressionListLevel(v string) *SetSuppressionListLevelRequest {
	s.SuppressionListLevel = &v
	return s
}

func (s *SetSuppressionListLevelRequest) Validate() error {
	return dara.Validate(s)
}
