// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizTags(v []*string) *UpdatePromptRequest
	GetBizTags() []*string
	SetDescription(v string) *UpdatePromptRequest
	GetDescription() *string
	SetLabels(v map[string]interface{}) *UpdatePromptRequest
	GetLabels() map[string]interface{}
	SetNamespaceId(v string) *UpdatePromptRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *UpdatePromptRequest
	GetPromptKey() *string
}

type UpdatePromptRequest struct {
	// The list of business tags. The value is a string array.
	//
	// example:
	//
	// ["cs","qa","support"]
	BizTags []*string `json:"BizTags,omitempty" xml:"BizTags,omitempty" type:"Repeated"`
	// The description of the prompt.
	//
	// example:
	//
	// 客服问答 Prompt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mapping between prompt versions and labels.
	//
	// example:
	//
	// {"latest":"0.0.1","stable":"0.0.1"}
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The unique identifier of the prompt.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer-service-qa
	PromptKey *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
}

func (s UpdatePromptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptRequest) GoString() string {
	return s.String()
}

func (s *UpdatePromptRequest) GetBizTags() []*string {
	return s.BizTags
}

func (s *UpdatePromptRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePromptRequest) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *UpdatePromptRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdatePromptRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *UpdatePromptRequest) SetBizTags(v []*string) *UpdatePromptRequest {
	s.BizTags = v
	return s
}

func (s *UpdatePromptRequest) SetDescription(v string) *UpdatePromptRequest {
	s.Description = &v
	return s
}

func (s *UpdatePromptRequest) SetLabels(v map[string]interface{}) *UpdatePromptRequest {
	s.Labels = v
	return s
}

func (s *UpdatePromptRequest) SetNamespaceId(v string) *UpdatePromptRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdatePromptRequest) SetPromptKey(v string) *UpdatePromptRequest {
	s.PromptKey = &v
	return s
}

func (s *UpdatePromptRequest) Validate() error {
	return dara.Validate(s)
}
