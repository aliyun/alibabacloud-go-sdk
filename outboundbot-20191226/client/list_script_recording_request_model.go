// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptRecordingRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListScriptRecordingRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptRecordingRequest
	GetPageSize() *int32
	SetRefIdsJson(v string) *ListScriptRecordingRequest
	GetRefIdsJson() *string
	SetScriptId(v string) *ListScriptRecordingRequest
	GetScriptId() *string
	SetSearch(v string) *ListScriptRecordingRequest
	GetSearch() *string
	SetStatesJson(v string) *ListScriptRecordingRequest
	GetStatesJson() *string
	SetUuidsJson(v string) *ListScriptRecordingRequest
	GetUuidsJson() *string
}

type ListScriptRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c209abb3-6804-4a75-b2c7-dd55c8c61b6a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RefIdsJson *string `json:"RefIdsJson,omitempty" xml:"RefIdsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	Search   *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// [
	//
	//      "8"
	//
	// ]
	StatesJson *string `json:"StatesJson,omitempty" xml:"StatesJson,omitempty"`
	// example:
	//
	// ["d17d5bfa-4972-4389-9718-f9602edabe48"]
	UuidsJson *string `json:"UuidsJson,omitempty" xml:"UuidsJson,omitempty"`
}

func (s ListScriptRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptRecordingRequest) GoString() string {
	return s.String()
}

func (s *ListScriptRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptRecordingRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptRecordingRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptRecordingRequest) GetRefIdsJson() *string {
	return s.RefIdsJson
}

func (s *ListScriptRecordingRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptRecordingRequest) GetSearch() *string {
	return s.Search
}

func (s *ListScriptRecordingRequest) GetStatesJson() *string {
	return s.StatesJson
}

func (s *ListScriptRecordingRequest) GetUuidsJson() *string {
	return s.UuidsJson
}

func (s *ListScriptRecordingRequest) SetInstanceId(v string) *ListScriptRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptRecordingRequest) SetPageNumber(v int32) *ListScriptRecordingRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptRecordingRequest) SetPageSize(v int32) *ListScriptRecordingRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptRecordingRequest) SetRefIdsJson(v string) *ListScriptRecordingRequest {
	s.RefIdsJson = &v
	return s
}

func (s *ListScriptRecordingRequest) SetScriptId(v string) *ListScriptRecordingRequest {
	s.ScriptId = &v
	return s
}

func (s *ListScriptRecordingRequest) SetSearch(v string) *ListScriptRecordingRequest {
	s.Search = &v
	return s
}

func (s *ListScriptRecordingRequest) SetStatesJson(v string) *ListScriptRecordingRequest {
	s.StatesJson = &v
	return s
}

func (s *ListScriptRecordingRequest) SetUuidsJson(v string) *ListScriptRecordingRequest {
	s.UuidsJson = &v
	return s
}

func (s *ListScriptRecordingRequest) Validate() error {
	return dara.Validate(s)
}
