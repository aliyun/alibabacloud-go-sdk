// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListScriptsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListScriptsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListScriptsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListScriptsRequest
	GetRegionId() *string
	SetScriptId(v string) *ListScriptsRequest
	GetScriptId() *string
	SetScriptName(v string) *ListScriptsRequest
	GetScriptName() *string
	SetScriptType(v string) *ListScriptsRequest
	GetScriptType() *string
	SetStatuses(v []*string) *ListScriptsRequest
	GetStatuses() []*string
}

type ListScriptsRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum number of records to retrieve at once.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Marks the current position to start reading from.
	//
	// example:
	//
	// dd6b1b2a-5837-5237-abe4-ff0c89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The script ID. Only common scripts are supported.
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the script. Only common scripts are supported. Fuzzy search is supported.
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Type of cluster script. Possible values:
	//
	// - BOOTSTRAP: Bootstrap script.
	//
	// - NORMAL: Regular cluster script.
	//
	// This parameter is required.
	//
	// example:
	//
	// BOOTSTRAP
	ScriptType *string `json:"ScriptType,omitempty" xml:"ScriptType,omitempty"`
	// The script status list.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListScriptsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListScriptsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScriptsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListScriptsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptsRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListScriptsRequest) GetScriptType() *string {
	return s.ScriptType
}

func (s *ListScriptsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListScriptsRequest) SetClusterId(v string) *ListScriptsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListScriptsRequest) SetMaxResults(v int32) *ListScriptsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScriptsRequest) SetNextToken(v string) *ListScriptsRequest {
	s.NextToken = &v
	return s
}

func (s *ListScriptsRequest) SetRegionId(v string) *ListScriptsRequest {
	s.RegionId = &v
	return s
}

func (s *ListScriptsRequest) SetScriptId(v string) *ListScriptsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListScriptsRequest) SetScriptName(v string) *ListScriptsRequest {
	s.ScriptName = &v
	return s
}

func (s *ListScriptsRequest) SetScriptType(v string) *ListScriptsRequest {
	s.ScriptType = &v
	return s
}

func (s *ListScriptsRequest) SetStatuses(v []*string) *ListScriptsRequest {
	s.Statuses = v
	return s
}

func (s *ListScriptsRequest) Validate() error {
	return dara.Validate(s)
}
