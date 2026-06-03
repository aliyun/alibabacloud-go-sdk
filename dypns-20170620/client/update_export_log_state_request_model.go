// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExportLogStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateExportLogStateRequest
	GetId() *int64
	SetOwnerId(v int64) *UpdateExportLogStateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateExportLogStateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateExportLogStateRequest
	GetResourceOwnerId() *int64
	SetState(v int32) *UpdateExportLogStateRequest
	GetState() *int32
}

type UpdateExportLogStateRequest struct {
	// This parameter is required.
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	State                *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateExportLogStateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExportLogStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateExportLogStateRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateExportLogStateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateExportLogStateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateExportLogStateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateExportLogStateRequest) GetState() *int32 {
	return s.State
}

func (s *UpdateExportLogStateRequest) SetId(v int64) *UpdateExportLogStateRequest {
	s.Id = &v
	return s
}

func (s *UpdateExportLogStateRequest) SetOwnerId(v int64) *UpdateExportLogStateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateExportLogStateRequest) SetResourceOwnerAccount(v string) *UpdateExportLogStateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateExportLogStateRequest) SetResourceOwnerId(v int64) *UpdateExportLogStateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateExportLogStateRequest) SetState(v int32) *UpdateExportLogStateRequest {
	s.State = &v
	return s
}

func (s *UpdateExportLogStateRequest) Validate() error {
	return dara.Validate(s)
}
