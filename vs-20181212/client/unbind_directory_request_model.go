// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *UnbindDirectoryRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *UnbindDirectoryRequest
	GetDirectoryId() *string
	SetOwnerId(v int64) *UnbindDirectoryRequest
	GetOwnerId() *int64
}

type UnbindDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 34871************3380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3998**************9488-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UnbindDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UnbindDirectoryRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnbindDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UnbindDirectoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindDirectoryRequest) SetDeviceId(v string) *UnbindDirectoryRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindDirectoryRequest) SetDirectoryId(v string) *UnbindDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UnbindDirectoryRequest) SetOwnerId(v int64) *UnbindDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
