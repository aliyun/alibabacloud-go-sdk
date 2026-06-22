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
	// The page number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the asset to scan. You can obtain this value from the Host Assets and Cloud Products page. In host security detection scenarios, this parameter specifies the UUID of the host to scan. In snapshot security detection and custom image security detection scenarios, this parameter specifies the ID of the image or snapshot to scan.
	//
	// example:
	//
	// 06293273b67d19516cfcc712194f****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The instance ID of the asset to query. You can obtain this value from Node Management > Details > Detection Objects.
	//
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page in a paging query. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The asset name or IP address.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The collection of scan ranges.
	ScanRangeShrink *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The type of the sensitive file.
	//
	// example:
	//
	// sshpasswd
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The fix status of the baseline risk. Valid values:
	//
	// - **0**: unfixed
	//
	// - **1**: fixed.
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
