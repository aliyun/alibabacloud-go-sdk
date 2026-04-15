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
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
