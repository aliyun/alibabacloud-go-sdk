// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyFilesetRequest
	GetClientToken() *string
	SetDeletionProtection(v bool) *ModifyFilesetRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *ModifyFilesetRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyFilesetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyFilesetRequest
	GetFileSystemId() *string
	SetFsetId(v string) *ModifyFilesetRequest
	GetFsetId() *string
}

type ModifyFilesetRequest struct {
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The release protection property of the instance. Specifies whether the instance can be released by using the console or the [DeleteFileset](https://help.aliyun.com/document_detail/2402263.html) operation.
	//
	//
	//
	// -   true: Enables release protection for the instance.
	//
	// -   false: Disables release protection for the instance.
	//
	// > This property only restricts manual release operations and does not take effect on system-initiated release operations.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the fileset.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this modification request.
	//
	// A dry run checks parameter validity and inventory without actually modifying the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without modifying the instance. The check items include whether required parameters are specified, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-099394bd928c****.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya****.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Fileset ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s ModifyFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFilesetRequest) GoString() string {
	return s.String()
}

func (s *ModifyFilesetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyFilesetRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyFilesetRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFilesetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFilesetRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *ModifyFilesetRequest) SetClientToken(v string) *ModifyFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFilesetRequest) SetDeletionProtection(v bool) *ModifyFilesetRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyFilesetRequest) SetDescription(v string) *ModifyFilesetRequest {
	s.Description = &v
	return s
}

func (s *ModifyFilesetRequest) SetDryRun(v bool) *ModifyFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFilesetRequest) SetFileSystemId(v string) *ModifyFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFilesetRequest) SetFsetId(v string) *ModifyFilesetRequest {
	s.FsetId = &v
	return s
}

func (s *ModifyFilesetRequest) Validate() error {
	return dara.Validate(s)
}
