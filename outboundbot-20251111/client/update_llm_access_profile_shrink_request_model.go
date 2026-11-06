// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLlmAccessProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *UpdateLlmAccessProfileShrinkRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *UpdateLlmAccessProfileShrinkRequest
	GetInstanceId() *string
	SetProfileShrink(v string) *UpdateLlmAccessProfileShrinkRequest
	GetProfileShrink() *string
}

type UpdateLlmAccessProfileShrinkRequest struct {
	// The access profile ID.
	//
	// example:
	//
	// ed06946e-dcbb-4f3a-baf8-a5feae0198e9
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language model access configuration.
	ProfileShrink *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s UpdateLlmAccessProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLlmAccessProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLlmAccessProfileShrinkRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateLlmAccessProfileShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLlmAccessProfileShrinkRequest) GetProfileShrink() *string {
	return s.ProfileShrink
}

func (s *UpdateLlmAccessProfileShrinkRequest) SetAccessProfileId(v string) *UpdateLlmAccessProfileShrinkRequest {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateLlmAccessProfileShrinkRequest) SetInstanceId(v string) *UpdateLlmAccessProfileShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLlmAccessProfileShrinkRequest) SetProfileShrink(v string) *UpdateLlmAccessProfileShrinkRequest {
	s.ProfileShrink = &v
	return s
}

func (s *UpdateLlmAccessProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
