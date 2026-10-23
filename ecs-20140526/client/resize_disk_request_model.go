// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResizeDiskRequest
	GetClientToken() *string
	SetDiskId(v string) *ResizeDiskRequest
	GetDiskId() *string
	SetNewSize(v int32) *ResizeDiskRequest
	GetNewSize() *int32
	SetOwnerAccount(v string) *ResizeDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResizeDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResizeDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResizeDiskRequest
	GetResourceOwnerId() *int64
	SetType(v string) *ResizeDiskRequest
	GetType() *string
}

type ResizeDiskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cloud disk. You can call the [DescribeDisks](https://help.aliyun.com/document_detail/25514.html) operation to query the ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The new capacity of the cloud disk. Unit: GiB. Valid values:
	//
	// - System cloud disk:
	//
	//   - Basic cloud disk: 20 to 500.
	//
	//   - ESSD:
	//
	//     - PL0: 1 to 2048.
	//
	//     - PL1: 20 to 2048.
	//
	//     - PL2: 461 to 2048.
	//
	//     - PL3: 1261 to 2048.
	//
	//   - ESSD AutoPL disk: 1 to 2048.
	//
	//   - Other cloud disk types: 20 to 2048.
	//
	// - Data cloud disk:
	//
	//     - Ultra cloud disk (cloud_efficiency): 20 to 32768.
	//
	//     - Standard SSD (cloud_ssd): 20 to 32768.
	//
	//     - ESSD (cloud_essd): The valid values vary based on the `PerformanceLevel` parameter. You can call the [DescribeDisks](https://help.aliyun.com/document_detail/25514.html) operation to query the cloud disk information and check the valid values based on the `PerformanceLevel` parameter in the response.
	//
	//         - PL0: 1 to 65536.
	//
	//         - PL1: 20 to 65536.
	//
	//         - PL2: 461 to 65536.
	//
	//         - PL3: 1261 to 65536.
	//
	//     - Basic cloud disk (cloud): 5 to 2000.
	//
	//     - ESSD AutoPL disk (cloud_auto): 1 to 65536.
	//
	// <props="china">
	//
	//     - ESSD Entry disk (cloud_essd_entry): 10 to 32768.
	//
	//   - Standard elastic ephemeral disk (elastic_ephemeral_disk_standard): 64 to 8192.
	//
	//   - Premium elastic ephemeral disk (elastic_ephemeral_disk_premium): 64 to 8192.
	//
	// > The new capacity must be greater than the original capacity. Otherwise, an error is returned when you call this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1900
	NewSize              *int32  `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The method used to resize the cloud disk. Valid values:
	//
	// - offline: offline resizing. This is the default value. After the cloud disk is resized, you must [restart the instance](https://help.aliyun.com/document_detail/25440.html) in the console or call the [RebootInstance](https://help.aliyun.com/document_detail/25502.html) operation for the resize to take effect.
	//
	//
	//
	// - online: online resizing. The cloud disk is resized without the need to restart the instance. This value supports ultra cloud disks, standard SSDs, ESSDs, and elastic ephemeral disks.
	//
	// example:
	//
	// offline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResizeDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskRequest) GoString() string {
	return s.String()
}

func (s *ResizeDiskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResizeDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResizeDiskRequest) GetNewSize() *int32 {
	return s.NewSize
}

func (s *ResizeDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResizeDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResizeDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResizeDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResizeDiskRequest) GetType() *string {
	return s.Type
}

func (s *ResizeDiskRequest) SetClientToken(v string) *ResizeDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *ResizeDiskRequest) SetDiskId(v string) *ResizeDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResizeDiskRequest) SetNewSize(v int32) *ResizeDiskRequest {
	s.NewSize = &v
	return s
}

func (s *ResizeDiskRequest) SetOwnerAccount(v string) *ResizeDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResizeDiskRequest) SetOwnerId(v int64) *ResizeDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *ResizeDiskRequest) SetResourceOwnerAccount(v string) *ResizeDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResizeDiskRequest) SetResourceOwnerId(v int64) *ResizeDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResizeDiskRequest) SetType(v string) *ResizeDiskRequest {
	s.Type = &v
	return s
}

func (s *ResizeDiskRequest) Validate() error {
	return dara.Validate(s)
}
