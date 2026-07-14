// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViberPauseTimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetViberPauseTimesRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetViberPauseTimesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetViberPauseTimesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetViberPauseTimesRequest
	GetResourceOwnerId() *int64
}

type GetViberPauseTimesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-x***
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetViberPauseTimesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetViberPauseTimesRequest) GoString() string {
	return s.String()
}

func (s *GetViberPauseTimesRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetViberPauseTimesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetViberPauseTimesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetViberPauseTimesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetViberPauseTimesRequest) SetCustSpaceId(v string) *GetViberPauseTimesRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetViberPauseTimesRequest) SetOwnerId(v int64) *GetViberPauseTimesRequest {
	s.OwnerId = &v
	return s
}

func (s *GetViberPauseTimesRequest) SetResourceOwnerAccount(v string) *GetViberPauseTimesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetViberPauseTimesRequest) SetResourceOwnerId(v int64) *GetViberPauseTimesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetViberPauseTimesRequest) Validate() error {
	return dara.Validate(s)
}
