// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ResizeDiskSizeRequest
	GetClusterId() *string
	SetNodeDiskSize(v int32) *ResizeDiskSizeRequest
	GetNodeDiskSize() *int32
}

type ResizeDiskSizeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 120
	NodeDiskSize *int32 `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
}

func (s ResizeDiskSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResizeDiskSizeRequest) GetNodeDiskSize() *int32 {
	return s.NodeDiskSize
}

func (s *ResizeDiskSizeRequest) SetClusterId(v string) *ResizeDiskSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeDiskSizeRequest) SetNodeDiskSize(v int32) *ResizeDiskSizeRequest {
	s.NodeDiskSize = &v
	return s
}

func (s *ResizeDiskSizeRequest) Validate() error {
	return dara.Validate(s)
}
