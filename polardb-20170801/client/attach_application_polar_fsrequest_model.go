// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplicationPolarFSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AttachApplicationPolarFSRequest
	GetApplicationId() *string
	SetContainerMountPath(v string) *AttachApplicationPolarFSRequest
	GetContainerMountPath() *string
	SetPolarFSAccessKeyId(v string) *AttachApplicationPolarFSRequest
	GetPolarFSAccessKeyId() *string
	SetPolarFSAccessKeySecret(v string) *AttachApplicationPolarFSRequest
	GetPolarFSAccessKeySecret() *string
	SetPolarFSInstanceId(v string) *AttachApplicationPolarFSRequest
	GetPolarFSInstanceId() *string
}

type AttachApplicationPolarFSRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The mount point of the file system visible in the application container. If this parameter is not specified, the default value /mnt/<PolarFSInstanceId> is used for the first mount. For subsequent mounts, the current mount point of the instance remains unchanged. The value must be an absolute path that starts with a forward slash (/) but cannot be the root directory (/). It can only contain letters, digits, periods (.), underscores (_), hyphens (-), and forward slashes (/). The total length cannot exceed 128 characters, the directory depth cannot exceed 6 levels, and the length of a single level cannot exceed 63 characters. The mount point cannot be identical to or a subdirectory of the mount points of other file systems under the same application.
	//
	// example:
	//
	// /mnt/datasets
	ContainerMountPath *string `json:"ContainerMountPath,omitempty" xml:"ContainerMountPath,omitempty"`
	// The AccessKey ID of the Alibaba Cloud account or RAM user. Ensure that the account or RAM user is authorized when you create the PolarLakebase instance.
	//
	// example:
	//
	// LT**********************
	PolarFSAccessKeyId *string `json:"PolarFSAccessKeyId,omitempty" xml:"PolarFSAccessKeyId,omitempty"`
	// The AccessKey secret of the Alibaba Cloud account or RAM user.
	//
	// example:
	//
	// H3****************************
	PolarFSAccessKeySecret *string `json:"PolarFSAccessKeySecret,omitempty" xml:"PolarFSAccessKeySecret,omitempty"`
	// The PolarLakebase cold storage instance ID.
	//
	// Currently, only the following applications support cold storage instances:
	//
	// - supabase
	//
	// This parameter is required.
	//
	// example:
	//
	// pcs-**************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
}

func (s AttachApplicationPolarFSRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachApplicationPolarFSRequest) GoString() string {
	return s.String()
}

func (s *AttachApplicationPolarFSRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AttachApplicationPolarFSRequest) GetContainerMountPath() *string {
	return s.ContainerMountPath
}

func (s *AttachApplicationPolarFSRequest) GetPolarFSAccessKeyId() *string {
	return s.PolarFSAccessKeyId
}

func (s *AttachApplicationPolarFSRequest) GetPolarFSAccessKeySecret() *string {
	return s.PolarFSAccessKeySecret
}

func (s *AttachApplicationPolarFSRequest) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *AttachApplicationPolarFSRequest) SetApplicationId(v string) *AttachApplicationPolarFSRequest {
	s.ApplicationId = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetContainerMountPath(v string) *AttachApplicationPolarFSRequest {
	s.ContainerMountPath = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetPolarFSAccessKeyId(v string) *AttachApplicationPolarFSRequest {
	s.PolarFSAccessKeyId = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetPolarFSAccessKeySecret(v string) *AttachApplicationPolarFSRequest {
	s.PolarFSAccessKeySecret = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetPolarFSInstanceId(v string) *AttachApplicationPolarFSRequest {
	s.PolarFSInstanceId = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) Validate() error {
	return dara.Validate(s)
}
