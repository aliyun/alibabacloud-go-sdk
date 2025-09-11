// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppPatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListCloudAppPatchesRequest
	GetAppId() *string
	SetEndTime(v string) *ListCloudAppPatchesRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListCloudAppPatchesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppPatchesRequest
	GetPageSize() *int64
	SetPatchId(v string) *ListCloudAppPatchesRequest
	GetPatchId() *string
	SetPatchName(v string) *ListCloudAppPatchesRequest
	GetPatchName() *string
	SetStartTime(v string) *ListCloudAppPatchesRequest
	GetStartTime() *string
}

type ListCloudAppPatchesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// 2015-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCloudAppPatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppPatchesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAppPatchesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListCloudAppPatchesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListCloudAppPatchesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppPatchesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppPatchesRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *ListCloudAppPatchesRequest) GetPatchName() *string {
	return s.PatchName
}

func (s *ListCloudAppPatchesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCloudAppPatchesRequest) SetAppId(v string) *ListCloudAppPatchesRequest {
	s.AppId = &v
	return s
}

func (s *ListCloudAppPatchesRequest) SetEndTime(v string) *ListCloudAppPatchesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudAppPatchesRequest) SetPageNumber(v int64) *ListCloudAppPatchesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppPatchesRequest) SetPageSize(v int64) *ListCloudAppPatchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppPatchesRequest) SetPatchId(v string) *ListCloudAppPatchesRequest {
	s.PatchId = &v
	return s
}

func (s *ListCloudAppPatchesRequest) SetPatchName(v string) *ListCloudAppPatchesRequest {
	s.PatchName = &v
	return s
}

func (s *ListCloudAppPatchesRequest) SetStartTime(v string) *ListCloudAppPatchesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudAppPatchesRequest) Validate() error {
	return dara.Validate(s)
}
