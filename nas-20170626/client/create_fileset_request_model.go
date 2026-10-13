// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFilesetRequest
	GetClientToken() *string
	SetDeletionProtection(v bool) *CreateFilesetRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *CreateFilesetRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateFilesetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateFilesetRequest
	GetFileSystemId() *string
	SetFileSystemPath(v string) *CreateFilesetRequest
	GetFileSystemPath() *string
	SetQuota(v *CreateFilesetRequestQuota) *CreateFilesetRequest
	GetQuota() *CreateFilesetRequestQuota
}

type CreateFilesetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266554*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The release protection property of the instance. Specifies whether the instance can be released by using the console or the [DeleteFileset](https://help.aliyun.com/document_detail/2402263.html) operation.
	//
	// - true: Enables release protection for the instance.
	//
	// - false (default): Disables release protection for the instance.
	//
	// > This property only restricts manual release operations and does not take effect on system-initiated release operations.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the fileset.
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter or Chinese character and cannot start with http:// or https://.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// Internal affairs directory
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating the instance. The check items include required parameters, request format, business restrictions, and NAS inventory. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned, but FsetId is empty.
	//
	// - false (default): Sends a normal request. After the check succeeds, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64y*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the fileset to create.
	//
	// - Path restrictions for CPFS:
	//
	//   - The parent directory of the specified directory must be an existing directory in the file system.
	//
	//   - The path must be 2 to 1,024 characters in length.
	//
	//   - The specified directory must start and end with a forward slash (/).
	//
	// - Path restrictions for CPFS for Lingjun:
	//
	//   - The path must be 2 to 1,024 characters in length.
	//
	//   - The specified directory must start and end with a forward slash (/).
	//
	//   - The fileset path must be a new path and cannot be an existing path. Renaming a fileset path is not supported, and the path cannot be a symbolic link.
	//
	//   - The maximum depth of a fileset path is 8 levels, with the root directory / as level 0. For example, a fileset path of /test/aaa/ccc/ has a depth of 3 levels.
	//
	//   - If the fileset path contains multiple directory levels, the parent directory must be an existing directory.
	//
	//   - Nesting filesets within a fileset is not supported. If a parent directory is already specified as a fileset, its subdirectories cannot be specified as filesets. Only one quota is supported for each fileset path.
	//
	//   - The total path length cannot exceed 990 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The quota information.
	//
	// > This parameter is supported only for CPFS for Lingjun 2.7.0 and later.
	Quota *CreateFilesetRequestQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s CreateFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFilesetRequest) GoString() string {
	return s.String()
}

func (s *CreateFilesetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFilesetRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateFilesetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFilesetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateFilesetRequest) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *CreateFilesetRequest) GetQuota() *CreateFilesetRequestQuota {
	return s.Quota
}

func (s *CreateFilesetRequest) SetClientToken(v string) *CreateFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFilesetRequest) SetDeletionProtection(v bool) *CreateFilesetRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateFilesetRequest) SetDescription(v string) *CreateFilesetRequest {
	s.Description = &v
	return s
}

func (s *CreateFilesetRequest) SetDryRun(v bool) *CreateFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemId(v string) *CreateFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemPath(v string) *CreateFilesetRequest {
	s.FileSystemPath = &v
	return s
}

func (s *CreateFilesetRequest) SetQuota(v *CreateFilesetRequestQuota) *CreateFilesetRequest {
	s.Quota = v
	return s
}

func (s *CreateFilesetRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFilesetRequestQuota struct {
	// The file count limit of the quota. Valid values:
	//
	// - Minimum value: 10,000.
	//
	// - Maximum value: 10,000,000,000.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The total capacity limit of the quota. Unit: bytes.
	//
	// Valid values:
	//
	// - Minimum value: 10,737,418,240 (10 GiB).
	//
	// - Increment: 1,073,741,824 (1 GiB).
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateFilesetRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateFilesetRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateFilesetRequestQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *CreateFilesetRequestQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *CreateFilesetRequestQuota) SetFileCountLimit(v int64) *CreateFilesetRequestQuota {
	s.FileCountLimit = &v
	return s
}

func (s *CreateFilesetRequestQuota) SetSizeLimit(v int64) *CreateFilesetRequestQuota {
	s.SizeLimit = &v
	return s
}

func (s *CreateFilesetRequestQuota) Validate() error {
	return dara.Validate(s)
}
