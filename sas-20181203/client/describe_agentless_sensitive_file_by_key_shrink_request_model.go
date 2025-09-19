// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentlessSensitiveFileByKeyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetImageUuid() *string
	SetInstanceId(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetRemark() *string
	SetScanRangeShrink(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetScanRangeShrink() *string
	SetSensitiveFileKey(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetSensitiveFileKey() *string
	SetStatus(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest
	GetStatus() *string
}

type DescribeAgentlessSensitiveFileByKeyShrinkRequest struct {
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
	ScanRangeShrink *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
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

func (s DescribeAgentlessSensitiveFileByKeyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentlessSensitiveFileByKeyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetScanRangeShrink() *string {
	return s.ScanRangeShrink
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetCurrentPage(v int32) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetImageUuid(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetInstanceId(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetPageSize(v int32) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetRemark(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.Remark = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetScanRangeShrink(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.ScanRangeShrink = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetSensitiveFileKey(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) SetStatus(v string) *DescribeAgentlessSensitiveFileByKeyShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
