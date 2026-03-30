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
	// example:
	//
	// 19ac2375-53e3-477f-abe9-6cd334227981
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b091
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
