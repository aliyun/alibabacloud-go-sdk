// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetWatermarkConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetWatermarkConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetWatermarkConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *GetWatermarkConsoleRequest
	GetResourceRealOwnerId() *int64
	SetWatermarkId(v string) *GetWatermarkConsoleRequest
	GetWatermarkId() *string
}

type GetWatermarkConsoleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s GetWatermarkConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkConsoleRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarkConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetWatermarkConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetWatermarkConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetWatermarkConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetWatermarkConsoleRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarkConsoleRequest) SetOwnerId(v int64) *GetWatermarkConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWatermarkConsoleRequest) SetResourceOwnerAccount(v string) *GetWatermarkConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWatermarkConsoleRequest) SetResourceOwnerId(v int64) *GetWatermarkConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWatermarkConsoleRequest) SetResourceRealOwnerId(v int64) *GetWatermarkConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetWatermarkConsoleRequest) SetWatermarkId(v string) *GetWatermarkConsoleRequest {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarkConsoleRequest) Validate() error {
	return dara.Validate(s)
}
