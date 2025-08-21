// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BatchBindDirectoriesRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *BatchBindDirectoriesRequest
	GetDirectoryId() *string
	SetOwnerId(v int64) *BatchBindDirectoriesRequest
	GetOwnerId() *int64
}

type BatchBindDirectoriesRequest struct {
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

func (s BatchBindDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchBindDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchBindDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *BatchBindDirectoriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchBindDirectoriesRequest) SetDeviceId(v string) *BatchBindDirectoriesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchBindDirectoriesRequest) SetDirectoryId(v string) *BatchBindDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *BatchBindDirectoriesRequest) SetOwnerId(v int64) *BatchBindDirectoriesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
