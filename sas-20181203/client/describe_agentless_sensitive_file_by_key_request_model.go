// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentlessSensitiveFileByKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAgentlessSensitiveFileByKeyRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeAgentlessSensitiveFileByKeyRequest
	GetImageUuid() *string
	SetInstanceId(v string) *DescribeAgentlessSensitiveFileByKeyRequest
	GetInstanceId() *string
	SetPageSize(v int32) *DescribeAgentlessSensitiveFileByKeyRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeAgentlessSensitiveFileByKeyRequest
	GetRemark() *string
	SetScanRange(v []*string) *DescribeAgentlessSensitiveFileByKeyRequest
	GetScanRange() []*string
	SetSensitiveFileKey(v string) *DescribeAgentlessSensitiveFileByKeyRequest
	GetSensitiveFileKey() *string
	SetStatus(v string) *DescribeAgentlessSensitiveFileByKeyRequest
	GetStatus() *string
}

type DescribeAgentlessSensitiveFileByKeyRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the asset that is scanned. You can query the UUID on the Host or Cloud Product page. If you scan a host, set this parameter to the UUID of the scanned host. If you scan a snapshot or a custom image, set this parameter to the ID of the scanned snapshot or image.
	//
	// example:
	//
	// 06293273b67d19516cfcc712194f****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The instance ID of the asset that is scanned. To query the instance ID, go to the Task Management page, click Details of a task, and then view the value of Check On.
	//
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the asset.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The types of the assets that are scanned.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The type of the sensitive file.
	//
	// example:
	//
	// sshpasswd
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The status of the baseline risk. Valid values:
	//
	// 	- **0**: unfixed.
	//
	// 	- **1**: fixed.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAgentlessSensitiveFileByKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentlessSensitiveFileByKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetCurrentPage(v int32) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetImageUuid(v string) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetInstanceId(v string) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetPageSize(v int32) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetRemark(v string) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.Remark = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetScanRange(v []*string) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetSensitiveFileKey(v string) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) SetStatus(v string) *DescribeAgentlessSensitiveFileByKeyRequest {
	s.Status = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyRequest) Validate() error {
	return dara.Validate(s)
}
