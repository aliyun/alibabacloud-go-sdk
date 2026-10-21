// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRefreshInterval(v int64) *CreateDataFlowRequest
	GetAutoRefreshInterval() *int64
	SetAutoRefreshPolicy(v string) *CreateDataFlowRequest
	GetAutoRefreshPolicy() *string
	SetAutoRefreshs(v []*CreateDataFlowRequestAutoRefreshs) *CreateDataFlowRequest
	GetAutoRefreshs() []*CreateDataFlowRequestAutoRefreshs
	SetClientToken(v string) *CreateDataFlowRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDataFlowRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateDataFlowRequest
	GetFileSystemId() *string
	SetFileSystemPath(v string) *CreateDataFlowRequest
	GetFileSystemPath() *string
	SetFsetId(v string) *CreateDataFlowRequest
	GetFsetId() *string
	SetSourceSecurityType(v string) *CreateDataFlowRequest
	GetSourceSecurityType() *string
	SetSourceStorage(v string) *CreateDataFlowRequest
	GetSourceStorage() *string
	SetSourceStoragePath(v string) *CreateDataFlowRequest
	GetSourceStoragePath() *string
	SetThroughput(v int64) *CreateDataFlowRequest
	GetThroughput() *int64
}

type CreateDataFlowRequest struct {
	// The auto-refresh interval. General-purpose CPFS checks whether data updates exist in the directory at this interval. If data updates exist, an auto-refresh task is started. Unit: minutes.
	//
	// Valid values: 10 to 525600. Default value: 10.
	//
	// > This parameter takes effect only when the file system type is General-purpose CPFS.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The auto-refresh policy. This policy determines how data updates from the source storage are imported to General-purpose CPFS.
	//
	// - None: Data updates from the source storage are not automatically imported to General-purpose CPFS. You can use a data flow task to import the updates. This is the default value.
	//
	// - ImportChanged: Data updates from the source storage are automatically imported to General-purpose CPFS.
	//
	// > This parameter takes effect only when the file system type is General-purpose CPFS.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	// The auto-refresh configurations.
	//
	// > This parameter takes effect only when the file system type is General-purpose CPFS.
	//
	// if can be null:
	// false
	AutoRefreshs []*CreateDataFlowRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can generate the token on your client to ensure that the token is unique across different requests.
	//
	// The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the data flow.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Bucket01 DataFlow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run on the creation request.
	//
	// A dry run helps you check parameter validity and resource availability without actually creating the instance or incurring fees.
	//
	// Valid values:
	//
	// - true: A dry run is performed and no instance is created. The system checks whether required parameters are specified, the request format, business limits, and NAS resource availability. If the request fails the dry run, an error is returned. If the request passes the dry run, the HTTP status code 200 is returned, but the FileSystemId parameter is empty.
	//
	// - false: A regular request is sent. If the request passes the dry run, the instance is created. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - General-purpose CPFS: The ID must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directory in the CPFS for Lingjun file system. Limits:
	//
	// - The directory must start and end with a forward slash (/).
	//
	// - The directory must already exist in the CPFS for Lingjun file system.
	//
	// - The directory must be 1 to 1,023 characters in length.
	//
	// - The directory must be encoded in UTF-8.
	//
	// > This parameter is required when the file system type is CPFS for Lingjun.
	//
	// example:
	//
	// /path/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The fileset ID.
	//
	// > This parameter is required when the file system type is General-purpose CPFS.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type. Valid values:
	//
	// - None: The source storage does not require security protection for access. This is the default value.
	//
	// - SSL: Access is protected by an SSL certificate.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// SSL
	SourceSecurityType *string `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	// The access address of the source storage. Format: `<storage type>://[<account id>:]<path>`.
	//
	// Where:
	//
	// - storage type: Only oss is supported.
	//
	// - account id: Optional. The UID of the account to which the source storage belongs. This parameter is required when you use a cross-account OSS bucket.
	//
	// - path: The name of the OSS bucket. Limits:
	//
	//     - The name can contain only lowercase letters, digits, and hyphens (-), and must start and end with a lowercase letter or a digit.
	//
	//     - The name can be up to 128 characters in length.
	//
	//     - The name must be encoded in UTF-8.
	//
	// > - The OSS bucket must already exist in the region.
	//
	// > - The account id parameter is supported only in CPFS for Lingjun 2.6.0 and later.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://178321033379****:bucket-01
	SourceStorage *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// The access path in the source storage bucket. Limits:
	//
	// - The path must start and end with a forward slash (/).
	//
	// - The path is case-sensitive.
	//
	// - The path must be 1 to 1,023 characters in length.
	//
	// - The path must be encoded in UTF-8.
	//
	// > This parameter is required when the file system type is CPFS for Lingjun.
	//
	// example:
	//
	// /prefix/
	SourceStoragePath *string `json:"SourceStoragePath,omitempty" xml:"SourceStoragePath,omitempty"`
	// The maximum transmission bandwidth of the data flow. Unit: MB/s. Valid values:
	//
	// - 600
	//
	// - 1200
	//
	// - 1500
	//
	// > The transmission bandwidth of the data flow must be less than the I/O bandwidth of the file system.
	//
	// > This parameter is required when the file system type is General-purpose CPFS.
	//
	// example:
	//
	// 600
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s CreateDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequest) GetAutoRefreshInterval() *int64 {
	return s.AutoRefreshInterval
}

