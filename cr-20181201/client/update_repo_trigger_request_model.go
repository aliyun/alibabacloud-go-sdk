// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateRepoTriggerRequest
	GetInstanceId() *string
	SetRepoId(v string) *UpdateRepoTriggerRequest
	GetRepoId() *string
	SetTriggerId(v string) *UpdateRepoTriggerRequest
	GetTriggerId() *string
	SetTriggerName(v string) *UpdateRepoTriggerRequest
	GetTriggerName() *string
	SetTriggerTag(v string) *UpdateRepoTriggerRequest
	GetTriggerTag() *string
	SetTriggerType(v string) *UpdateRepoTriggerRequest
	GetTriggerType() *string
	SetTriggerUrl(v string) *UpdateRepoTriggerRequest
	GetTriggerUrl() *string
}

type UpdateRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// crw-k7bdx4kt52ty****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
	// The name of the trigger.
	//
	// You can specify the TriggerName or TriggerUrl parameter. The TriggerName parameter is optional.
	//
	// example:
	//
	// test_trigger
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// example:
	//
	// master
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- `ALL`: a trigger that supports both tags and regular expressions.
	//
	// 	- `TAG_LISTTAG`: a tag-based trigger.
	//
	// 	- `TAG_REG_EXP`: a regular expression-based trigger.
	//
	// example:
	//
	// TAG_LIST
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
	//
	// example:
	//
	// https://www.test.com
	TriggerUrl *string `json:"TriggerUrl,omitempty" xml:"TriggerUrl,omitempty"`
}

func (s UpdateRepoTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepoTriggerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRepoTriggerRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *UpdateRepoTriggerRequest) GetTriggerId() *string {
	return s.TriggerId
}

func (s *UpdateRepoTriggerRequest) GetTriggerName() *string {
	return s.TriggerName
}

func (s *UpdateRepoTriggerRequest) GetTriggerTag() *string {
	return s.TriggerTag
}

func (s *UpdateRepoTriggerRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateRepoTriggerRequest) GetTriggerUrl() *string {
	return s.TriggerUrl
}

func (s *UpdateRepoTriggerRequest) SetInstanceId(v string) *UpdateRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetRepoId(v string) *UpdateRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerId(v string) *UpdateRepoTriggerRequest {
	s.TriggerId = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerName(v string) *UpdateRepoTriggerRequest {
	s.TriggerName = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerTag(v string) *UpdateRepoTriggerRequest {
	s.TriggerTag = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerType(v string) *UpdateRepoTriggerRequest {
	s.TriggerType = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerUrl(v string) *UpdateRepoTriggerRequest {
	s.TriggerUrl = &v
	return s
}

func (s *UpdateRepoTriggerRequest) Validate() error {
	return dara.Validate(s)
}
