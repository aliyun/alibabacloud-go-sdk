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
	// Instance ID to which the recording belongs
	//
	// This parameter is required.
	//
	// example:
	//
	// c209abb3-6804-4a75-b2c7-dd55c8c61b6a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries displayed per page
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of recording IDs (in JSON format), identical to UuidsJson
	//
	// example:
	//
	// ["d9fad189-760b-47b9-837f-aeabb4fc9109\\"]
	RefIdsJson *string `json:"RefIdsJson,omitempty" xml:"RefIdsJson,omitempty"`
	// Scenario ID to which the recording belongs
	//
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Search criteria: you can query by recording Name or Content
	//
	// example:
	//
	// 您好
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// Recording status:
	//
	// - Validating: 1
	//
	// - Validation failed: 2
	//
	// - Processing: 3
	//
	// - Processing failed: 4
	//
	// - Reviewing: 5
	//
	// - Review failed: 6
	//
	// - Publish failed: 7
	//
	// - Published: 8
	//
	// - Pending publish: 9
	//
	// example:
	//
	// [
	//
	//      "8"
	//
	// ]
	StatesJson *string `json:"StatesJson,omitempty" xml:"StatesJson,omitempty"`
	// List of recording IDs
	//
	// > That is, a list of Unique recording IDs. If not specified, all recordings are displayed.
	//
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
