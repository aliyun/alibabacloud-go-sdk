// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteScriptRecordingRequest
	GetInstanceId() *string
	SetScriptId(v string) *DeleteScriptRecordingRequest
	GetScriptId() *string
	SetUuidsJson(v string) *DeleteScriptRecordingRequest
	GetUuidsJson() *string
}

type DeleteScriptRecordingRequest struct {
	// Instance ID to which the recording belongs
	//
	// This parameter is required.
	//
	// example:
	//
	// 1eefcb81-cd58-4143-8180-6a962d79d708
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID to which the recording belongs
	//
	// This parameter is required.
	//
	// example:
	//
	// dcc42f0d-cfd8-4866-9bbf-002042503745
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// List of recording IDs. If empty, all recordings are selected by default.
	//
	// > You can obtain the recording IDs from the ListScriptRecording API.
	//
	// example:
	//
	// ["d17d5bfa-4972-4389-9718-f9602edabe48"]
	UuidsJson *string `json:"UuidsJson,omitempty" xml:"UuidsJson,omitempty"`
}

func (s DeleteScriptRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptRecordingRequest) GoString() string {
	return s.String()
}

func (s *DeleteScriptRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteScriptRecordingRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DeleteScriptRecordingRequest) GetUuidsJson() *string {
	return s.UuidsJson
}

func (s *DeleteScriptRecordingRequest) SetInstanceId(v string) *DeleteScriptRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScriptRecordingRequest) SetScriptId(v string) *DeleteScriptRecordingRequest {
	s.ScriptId = &v
	return s
}

func (s *DeleteScriptRecordingRequest) SetUuidsJson(v string) *DeleteScriptRecordingRequest {
	s.UuidsJson = &v
	return s
}

func (s *DeleteScriptRecordingRequest) Validate() error {
	return dara.Validate(s)
}
