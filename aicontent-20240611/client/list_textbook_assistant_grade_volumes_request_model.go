// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantGradeVolumesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *ListTextbookAssistantGradeVolumesRequest
	GetAuthToken() *string
	SetScenario(v string) *ListTextbookAssistantGradeVolumesRequest
	GetScenario() *string
}

type ListTextbookAssistantGradeVolumesRequest struct {
	// example:
	//
	// tc_197bf5bb81889cc79eb51ae9b8c0cea3
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListTextbookAssistantGradeVolumesRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ListTextbookAssistantGradeVolumesRequest) SetAuthToken(v string) *ListTextbookAssistantGradeVolumesRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesRequest) SetScenario(v string) *ListTextbookAssistantGradeVolumesRequest {
	s.Scenario = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesRequest) Validate() error {
	return dara.Validate(s)
}
