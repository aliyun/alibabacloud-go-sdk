// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLlmAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateLlmAccessProfileShrinkRequest
	GetInstanceId() *string
	SetProfileShrink(v string) *CreateLlmAccessProfileShrinkRequest
	GetProfileShrink() *string
}

type CreateLlmAccessProfileShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 33014787-cc13-49d3-ab2f-a98aa8f15fbb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language model access configuration.
	ProfileShrink *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s CreateLlmAccessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLlmAccessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLlmAccessProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLlmAccessProfileShrinkRequest) GetProfileShrink() *string {
	return s.ProfileShrink
}

func (s *CreateLlmAccessProfileShrinkRequest) SetInstanceId(v string) *CreateLlmAccessProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLlmAccessProfileShrinkRequest) SetProfileShrink(v string) *CreateLlmAccessProfileShrinkRequest {
	s.ProfileShrink = &v
	return s
}

func (s *CreateLlmAccessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
