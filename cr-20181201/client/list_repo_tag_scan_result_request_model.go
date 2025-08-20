// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTagScanResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *ListRepoTagScanResultRequest
	GetDigest() *string
	SetFilterValue(v string) *ListRepoTagScanResultRequest
	GetFilterValue() *string
	SetInstanceId(v string) *ListRepoTagScanResultRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListRepoTagScanResultRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoTagScanResultRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListRepoTagScanResultRequest
	GetRepoId() *string
	SetScanTaskId(v string) *ListRepoTagScanResultRequest
	GetScanTaskId() *string
	SetScanType(v string) *ListRepoTagScanResultRequest
	GetScanType() *string
	SetSeverity(v string) *ListRepoTagScanResultRequest
	GetSeverity() *string
	SetTag(v string) *ListRepoTagScanResultRequest
	GetTag() *string
	SetVulQueryKey(v string) *ListRepoTagScanResultRequest
	GetVulQueryKey() *string
}

type ListRepoTagScanResultRequest struct {
	// The digest of the image.
	//
	// example:
	//
	// sha256:6b0b094f8a904f8fb6602427aed0d1fa
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The parameter whose value that you want to query. Fox example, if the value is `FixCmd`, only the `FixCmd` parameter is returned.
	//
	// example:
	//
	// FixCmd
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-2j88dtld8yel****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-uf082u9dg8do****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the security scan task.
	//
	// example:
	//
	// 6b0b094f-8a90-4f8f-b660-2427aed0****
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- `cve`: image system vulnerability
	//
	// 	- `sca`: image application vulnerability
	//
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The severity of the vulnerability. Valid values:
	//
	// 	- `High`
	//
	// 	- `Medium`
	//
	// 	- `Low`
	//
	// 	- `Unknown`
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The name of the image tag.
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The keyword for fuzzy search used in scanning. The value can be a CVE name.
	//
	// example:
	//
	// CVE-2021
	VulQueryKey *string `json:"VulQueryKey,omitempty" xml:"VulQueryKey,omitempty"`
}

func (s ListRepoTagScanResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagScanResultRequest) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultRequest) GetDigest() *string {
	return s.Digest
}

func (s *ListRepoTagScanResultRequest) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ListRepoTagScanResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoTagScanResultRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoTagScanResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoTagScanResultRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoTagScanResultRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListRepoTagScanResultRequest) GetScanType() *string {
	return s.ScanType
}

func (s *ListRepoTagScanResultRequest) GetSeverity() *string {
	return s.Severity
}

func (s *ListRepoTagScanResultRequest) GetTag() *string {
	return s.Tag
}

func (s *ListRepoTagScanResultRequest) GetVulQueryKey() *string {
	return s.VulQueryKey
}

func (s *ListRepoTagScanResultRequest) SetDigest(v string) *ListRepoTagScanResultRequest {
	s.Digest = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetFilterValue(v string) *ListRepoTagScanResultRequest {
	s.FilterValue = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetInstanceId(v string) *ListRepoTagScanResultRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetPageNo(v int32) *ListRepoTagScanResultRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetPageSize(v int32) *ListRepoTagScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetRepoId(v string) *ListRepoTagScanResultRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetScanTaskId(v string) *ListRepoTagScanResultRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetScanType(v string) *ListRepoTagScanResultRequest {
	s.ScanType = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetSeverity(v string) *ListRepoTagScanResultRequest {
	s.Severity = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetTag(v string) *ListRepoTagScanResultRequest {
	s.Tag = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetVulQueryKey(v string) *ListRepoTagScanResultRequest {
	s.VulQueryKey = &v
	return s
}

func (s *ListRepoTagScanResultRequest) Validate() error {
	return dara.Validate(s)
}
