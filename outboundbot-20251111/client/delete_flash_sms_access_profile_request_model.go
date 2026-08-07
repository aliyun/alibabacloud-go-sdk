// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlashSmsAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *DeleteFlashSmsAccessProfileRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *DeleteFlashSmsAccessProfileRequest
	GetInstanceId() *string
}

type DeleteFlashSmsAccessProfileRequest struct {
	// 接入配置ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteFlashSmsAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlashSmsAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlashSmsAccessProfileRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *DeleteFlashSmsAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteFlashSmsAccessProfileRequest) SetAccessProfileId(v string) *DeleteFlashSmsAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileRequest) SetInstanceId(v string) *DeleteFlashSmsAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileRequest) Validate() error {
	return dara.Validate(s)
}
