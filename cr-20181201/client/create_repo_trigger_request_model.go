// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateRepoTriggerRequest
	GetInstanceId() *string
	SetRepoId(v string) *CreateRepoTriggerRequest
	GetRepoId() *string
	SetTriggerName(v string) *CreateRepoTriggerRequest
	GetTriggerName() *string
	SetTriggerTag(v string) *CreateRepoTriggerRequest
	GetTriggerTag() *string
	SetTriggerType(v string) *CreateRepoTriggerRequest
	GetTriggerType() *string
	SetTriggerUrl(v string) *CreateRepoTriggerRequest
	GetTriggerUrl() *string
}

type CreateRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger1
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// >
	//
	// 	- If `TriggerType` is set to `ALL`, `TriggerTag` can be set to a string or an array, for example, `*`.
	//
	// 	- If `TriggerType` is set to `TAG_LIST`, `TriggerTag` must be set to an array, for example, `[1]`.
	//
	// 	- If `TriggerType` is set to `TAG_REG_EXP`, `TriggerTag` must be set to a string, for example, `*`.
	//
	// example:
	//
	// [1]
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- `ALL`: a trigger that supports both tags and regular expressions.
	//
	// 	- `TAG_LIST`: a tag-based trigger.
	//
	// 	- `TAG_REG_EXP`: a regular expression-based trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.mysite.com
	TriggerUrl *string `json:"TriggerUrl,omitempty" xml:"TriggerUrl,omitempty"`
}

func (s CreateRepoTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoTriggerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoTriggerRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepoTriggerRequest) GetTriggerName() *string {
	return s.TriggerName
}

func (s *CreateRepoTriggerRequest) GetTriggerTag() *string {
	return s.TriggerTag
}

func (s *CreateRepoTriggerRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateRepoTriggerRequest) GetTriggerUrl() *string {
	return s.TriggerUrl
}

func (s *CreateRepoTriggerRequest) SetInstanceId(v string) *CreateRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetRepoId(v string) *CreateRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerName(v string) *CreateRepoTriggerRequest {
	s.TriggerName = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerTag(v string) *CreateRepoTriggerRequest {
	s.TriggerTag = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerType(v string) *CreateRepoTriggerRequest {
	s.TriggerType = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerUrl(v string) *CreateRepoTriggerRequest {
	s.TriggerUrl = &v
	return s
}

func (s *CreateRepoTriggerRequest) Validate() error {
	return dara.Validate(s)
}
