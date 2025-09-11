// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppPatchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListCloudAppPatchesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppPatchesResponseBody
	GetPageSize() *int64
	SetPatches(v []*ListCloudAppPatchesResponseBodyPatches) *ListCloudAppPatchesResponseBody
	GetPatches() []*ListCloudAppPatchesResponseBodyPatches
	SetRequestId(v string) *ListCloudAppPatchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCloudAppPatchesResponseBody
	GetTotalCount() *int64
}

type ListCloudAppPatchesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Patches  []*ListCloudAppPatchesResponseBodyPatches `json:"Patches,omitempty" xml:"Patches,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAppPatchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppPatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAppPatchesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppPatchesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppPatchesResponseBody) GetPatches() []*ListCloudAppPatchesResponseBodyPatches {
	return s.Patches
}

func (s *ListCloudAppPatchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAppPatchesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudAppPatchesResponseBody) SetPageNumber(v int64) *ListCloudAppPatchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppPatchesResponseBody) SetPageSize(v int64) *ListCloudAppPatchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppPatchesResponseBody) SetPatches(v []*ListCloudAppPatchesResponseBodyPatches) *ListCloudAppPatchesResponseBody {
	s.Patches = v
	return s
}

func (s *ListCloudAppPatchesResponseBody) SetRequestId(v string) *ListCloudAppPatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAppPatchesResponseBody) SetTotalCount(v int64) *ListCloudAppPatchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAppPatchesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudAppPatchesResponseBodyPatches struct {
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// example:
	//
	// patch-1
	PatchName *string `json:"PatchName,omitempty" xml:"PatchName,omitempty"`
	// example:
	//
	// Doing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Uploading
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// example:
	//
	// 2024-09-23T02:12:28
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 2025-07-24T14:45:36+08:00
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s ListCloudAppPatchesResponseBodyPatches) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppPatchesResponseBodyPatches) GoString() string {
	return s.String()
}

func (s *ListCloudAppPatchesResponseBodyPatches) GetPatchId() *string {
	return s.PatchId
}

func (s *ListCloudAppPatchesResponseBodyPatches) GetPatchName() *string {
	return s.PatchName
}

func (s *ListCloudAppPatchesResponseBodyPatches) GetStatus() *string {
	return s.Status
}

func (s *ListCloudAppPatchesResponseBodyPatches) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *ListCloudAppPatchesResponseBodyPatches) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudAppPatchesResponseBodyPatches) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListCloudAppPatchesResponseBodyPatches) SetPatchId(v string) *ListCloudAppPatchesResponseBodyPatches {
	s.PatchId = &v
	return s
}

func (s *ListCloudAppPatchesResponseBodyPatches) SetPatchName(v string) *ListCloudAppPatchesResponseBodyPatches {
	s.PatchName = &v
	return s
}

func (s *ListCloudAppPatchesResponseBodyPatches) SetStatus(v string) *ListCloudAppPatchesResponseBodyPatches {
	s.Status = &v
	return s
}

func (s *ListCloudAppPatchesResponseBodyPatches) SetStatusDescription(v string) *ListCloudAppPatchesResponseBodyPatches {
	s.StatusDescription = &v
	return s
}

func (s *ListCloudAppPatchesResponseBodyPatches) SetUpdateTime(v string) *ListCloudAppPatchesResponseBodyPatches {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAppPatchesResponseBodyPatches) SetUploadTime(v string) *ListCloudAppPatchesResponseBodyPatches {
	s.UploadTime = &v
	return s
}

func (s *ListCloudAppPatchesResponseBodyPatches) Validate() error {
	return dara.Validate(s)
}
