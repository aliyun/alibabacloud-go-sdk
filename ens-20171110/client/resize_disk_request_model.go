// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *ResizeDiskRequest
	GetDiskId() *string
	SetNewSize(v string) *ResizeDiskRequest
	GetNewSize() *string
}

type ResizeDiskRequest struct {
	// The ID of the disk that you want to resize.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5tzm9wnhzlhjzcbtxo465****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The size of the disk that you want to resize. Unit: GiB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	NewSize *string `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
}

func (s ResizeDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskRequest) GoString() string {
	return s.String()
}

func (s *ResizeDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResizeDiskRequest) GetNewSize() *string {
	return s.NewSize
}

func (s *ResizeDiskRequest) SetDiskId(v string) *ResizeDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResizeDiskRequest) SetNewSize(v string) *ResizeDiskRequest {
	s.NewSize = &v
	return s
}

func (s *ResizeDiskRequest) Validate() error {
	return dara.Validate(s)
}
