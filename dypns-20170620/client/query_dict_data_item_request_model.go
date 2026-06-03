// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDictDataItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassifyItemCode(v string) *QueryDictDataItemRequest
	GetClassifyItemCode() *string
	SetOwnerId(v int64) *QueryDictDataItemRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryDictDataItemRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryDictDataItemRequest
	GetResourceOwnerId() *int64
}

type QueryDictDataItemRequest struct {
	// This parameter is required.
	ClassifyItemCode     *string `json:"ClassifyItemCode,omitempty" xml:"ClassifyItemCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryDictDataItemRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDictDataItemRequest) GoString() string {
	return s.String()
}

func (s *QueryDictDataItemRequest) GetClassifyItemCode() *string {
	return s.ClassifyItemCode
}

func (s *QueryDictDataItemRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryDictDataItemRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryDictDataItemRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryDictDataItemRequest) SetClassifyItemCode(v string) *QueryDictDataItemRequest {
	s.ClassifyItemCode = &v
	return s
}

func (s *QueryDictDataItemRequest) SetOwnerId(v int64) *QueryDictDataItemRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDictDataItemRequest) SetResourceOwnerAccount(v string) *QueryDictDataItemRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDictDataItemRequest) SetResourceOwnerId(v int64) *QueryDictDataItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryDictDataItemRequest) Validate() error {
	return dara.Validate(s)
}
