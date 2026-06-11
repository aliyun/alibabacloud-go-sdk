// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnnotationMissionSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotationMissionId(v string) *ListAnnotationMissionSessionRequest
	GetAnnotationMissionId() *string
	SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionRequest
	GetAnnotationMissionSessionId() *string
	SetEnvironment(v int32) *ListAnnotationMissionSessionRequest
	GetEnvironment() *int32
	SetIncludeStatusListJsonString(v string) *ListAnnotationMissionSessionRequest
	GetIncludeStatusListJsonString() *string
	SetPageIndex(v int32) *ListAnnotationMissionSessionRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListAnnotationMissionSessionRequest
	GetPageSize() *int32
}

type ListAnnotationMissionSessionRequest struct {
	// The task ID.
	//
	// example:
	//
	// 8434a4b0-41fc-41b1-aa75-bbd1f2ab0c8d
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// The ID of the chat instance for the annotation task.
	//
	// > Obtain this ID by calling the SaveAnnotationMissionSessionList operation.
	//
	// example:
	//
	// 8434a4b0-41fc-41b1-aa75-bbd1f2ab0c8d
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// The environment.
	//
	// > The default value is 2.
	//
	// **Enumeration values**
	//
	// - 0: NONE
	//
	// - 1: Private cloud
	//
	// - 2: Public cloud
	//
	// example:
	//
	// 0
	Environment *int32 `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The list of statuses to include.
	//
	// > The format is [1]. Fill the array with specific enumeration values.
	//
	// **Enumeration values**
	//
	// - 1: Correctly detected
	//
	// - 2: Incorrectly detected
	//
	// - 8: Invalid data
	//
	// - 4: Uncovered intent
	//
	// - 16: Transcription error
	//
	// example:
	//
	// [1]
	IncludeStatusListJsonString *string `json:"IncludeStatusListJsonString,omitempty" xml:"IncludeStatusListJsonString,omitempty"`
	// The page number.
	//
	// example:
	//
	// 3
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of records to display on each page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAnnotationMissionSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionSessionRequest) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionSessionRequest) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionSessionRequest) GetAnnotationMissionSessionId() *string {
	return s.AnnotationMissionSessionId
}

func (s *ListAnnotationMissionSessionRequest) GetEnvironment() *int32 {
	return s.Environment
}

func (s *ListAnnotationMissionSessionRequest) GetIncludeStatusListJsonString() *string {
	return s.IncludeStatusListJsonString
}

func (s *ListAnnotationMissionSessionRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListAnnotationMissionSessionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnnotationMissionSessionRequest) SetAnnotationMissionId(v string) *ListAnnotationMissionSessionRequest {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionSessionRequest) SetAnnotationMissionSessionId(v string) *ListAnnotationMissionSessionRequest {
	s.AnnotationMissionSessionId = &v
	return s
}

func (s *ListAnnotationMissionSessionRequest) SetEnvironment(v int32) *ListAnnotationMissionSessionRequest {
	s.Environment = &v
	return s
}

func (s *ListAnnotationMissionSessionRequest) SetIncludeStatusListJsonString(v string) *ListAnnotationMissionSessionRequest {
	s.IncludeStatusListJsonString = &v
	return s
}

func (s *ListAnnotationMissionSessionRequest) SetPageIndex(v int32) *ListAnnotationMissionSessionRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAnnotationMissionSessionRequest) SetPageSize(v int32) *ListAnnotationMissionSessionRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnnotationMissionSessionRequest) Validate() error {
	return dara.Validate(s)
}
