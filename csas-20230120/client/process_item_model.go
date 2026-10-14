// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessItem interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v string) *ProcessItem
	GetBundleId() *string
	SetDevType(v string) *ProcessItem
	GetDevType() *string
	SetDirectory(v string) *ProcessItem
	GetDirectory() *string
	SetProcess(v string) *ProcessItem
	GetProcess() *string
}

type ProcessItem struct {
	// The bundle identity (Bundle ID) of the application to which the process belongs. The example value com.apple.Terminal is the Bundle ID of the macOS terminal application.
	//
	// example:
	//
	// com.apple.Terminal
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The operating system type of the user terminal device to which the process belongs. Valid values:
	//
	// - windows: Windows operating system
	//
	// - macOS: macOS operating system
	//
	// - linux: Linux operating system
	//
	// example:
	//
	// macOS
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// The process folder path.
	//
	// example:
	//
	// /System/Applications/Utilities/Terminal.app/Contents/MacOS
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// Terminal
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s ProcessItem) String() string {
	return dara.Prettify(s)
}

func (s ProcessItem) GoString() string {
	return s.String()
}

func (s *ProcessItem) GetBundleId() *string {
	return s.BundleId
}

func (s *ProcessItem) GetDevType() *string {
	return s.DevType
}

func (s *ProcessItem) GetDirectory() *string {
	return s.Directory
}

func (s *ProcessItem) GetProcess() *string {
	return s.Process
}

func (s *ProcessItem) SetBundleId(v string) *ProcessItem {
	s.BundleId = &v
	return s
}

func (s *ProcessItem) SetDevType(v string) *ProcessItem {
	s.DevType = &v
	return s
}

func (s *ProcessItem) SetDirectory(v string) *ProcessItem {
	s.Directory = &v
	return s
}

func (s *ProcessItem) SetProcess(v string) *ProcessItem {
	s.Process = &v
	return s
}

func (s *ProcessItem) Validate() error {
	return dara.Validate(s)
}
