// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeStorageVolumeResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeStorageVolumeResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeStorageVolumeResponseBody
	GetRequestId() *string
	SetStorageVolumes(v []*DescribeStorageVolumeResponseBodyStorageVolumes) *DescribeStorageVolumeResponseBody
	GetStorageVolumes() []*DescribeStorageVolumeResponseBodyStorageVolumes
	SetTotalCount(v string) *DescribeStorageVolumeResponseBody
	GetTotalCount() *string
}

type DescribeStorageVolumeResponseBody struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of returned results.
	StorageVolumes []*DescribeStorageVolumeResponseBodyStorageVolumes `json:"StorageVolumes,omitempty" xml:"StorageVolumes,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStorageVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageVolumeResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeStorageVolumeResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeStorageVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageVolumeResponseBody) GetStorageVolumes() []*DescribeStorageVolumeResponseBodyStorageVolumes {
	return s.StorageVolumes
}

func (s *DescribeStorageVolumeResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeStorageVolumeResponseBody) SetPageNumber(v string) *DescribeStorageVolumeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageVolumeResponseBody) SetPageSize(v string) *DescribeStorageVolumeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageVolumeResponseBody) SetRequestId(v string) *DescribeStorageVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageVolumeResponseBody) SetStorageVolumes(v []*DescribeStorageVolumeResponseBodyStorageVolumes) *DescribeStorageVolumeResponseBody {
	s.StorageVolumes = v
	return s
}

func (s *DescribeStorageVolumeResponseBody) SetTotalCount(v string) *DescribeStorageVolumeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStorageVolumeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStorageVolumeResponseBodyStorageVolumes struct {
	// The authentication protocol. The value is set to **CHAP**.
	//
	// example:
	//
	// CHAP
	AuthProtocol *string `json:"AuthProtocol,omitempty" xml:"AuthProtocol,omitempty"`
	// The time when the volume was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-03-14T09:35:32Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the volume.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-shenzhen-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// Indicates whether authentication is enabled. Valid values:
	//
	// 	- **1**: Authentication is enabled.
	//
	// 	- **0*	- (default): Authentication is disabled.
	//
	// example:
	//
	// 0
	IsAuth *int32 `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// Indicates whether the volume is enabled. Valid values:
	//
	// 	- **1*	- (default): The volume is enabled.
	//
	// 	- **0**: The volume is disabled.
	//
	// example:
	//
	// 1
	IsEnable *int32 `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The status of the volume. Valid values:
	//
	// 	- creating
	//
	// 	- available
	//
	// 	- deleting
	//
	// 	- deleted
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the storage gateway.
	//
	// example:
	//
	// sgw-***
	StorageGatewayId *string `json:"StorageGatewayId,omitempty" xml:"StorageGatewayId,omitempty"`
	// The ID of the storage medium.
	//
	// example:
	//
	// d-***
	StorageId *string `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
	// The ID of the volume.
	//
	// example:
	//
	// sv-***
	StorageVolumeId *string `json:"StorageVolumeId,omitempty" xml:"StorageVolumeId,omitempty"`
	// The name of the volume.
	//
	// example:
	//
	// testVolumeName
	StorageVolumeName *string `json:"StorageVolumeName,omitempty" xml:"StorageVolumeName,omitempty"`
	// The destination of the volume.
	//
	// example:
	//
	// iqn.*.*.*:*
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
}

func (s DescribeStorageVolumeResponseBodyStorageVolumes) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageVolumeResponseBodyStorageVolumes) GoString() string {
	return s.String()
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetAuthProtocol() *string {
	return s.AuthProtocol
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetDescription() *string {
	return s.Description
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetIsAuth() *int32 {
	return s.IsAuth
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetIsEnable() *int32 {
	return s.IsEnable
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetStatus() *string {
	return s.Status
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetStorageGatewayId() *string {
	return s.StorageGatewayId
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetStorageId() *string {
	return s.StorageId
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetStorageVolumeId() *string {
	return s.StorageVolumeId
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetStorageVolumeName() *string {
	return s.StorageVolumeName
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetAuthProtocol(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.AuthProtocol = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetCreationTime(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.CreationTime = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetDescription(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.Description = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetEnsRegionId(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetIsAuth(v int32) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.IsAuth = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetIsEnable(v int32) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.IsEnable = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetStatus(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.Status = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetStorageGatewayId(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.StorageGatewayId = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetStorageId(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.StorageId = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetStorageVolumeId(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.StorageVolumeId = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetStorageVolumeName(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.StorageVolumeName = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) SetTargetName(v string) *DescribeStorageVolumeResponseBodyStorageVolumes {
	s.TargetName = &v
	return s
}

func (s *DescribeStorageVolumeResponseBodyStorageVolumes) Validate() error {
	return dara.Validate(s)
}
