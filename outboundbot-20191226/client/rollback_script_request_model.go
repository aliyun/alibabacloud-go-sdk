// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RollbackScriptRequest
	GetInstanceId() *string
	SetRollbackVersion(v string) *RollbackScriptRequest
	GetRollbackVersion() *string
	SetScriptId(v string) *RollbackScriptRequest
	GetScriptId() *string
}

type RollbackScriptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1579055782000
	RollbackVersion *string `json:"RollbackVersion,omitempty" xml:"RollbackVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 303523b1-0094-4ebe-b9ed-c23d11c91d61
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s RollbackScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackScriptRequest) GoString() string {
	return s.String()
}

func (s *RollbackScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RollbackScriptRequest) GetRollbackVersion() *string {
	return s.RollbackVersion
}

func (s *RollbackScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *RollbackScriptRequest) SetInstanceId(v string) *RollbackScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *RollbackScriptRequest) SetRollbackVersion(v string) *RollbackScriptRequest {
	s.RollbackVersion = &v
	return s
}

func (s *RollbackScriptRequest) SetScriptId(v string) *RollbackScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *RollbackScriptRequest) Validate() error {
	return dara.Validate(s)
}
