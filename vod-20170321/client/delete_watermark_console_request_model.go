// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWatermarkConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteWatermarkConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteWatermarkConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteWatermarkConsoleRequest
	GetResourceOwnerId() *int64
	SetWatermarkId(v string) *DeleteWatermarkConsoleRequest
	GetWatermarkId() *string
}

type DeleteWatermarkConsoleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s DeleteWatermarkConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWatermarkConsoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteWatermarkConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteWatermarkConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteWatermarkConsoleRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *DeleteWatermarkConsoleRequest) SetOwnerId(v int64) *DeleteWatermarkConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteWatermarkConsoleRequest) SetResourceOwnerAccount(v string) *DeleteWatermarkConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteWatermarkConsoleRequest) SetResourceOwnerId(v int64) *DeleteWatermarkConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteWatermarkConsoleRequest) SetWatermarkId(v string) *DeleteWatermarkConsoleRequest {
	s.WatermarkId = &v
	return s
}

func (s *DeleteWatermarkConsoleRequest) Validate() error {
	return dara.Validate(s)
}
