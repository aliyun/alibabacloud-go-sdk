// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptsRequest
	GetInstanceId() *string
	SetName(v string) *ListScriptsRequest
	GetName() *string
	SetPageNumber(v int32) *ListScriptsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptsRequest
	GetPageSize() *int32
	SetPublishOnly(v bool) *ListScriptsRequest
	GetPublishOnly() *bool
	SetScriptIds(v []*string) *ListScriptsRequest
	GetScriptIds() []*string
}

type ListScriptsRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// Satisfaction Survey
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return only published scenarios.
	//
	// example:
	//
	// true
	PublishOnly *bool `json:"PublishOnly,omitempty" xml:"PublishOnly,omitempty"`
	// The list of scenario IDs.
	ScriptIds []*string `json:"ScriptIds,omitempty" xml:"ScriptIds,omitempty" type:"Repeated"`
}

func (s ListScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptsRequest) GetName() *string {
	return s.Name
}

func (s *ListScriptsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsRequest) GetPublishOnly() *bool {
	return s.PublishOnly
}

func (s *ListScriptsRequest) GetScriptIds() []*string {
	return s.ScriptIds
}

func (s *ListScriptsRequest) SetInstanceId(v string) *ListScriptsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptsRequest) SetName(v string) *ListScriptsRequest {
	s.Name = &v
	return s
}

func (s *ListScriptsRequest) SetPageNumber(v int32) *ListScriptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsRequest) SetPageSize(v int32) *ListScriptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptsRequest) SetPublishOnly(v bool) *ListScriptsRequest {
	s.PublishOnly = &v
	return s
}

func (s *ListScriptsRequest) SetScriptIds(v []*string) *ListScriptsRequest {
	s.ScriptIds = v
	return s
}

func (s *ListScriptsRequest) Validate() error {
	return dara.Validate(s)
}
