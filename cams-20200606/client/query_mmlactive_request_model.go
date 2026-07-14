// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMMLActiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *QueryMMLActiveRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *QueryMMLActiveRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMMLActiveRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMMLActiveRequest
	GetResourceOwnerId() *int64
}

type QueryMMLActiveRequest struct {
	// example:
	//
	// erer-retreerew**
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMMLActiveRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMMLActiveRequest) GoString() string {
	return s.String()
}

func (s *QueryMMLActiveRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *QueryMMLActiveRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMMLActiveRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMMLActiveRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMMLActiveRequest) SetCustSpaceId(v string) *QueryMMLActiveRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryMMLActiveRequest) SetOwnerId(v int64) *QueryMMLActiveRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMMLActiveRequest) SetResourceOwnerAccount(v string) *QueryMMLActiveRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMMLActiveRequest) SetResourceOwnerId(v int64) *QueryMMLActiveRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMMLActiveRequest) Validate() error {
	return dara.Validate(s)
}
