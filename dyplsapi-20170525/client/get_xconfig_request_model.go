// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerParentId(v int64) *GetXConfigRequest
	GetCallerParentId() *int64
	SetCustomerPoolKey(v string) *GetXConfigRequest
	GetCustomerPoolKey() *string
	SetOwnerId(v int64) *GetXConfigRequest
	GetOwnerId() *int64
	SetReqId(v string) *GetXConfigRequest
	GetReqId() *string
	SetResourceOwnerAccount(v string) *GetXConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetXConfigRequest
	GetResourceOwnerId() *int64
	SetTelX(v string) *GetXConfigRequest
	GetTelX() *string
}

type GetXConfigRequest struct {
	// The customer UID. This parameter is automatically obtained by the system and does not need to be specified.
	//
	// example:
	//
	// -
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// The key of the number pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC5**********************a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request deduplication ID. The maximum length is 20 characters. Ensure the atomicity of the request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 564**********879
	ReqId                *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The X number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17*******22
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
}

func (s GetXConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigRequest) GoString() string {
	return s.String()
}

func (s *GetXConfigRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *GetXConfigRequest) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *GetXConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetXConfigRequest) GetReqId() *string {
	return s.ReqId
}

func (s *GetXConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetXConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetXConfigRequest) GetTelX() *string {
	return s.TelX
}

func (s *GetXConfigRequest) SetCallerParentId(v int64) *GetXConfigRequest {
	s.CallerParentId = &v
	return s
}

func (s *GetXConfigRequest) SetCustomerPoolKey(v string) *GetXConfigRequest {
	s.CustomerPoolKey = &v
	return s
}

func (s *GetXConfigRequest) SetOwnerId(v int64) *GetXConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetXConfigRequest) SetReqId(v string) *GetXConfigRequest {
	s.ReqId = &v
	return s
}

func (s *GetXConfigRequest) SetResourceOwnerAccount(v string) *GetXConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetXConfigRequest) SetResourceOwnerId(v int64) *GetXConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetXConfigRequest) SetTelX(v string) *GetXConfigRequest {
	s.TelX = &v
	return s
}

func (s *GetXConfigRequest) Validate() error {
	return dara.Validate(s)
}