func (s *CreateDataFlowRequest) GetAutoRefreshPolicy() *string {
	return s.AutoRefreshPolicy
}

func (s *CreateDataFlowRequest) GetAutoRefreshs() []*CreateDataFlowRequestAutoRefreshs {
	return s.AutoRefreshs
}

func (s *CreateDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDataFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateDataFlowRequest) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *CreateDataFlowRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *CreateDataFlowRequest) GetSourceSecurityType() *string {
	return s.SourceSecurityType
}

func (s *CreateDataFlowRequest) GetSourceStorage() *string {
	return s.SourceStorage
}

func (s *CreateDataFlowRequest) GetSourceStoragePath() *string {
	return s.SourceStoragePath
}

func (s *CreateDataFlowRequest) GetThroughput() *int64 {
	return s.Throughput
}

func (s *CreateDataFlowRequest) SetAutoRefreshInterval(v int64) *CreateDataFlowRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshPolicy(v string) *CreateDataFlowRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshs(v []*CreateDataFlowRequestAutoRefreshs) *CreateDataFlowRequest {
	s.AutoRefreshs = v
	return s
}

func (s *CreateDataFlowRequest) SetClientToken(v string) *CreateDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowRequest) SetDescription(v string) *CreateDataFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateDataFlowRequest) SetDryRun(v bool) *CreateDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowRequest) SetFileSystemId(v string) *CreateDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowRequest) SetFileSystemPath(v string) *CreateDataFlowRequest {
	s.FileSystemPath = &v
	return s
}

func (s *CreateDataFlowRequest) SetFsetId(v string) *CreateDataFlowRequest {
	s.FsetId = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceSecurityType(v string) *CreateDataFlowRequest {
	s.SourceSecurityType = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceStorage(v string) *CreateDataFlowRequest {
	s.SourceStorage = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceStoragePath(v string) *CreateDataFlowRequest {
	s.SourceStoragePath = &v
	return s
}

func (s *CreateDataFlowRequest) SetThroughput(v int64) *CreateDataFlowRequest {
	s.Throughput = &v
	return s
}

func (s *CreateDataFlowRequest) Validate() error {
	if s.AutoRefreshs != nil {
		for _, item := range s.AutoRefreshs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataFlowRequestAutoRefreshs struct {
	// The auto-refresh directory. General-purpose CPFS registers data modification events of the source storage, checks whether the source data in this directory is updated, and automatically imports the updated data.
	//
	// By default, this parameter is empty. Data updates in the source storage are not automatically imported to General-purpose CPFS. You must create a manual task to import the updates.
	//
	// Limits:
	//
	// - The directory must be 2 to 1,024 characters in length.
	//
	// - The directory must be encoded in UTF-8.
	//
	// - The directory must start and end with a forward slash (/).
	//
	// - The directory must already exist in General-purpose CPFS and must be located in the fileset directory of the data flow.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// /prefix1/prefix2/
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CreateDataFlowRequestAutoRefreshs) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequestAutoRefreshs) GetRefreshPath() *string {
	return s.RefreshPath
}

func (s *CreateDataFlowRequestAutoRefreshs) SetRefreshPath(v string) *CreateDataFlowRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

func (s *CreateDataFlowRequestAutoRefreshs) Validate() error {
	return dara.Validate(s)
}
