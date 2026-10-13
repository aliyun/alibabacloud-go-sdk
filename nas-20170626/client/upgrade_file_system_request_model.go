// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *UpgradeFileSystemRequest
	GetCapacity() *int64
	SetClientToken(v string) *UpgradeFileSystemRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpgradeFileSystemRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *UpgradeFileSystemRequest
	GetFileSystemId() *string
}

type UpgradeFileSystemRequest struct {
	// The storage capacity of the file system after the scale-up.
	//
	// The new capacity must be greater than the current capacity of the file system. Unit: GiB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure the idempotency of the request. Generate a unique value from your client to ensure that the value is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating an instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating an instance. The check items include required parameters, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - Extreme NAS: The ID must start with `extreme-`, for example, extreme-0015\\*\\*\\*\\*.
	//
	// - CPFS: The ID must start with `cpfs-`, for example, cpfs-125487\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme-0015****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s UpgradeFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeFileSystemRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *UpgradeFileSystemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeFileSystemRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpgradeFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *UpgradeFileSystemRequest) SetCapacity(v int64) *UpgradeFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetClientToken(v string) *UpgradeFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetDryRun(v bool) *UpgradeFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetFileSystemId(v string) *UpgradeFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpgradeFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
