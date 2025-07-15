// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSmsMetadataRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListSmsMetadataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSmsMetadataRequest
	GetPageSize() *int32
	SetScenarioListJson(v string) *ListSmsMetadataRequest
	GetScenarioListJson() *string
}

type ListSmsMetadataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
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
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ["SATISFACTION_SURVEY","CONTACT_FLOW_NODE"]
	ScenarioListJson *string `json:"ScenarioListJson,omitempty" xml:"ScenarioListJson,omitempty"`
}

func (s ListSmsMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmsMetadataRequest) GoString() string {
	return s.String()
}

func (s *ListSmsMetadataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSmsMetadataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSmsMetadataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSmsMetadataRequest) GetScenarioListJson() *string {
	return s.ScenarioListJson
}

func (s *ListSmsMetadataRequest) SetInstanceId(v string) *ListSmsMetadataRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSmsMetadataRequest) SetPageNumber(v int32) *ListSmsMetadataRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSmsMetadataRequest) SetPageSize(v int32) *ListSmsMetadataRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmsMetadataRequest) SetScenarioListJson(v string) *ListSmsMetadataRequest {
	s.ScenarioListJson = &v
	return s
}

func (s *ListSmsMetadataRequest) Validate() error {
	return dara.Validate(s)
}
