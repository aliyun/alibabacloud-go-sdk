// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDuplicateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DuplicateScriptRequest
	GetInstanceId() *string
	SetName(v string) *DuplicateScriptRequest
	GetName() *string
	SetSourceScriptId(v string) *DuplicateScriptRequest
	GetSourceScriptId() *string
}

type DuplicateScriptRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// New scenario name
	//
	// This parameter is required.
	//
	// example:
	//
	// 复制的催收话术
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Source scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6114e7e8-4140-48d9-b46d-65ea29f13fe8
	SourceScriptId *string `json:"SourceScriptId,omitempty" xml:"SourceScriptId,omitempty"`
}

func (s DuplicateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s DuplicateScriptRequest) GoString() string {
	return s.String()
}

func (s *DuplicateScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DuplicateScriptRequest) GetName() *string {
	return s.Name
}

func (s *DuplicateScriptRequest) GetSourceScriptId() *string {
	return s.SourceScriptId
}

func (s *DuplicateScriptRequest) SetInstanceId(v string) *DuplicateScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *DuplicateScriptRequest) SetName(v string) *DuplicateScriptRequest {
	s.Name = &v
	return s
}

func (s *DuplicateScriptRequest) SetSourceScriptId(v string) *DuplicateScriptRequest {
	s.SourceScriptId = &v
	return s
}

func (s *DuplicateScriptRequest) Validate() error {
	return dara.Validate(s)
}
