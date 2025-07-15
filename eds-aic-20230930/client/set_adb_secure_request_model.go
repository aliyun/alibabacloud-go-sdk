// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAdbSecureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *SetAdbSecureRequest
	GetInstanceIds() []*string
	SetStatus(v int32) *SetAdbSecureRequest
	GetStatus() *int32
}

type SetAdbSecureRequest struct {
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The status of the ADB authentication feature.
	//
	// Valid values:
	//
	// 	- 0: The ADB authentication feature is disabled.
	//
	// 	- 1: The ADB authentication feature is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetAdbSecureRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAdbSecureRequest) GoString() string {
	return s.String()
}

func (s *SetAdbSecureRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *SetAdbSecureRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SetAdbSecureRequest) SetInstanceIds(v []*string) *SetAdbSecureRequest {
	s.InstanceIds = v
	return s
}

func (s *SetAdbSecureRequest) SetStatus(v int32) *SetAdbSecureRequest {
	s.Status = &v
	return s
}

func (s *SetAdbSecureRequest) Validate() error {
	return dara.Validate(s)
}
