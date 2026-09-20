// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeMultiZoneClusterDiskSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ResizeMultiZoneClusterDiskSizeRequest
	GetClusterId() *string
	SetCoreDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest
	GetCoreDiskSize() *int32
	SetLogDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest
	GetLogDiskSize() *int32
}

type ResizeMultiZoneClusterDiskSizeRequest struct {
	// The ID of the multi-zone instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-f5d6vc2r8d6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The disk size of core nodes. The value must be greater than the current disk size and must be a multiple of 40. Unit: GB. Maximum value: 64000.
	//
	// example:
	//
	// 480
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// The disk size of log nodes. The value must be greater than the current disk size of log nodes and must be a multiple of 40. Unit: GB. Maximum value: 8000.
	//
	// example:
	//
	// 440
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) GetCoreDiskSize() *int32 {
	return s.CoreDiskSize
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) GetLogDiskSize() *int32 {
	return s.LogDiskSize
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetClusterId(v string) *ResizeMultiZoneClusterDiskSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetCoreDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetLogDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest {
	s.LogDiskSize = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) Validate() error {
	return dara.Validate(s)
}
