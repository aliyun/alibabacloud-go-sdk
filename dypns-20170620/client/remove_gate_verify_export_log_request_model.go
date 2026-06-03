// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGateVerifyExportLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RemoveGateVerifyExportLogRequest
	GetId() *int64
	SetOwnerId(v int64) *RemoveGateVerifyExportLogRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveGateVerifyExportLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveGateVerifyExportLogRequest
	GetResourceOwnerId() *int64
}

type RemoveGateVerifyExportLogRequest struct {
	// This parameter is required.
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveGateVerifyExportLogRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveGateVerifyExportLogRequest) GoString() string {
	return s.String()
}

func (s *RemoveGateVerifyExportLogRequest) GetId() *int64 {
	return s.Id
}

func (s *RemoveGateVerifyExportLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveGateVerifyExportLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveGateVerifyExportLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveGateVerifyExportLogRequest) SetId(v int64) *RemoveGateVerifyExportLogRequest {
	s.Id = &v
	return s
}

func (s *RemoveGateVerifyExportLogRequest) SetOwnerId(v int64) *RemoveGateVerifyExportLogRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveGateVerifyExportLogRequest) SetResourceOwnerAccount(v string) *RemoveGateVerifyExportLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveGateVerifyExportLogRequest) SetResourceOwnerId(v int64) *RemoveGateVerifyExportLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveGateVerifyExportLogRequest) Validate() error {
	return dara.Validate(s)
}
