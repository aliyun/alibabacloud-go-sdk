// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QuerySmarttagJobRequest
	GetJobId() *string
	SetOwnerAccount(v string) *QuerySmarttagJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QuerySmarttagJobRequest
	GetOwnerId() *int64
	SetParams(v string) *QuerySmarttagJobRequest
	GetParams() *string
	SetResourceOwnerAccount(v string) *QuerySmarttagJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmarttagJobRequest
	GetResourceOwnerId() *int64
}

type QuerySmarttagJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// {"labelResultType":"auto"}
	Params               *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmarttagJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagJobRequest) GoString() string {
	return s.String()
}

func (s *QuerySmarttagJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QuerySmarttagJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QuerySmarttagJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmarttagJobRequest) GetParams() *string {
	return s.Params
}

func (s *QuerySmarttagJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmarttagJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmarttagJobRequest) SetJobId(v string) *QuerySmarttagJobRequest {
	s.JobId = &v
	return s
}

func (s *QuerySmarttagJobRequest) SetOwnerAccount(v string) *QuerySmarttagJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QuerySmarttagJobRequest) SetOwnerId(v int64) *QuerySmarttagJobRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmarttagJobRequest) SetParams(v string) *QuerySmarttagJobRequest {
	s.Params = &v
	return s
}

func (s *QuerySmarttagJobRequest) SetResourceOwnerAccount(v string) *QuerySmarttagJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmarttagJobRequest) SetResourceOwnerId(v int64) *QuerySmarttagJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmarttagJobRequest) Validate() error {
	return dara.Validate(s)
}
