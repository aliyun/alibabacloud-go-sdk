// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePatchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListInstancePatchesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancePatchesResponseBody
	GetNextToken() *string
	SetPatches(v []*ListInstancePatchesResponseBodyPatches) *ListInstancePatchesResponseBody
	GetPatches() []*ListInstancePatchesResponseBodyPatches
	SetRequestId(v string) *ListInstancePatchesResponseBody
	GetRequestId() *string
}

type ListInstancePatchesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the patch.
	Patches []*ListInstancePatchesResponseBodyPatches `json:"Patches,omitempty" xml:"Patches,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0A615755-9C86-5EA6-BF9E-6E8F1AFF9403
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePatchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancePatchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancePatchesResponseBody) GetPatches() []*ListInstancePatchesResponseBodyPatches {
	return s.Patches
}

func (s *ListInstancePatchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancePatchesResponseBody) SetMaxResults(v int32) *ListInstancePatchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetNextToken(v string) *ListInstancePatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetPatches(v []*ListInstancePatchesResponseBodyPatches) *ListInstancePatchesResponseBody {
	s.Patches = v
	return s
}

func (s *ListInstancePatchesResponseBody) SetRequestId(v string) *ListInstancePatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancePatchesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancePatchesResponseBodyPatches struct {
	// The classification of the patch.
	//
	// example:
	//
	// “”
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// The time when the patch was installed.
	//
	// example:
	//
	// 2021-01-28T07:07:20Z
	InstalledTime *string `json:"InstalledTime,omitempty" xml:"InstalledTime,omitempty"`
	// The Id of KBId.
	//
	// example:
	//
	// apt-utils.amd64
	KBId *string `json:"KBId,omitempty" xml:"KBId,omitempty"`
	// The level of the severity.
	//
	// example:
	//
	// important
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The status of the installation.
	//
	// example:
	//
	// Installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the patch.
	//
	// example:
	//
	// isc-dhcp-common.amd64:4.3.5-3ubuntu7.3
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListInstancePatchesResponseBodyPatches) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchesResponseBodyPatches) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponseBodyPatches) GetClassification() *string {
	return s.Classification
}

func (s *ListInstancePatchesResponseBodyPatches) GetInstalledTime() *string {
	return s.InstalledTime
}

func (s *ListInstancePatchesResponseBodyPatches) GetKBId() *string {
	return s.KBId
}

func (s *ListInstancePatchesResponseBodyPatches) GetSeverity() *string {
	return s.Severity
}

func (s *ListInstancePatchesResponseBodyPatches) GetStatus() *string {
	return s.Status
}

func (s *ListInstancePatchesResponseBodyPatches) GetTitle() *string {
	return s.Title
}

func (s *ListInstancePatchesResponseBodyPatches) SetClassification(v string) *ListInstancePatchesResponseBodyPatches {
	s.Classification = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetInstalledTime(v string) *ListInstancePatchesResponseBodyPatches {
	s.InstalledTime = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetKBId(v string) *ListInstancePatchesResponseBodyPatches {
	s.KBId = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetSeverity(v string) *ListInstancePatchesResponseBodyPatches {
	s.Severity = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetStatus(v string) *ListInstancePatchesResponseBodyPatches {
	s.Status = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetTitle(v string) *ListInstancePatchesResponseBodyPatches {
	s.Title = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) Validate() error {
	return dara.Validate(s)
}
