// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDevicesRequest
	GetInstanceId() *string
	SetUserId(v string) *ListDevicesRequest
	GetUserId() *string
}

type ListDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDevicesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListDevicesRequest) SetInstanceId(v string) *ListDevicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDevicesRequest) SetUserId(v string) *ListDevicesRequest {
	s.UserId = &v
	return s
}

func (s *ListDevicesRequest) Validate() error {
	return dara.Validate(s)
}
