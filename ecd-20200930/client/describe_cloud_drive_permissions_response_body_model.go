// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDrivePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudDrivePermissionModels(v []*DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) *DescribeCloudDrivePermissionsResponseBody
	GetCloudDrivePermissionModels() []*DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels
	SetRequestId(v string) *DescribeCloudDrivePermissionsResponseBody
	GetRequestId() *string
}

type DescribeCloudDrivePermissionsResponseBody struct {
	CloudDrivePermissionModels []*DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels `json:"CloudDrivePermissionModels,omitempty" xml:"CloudDrivePermissionModels,omitempty" type:"Repeated"`
	// example:
	//
	// A87DBB05-653A-5E4B-B72B-5F4A1E07E5B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudDrivePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDrivePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudDrivePermissionsResponseBody) GetCloudDrivePermissionModels() []*DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels {
	return s.CloudDrivePermissionModels
}

func (s *DescribeCloudDrivePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudDrivePermissionsResponseBody) SetCloudDrivePermissionModels(v []*DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) *DescribeCloudDrivePermissionsResponseBody {
	s.CloudDrivePermissionModels = v
	return s
}

func (s *DescribeCloudDrivePermissionsResponseBody) SetRequestId(v string) *DescribeCloudDrivePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudDrivePermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels struct {
	EndUsers []*string `json:"EndUsers,omitempty" xml:"EndUsers,omitempty" type:"Repeated"`
	// example:
	//
	// CDS_DOWNLOAD
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) GoString() string {
	return s.String()
}

func (s *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) GetEndUsers() []*string {
	return s.EndUsers
}

func (s *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) GetPermission() *string {
	return s.Permission
}

func (s *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) SetEndUsers(v []*string) *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels {
	s.EndUsers = v
	return s
}

func (s *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) SetPermission(v string) *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels {
	s.Permission = &v
	return s
}

func (s *DescribeCloudDrivePermissionsResponseBodyCloudDrivePermissionModels) Validate() error {
	return dara.Validate(s)
}
