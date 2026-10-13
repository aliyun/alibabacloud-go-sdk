// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteProtocolMountTargetRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteProtocolMountTargetRequest
	GetDryRun() *bool
	SetExportId(v string) *DeleteProtocolMountTargetRequest
	GetExportId() *string
	SetFileSystemId(v string) *DeleteProtocolMountTargetRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *DeleteProtocolMountTargetRequest
	GetProtocolServiceId() *string
}

type DeleteProtocolMountTargetRequest struct {
	// Ensures the idempotence of the request. Generate a value from your client that is unique among different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this delete request.
	//
	// A dry run checks parameter validity and dependencies without actually deleting the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without deleting the export directory. The dry run checks whether required parameters are specified, whether the request format is valid, and whether business limit dependencies are met. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly deleted.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the export directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-19abf5beab8d****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the protocol service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteProtocolMountTargetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteProtocolMountTargetRequest) GetExportId() *string {
	return s.ExportId
}

func (s *DeleteProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteProtocolMountTargetRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *DeleteProtocolMountTargetRequest) SetClientToken(v string) *DeleteProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetDryRun(v bool) *DeleteProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetExportId(v string) *DeleteProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetFileSystemId(v string) *DeleteProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetProtocolServiceId(v string) *DeleteProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
