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
	// The access configuration ID.
	//
	// example:
	//
	// 73ce1225-c3cc-43e5-b169-1ef56afaeb9f
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
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
