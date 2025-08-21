// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BindDirectoryRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *BindDirectoryRequest
	GetDirectoryId() *string
	SetOwnerId(v int64) *BindDirectoryRequest
	GetOwnerId() *int64
}

type BindDirectoryRequest struct {
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

func (s BindDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDirectoryRequest) GoString() string {
	return s.String()
}

func (s *BindDirectoryRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *BindDirectoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDirectoryRequest) SetDeviceId(v string) *BindDirectoryRequest {
	s.DeviceId = &v
	return s
}

func (s *BindDirectoryRequest) SetDirectoryId(v string) *BindDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *BindDirectoryRequest) SetOwnerId(v int64) *BindDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
