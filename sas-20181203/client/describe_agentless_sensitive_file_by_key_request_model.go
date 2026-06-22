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
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
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
