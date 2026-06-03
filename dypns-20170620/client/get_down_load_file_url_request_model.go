// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownLoadFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcctId(v int64) *GetDownLoadFileUrlRequest
	GetAcctId() *int64
	SetId(v int64) *GetDownLoadFileUrlRequest
	GetId() *int64
	SetOrderId(v int64) *GetDownLoadFileUrlRequest
	GetOrderId() *int64
	SetOwnerId(v int64) *GetDownLoadFileUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetDownLoadFileUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetDownLoadFileUrlRequest
	GetResourceOwnerId() *int64
}

type GetDownLoadFileUrlRequest struct {
	AcctId *int64 `json:"AcctId,omitempty" xml:"AcctId,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	OrderId              *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDownLoadFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDownLoadFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDownLoadFileUrlRequest) GetAcctId() *int64 {
	return s.AcctId
}

func (s *GetDownLoadFileUrlRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDownLoadFileUrlRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDownLoadFileUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDownLoadFileUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDownLoadFileUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetDownLoadFileUrlRequest) SetAcctId(v int64) *GetDownLoadFileUrlRequest {
	s.AcctId = &v
	return s
}

func (s *GetDownLoadFileUrlRequest) SetId(v int64) *GetDownLoadFileUrlRequest {
	s.Id = &v
	return s
}

func (s *GetDownLoadFileUrlRequest) SetOrderId(v int64) *GetDownLoadFileUrlRequest {
	s.OrderId = &v
	return s
}

func (s *GetDownLoadFileUrlRequest) SetOwnerId(v int64) *GetDownLoadFileUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDownLoadFileUrlRequest) SetResourceOwnerAccount(v string) *GetDownLoadFileUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDownLoadFileUrlRequest) SetResourceOwnerId(v int64) *GetDownLoadFileUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDownLoadFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
