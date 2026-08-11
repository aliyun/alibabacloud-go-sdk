// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PublishScriptRequest
	GetInstanceId() *string
	SetScriptId(v string) *PublishScriptRequest
	GetScriptId() *string
	SetVersionId(v string) *PublishScriptRequest
	GetVersionId() *string
}

type PublishScriptRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b89
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The version ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b11
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s PublishScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishScriptRequest) GoString() string {
	return s.String()
}

func (s *PublishScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PublishScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *PublishScriptRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PublishScriptRequest) SetInstanceId(v string) *PublishScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *PublishScriptRequest) SetScriptId(v string) *PublishScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *PublishScriptRequest) SetVersionId(v string) *PublishScriptRequest {
	s.VersionId = &v
	return s
}

func (s *PublishScriptRequest) Validate() error {
	return dara.Validate(s)
}
