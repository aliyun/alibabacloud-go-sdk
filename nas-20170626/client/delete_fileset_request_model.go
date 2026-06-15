// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFilesetRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteFilesetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteFilesetRequest
	GetFileSystemId() *string
	SetFsetId(v string) *DeleteFilesetRequest
	GetFsetId() *string
}

type DeleteFilesetRequest struct {
	// A client-generated token that you can use to ensure the idempotence of the request. The ClientToken must be unique across requests.
	//
	// The ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the ClientToken. The request ID is unique for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for the request.
	//
	// A dry run checks for issues such as parameter validity and resource availability, but does not delete the fileset.
	//
	// Valid values:
	//
	// - true: Sends a check request and does not delete the fileset. The system checks for required parameters, request format, and business limits. If the check fails, an error is returned. If the check passes, an HTTP 200 OK status code is returned.
	//
	// - false (Default): Sends a normal request and deletes the fileset after the check passes.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// - CPFS for AI and HPC: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The fileset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s DeleteFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFilesetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteFilesetRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *DeleteFilesetRequest) SetClientToken(v string) *DeleteFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFilesetRequest) SetDryRun(v bool) *DeleteFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFilesetRequest) SetFileSystemId(v string) *DeleteFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFilesetRequest) SetFsetId(v string) *DeleteFilesetRequest {
	s.FsetId = &v
	return s
}

func (s *DeleteFilesetRequest) Validate() error {
	return dara.Validate(s)
}
