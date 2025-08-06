// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultWatermarkConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *SetDefaultWatermarkConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetDefaultWatermarkConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetDefaultWatermarkConsoleRequest
	GetResourceOwnerId() *int64
	SetWatermarkId(v string) *SetDefaultWatermarkConsoleRequest
	GetWatermarkId() *string
}

type SetDefaultWatermarkConsoleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s SetDefaultWatermarkConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultWatermarkConsoleRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDefaultWatermarkConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDefaultWatermarkConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetDefaultWatermarkConsoleRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *SetDefaultWatermarkConsoleRequest) SetOwnerId(v int64) *SetDefaultWatermarkConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultWatermarkConsoleRequest) SetResourceOwnerAccount(v string) *SetDefaultWatermarkConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultWatermarkConsoleRequest) SetResourceOwnerId(v int64) *SetDefaultWatermarkConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultWatermarkConsoleRequest) SetWatermarkId(v string) *SetDefaultWatermarkConsoleRequest {
	s.WatermarkId = &v
	return s
}

func (s *SetDefaultWatermarkConsoleRequest) Validate() error {
	return dara.Validate(s)
}
