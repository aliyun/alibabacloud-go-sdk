// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadScriptRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DownloadScriptRecordingRequest
	GetInstanceId() *string
	SetScriptId(v string) *DownloadScriptRecordingRequest
	GetScriptId() *string
	SetUuid(v string) *DownloadScriptRecordingRequest
	GetUuid() *string
}

type DownloadScriptRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b3865dc3-40fa-4afd-9fe4-dc7cda305a24
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5feaab8a-97fd-4720-8108-79e017f2d3ac
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DownloadScriptRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadScriptRecordingRequest) GoString() string {
	return s.String()
}

func (s *DownloadScriptRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DownloadScriptRecordingRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DownloadScriptRecordingRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DownloadScriptRecordingRequest) SetInstanceId(v string) *DownloadScriptRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *DownloadScriptRecordingRequest) SetScriptId(v string) *DownloadScriptRecordingRequest {
	s.ScriptId = &v
	return s
}

func (s *DownloadScriptRecordingRequest) SetUuid(v string) *DownloadScriptRecordingRequest {
	s.Uuid = &v
	return s
}

func (s *DownloadScriptRecordingRequest) Validate() error {
	return dara.Validate(s)
}
