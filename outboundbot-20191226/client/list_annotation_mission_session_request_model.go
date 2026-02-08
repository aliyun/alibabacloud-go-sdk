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
	// Job ID
	//
	// example:
	//
	// 8434a4b0-41fc-41b1-aa75-bbd1f2ab0c8d
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// Annotation Job chat instance ID
	//
	// > Obtain it by calling the SaveAnnotationMissionSessionList API.
	//
	// example:
	//
	// 8434a4b0-41fc-41b1-aa75-bbd1f2ab0c8d
	AnnotationMissionSessionId *string `json:"AnnotationMissionSessionId,omitempty" xml:"AnnotationMissionSessionId,omitempty"`
	// Environment
	//
	// > Default value is 2.
	//
	// **Enumeration values**
	//
	// - 0: NONE
	//
	// - 1: Private Cloud
	//
	// - 2: Public Cloud
	//
	// example:
	//
	// 0
	Environment *int32 `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// List of included statuses
	//
	// > Format: [1]. Specify the exact enumeration values in the array.
	//
	// **Enumeration values*	-
	//
	// - 1: Detected correctly
	//
	// - 2: Detection fault
	//
	// - 8: Invalid data
	//
	// - 4: Intent not covered
	//
	// - 16: Misrecognized translation
	//
	// example:
	//
	// [1]
	IncludeStatusListJsonString *string `json:"IncludeStatusListJsonString,omitempty" xml:"IncludeStatusListJsonString,omitempty"`
	// Page index
	//
	// example:
	//
	// 3
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of records displayed per page.
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
