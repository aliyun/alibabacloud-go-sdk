// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *DeleteLlmAccessProfileRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *DeleteLlmAccessProfileRequest
	GetInstanceId() *string
}

type DeleteLlmAccessProfileRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteLlmAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteLlmAccessProfileRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *DeleteLlmAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteLlmAccessProfileRequest) SetAccessProfileId(v string) *DeleteLlmAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *DeleteLlmAccessProfileRequest) SetInstanceId(v string) *DeleteLlmAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLlmAccessProfileRequest) Validate() error {
	return dara.Validate(s)
}
