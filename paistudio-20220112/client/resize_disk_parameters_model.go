// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskParameters interface {
	dara.Model
	String() string
	GoString() string
	SetNewDiskSize(v string) *ResizeDiskParameters
	GetNewDiskSize() *string
}

type ResizeDiskParameters struct {
	// The target disk capacity after the change.
	NewDiskSize *string `json:"NewDiskSize,omitempty" xml:"NewDiskSize,omitempty"`
}

func (s ResizeDiskParameters) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskParameters) GoString() string {
	return s.String()
}

func (s *ResizeDiskParameters) GetNewDiskSize() *string {
	return s.NewDiskSize
}

func (s *ResizeDiskParameters) SetNewDiskSize(v string) *ResizeDiskParameters {
	s.NewDiskSize = &v
	return s
}

func (s *ResizeDiskParameters) Validate() error {
	return dara.Validate(s)
}
