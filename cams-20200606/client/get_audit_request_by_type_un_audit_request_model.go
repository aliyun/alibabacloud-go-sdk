// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditRequestByTypeUnAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetAuditRequestByTypeUnAuditRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetAuditRequestByTypeUnAuditRequest
	GetOwnerId() *int64
	SetRequestType(v string) *GetAuditRequestByTypeUnAuditRequest
	GetRequestType() *string
	SetResourceOwnerAccount(v string) *GetAuditRequestByTypeUnAuditRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAuditRequestByTypeUnAuditRequest
	GetResourceOwnerId() *int64
}

type GetAuditRequestByTypeUnAuditRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// viberOpen
	RequestType          *string `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAuditRequestByTypeUnAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditRequestByTypeUnAuditRequest) GoString() string {
	return s.String()
}

func (s *GetAuditRequestByTypeUnAuditRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetAuditRequestByTypeUnAuditRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAuditRequestByTypeUnAuditRequest) GetRequestType() *string {
	return s.RequestType
}

func (s *GetAuditRequestByTypeUnAuditRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAuditRequestByTypeUnAuditRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAuditRequestByTypeUnAuditRequest) SetCustSpaceId(v string) *GetAuditRequestByTypeUnAuditRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditRequest) SetOwnerId(v int64) *GetAuditRequestByTypeUnAuditRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditRequest) SetRequestType(v string) *GetAuditRequestByTypeUnAuditRequest {
	s.RequestType = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditRequest) SetResourceOwnerAccount(v string) *GetAuditRequestByTypeUnAuditRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditRequest) SetResourceOwnerId(v int64) *GetAuditRequestByTypeUnAuditRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuditRequestByTypeUnAuditRequest) Validate() error {
	return dara.Validate(s)
}
