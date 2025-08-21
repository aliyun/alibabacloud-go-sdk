// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BatchUnbindDirectoriesRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *BatchUnbindDirectoriesRequest
	GetDirectoryId() *string
	SetOwnerId(v int64) *BatchUnbindDirectoriesRequest
	GetOwnerId() *int64
}

type BatchUnbindDirectoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 399*****488-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchUnbindDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchUnbindDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *BatchUnbindDirectoriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchUnbindDirectoriesRequest) SetDeviceId(v string) *BatchUnbindDirectoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindDirectoriesRequest) SetDirectoryId(v string) *BatchUnbindDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *BatchUnbindDirectoriesRequest) SetOwnerId(v int64) *BatchUnbindDirectoriesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
