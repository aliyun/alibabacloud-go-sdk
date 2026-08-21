// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedSoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProhibitedSoftwareRequest
	GetDescription() *string
	SetLinuxProcesses(v []*UpdateProhibitedSoftwareRequestLinuxProcesses) *UpdateProhibitedSoftwareRequest
	GetLinuxProcesses() []*UpdateProhibitedSoftwareRequestLinuxProcesses
	SetMacOSProcesses(v []*UpdateProhibitedSoftwareRequestMacOSProcesses) *UpdateProhibitedSoftwareRequest
	GetMacOSProcesses() []*UpdateProhibitedSoftwareRequestMacOSProcesses
	SetName(v string) *UpdateProhibitedSoftwareRequest
	GetName() *string
	SetSoftwareId(v string) *UpdateProhibitedSoftwareRequest
	GetSoftwareId() *string
	SetTagIds(v []*string) *UpdateProhibitedSoftwareRequest
	GetTagIds() []*string
	SetWindowsProcesses(v []*UpdateProhibitedSoftwareRequestWindowsProcesses) *UpdateProhibitedSoftwareRequest
	GetWindowsProcesses() []*UpdateProhibitedSoftwareRequestWindowsProcesses
}

type UpdateProhibitedSoftwareRequest struct {
	// The description of the prohibited software.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of process configurations for the Linux operating system.
	LinuxProcesses []*UpdateProhibitedSoftwareRequestLinuxProcesses `json:"LinuxProcesses,omitempty" xml:"LinuxProcesses,omitempty" type:"Repeated"`
	// The list of process configurations for the macOS operating system.
	MacOSProcesses []*UpdateProhibitedSoftwareRequestMacOSProcesses `json:"MacOSProcesses,omitempty" xml:"MacOSProcesses,omitempty" type:"Repeated"`
	// The name of the prohibited software.
	//
	// example:
	//
	// Edge
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the custom prohibited software to update. Only custom prohibited software under the current Alibaba Cloud account can be updated. Built-in prohibited software cannot be updated. You can obtain the value from the following operations:
	//
	// - [ListProhibitedSoftware](~~ListProhibitedSoftware~~): queries prohibited software entries in batches.
	//
	// - [CreateProhibitedSoftware](~~CreateProhibitedSoftware~~): creates a custom prohibited software entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// swb-a43c9cbf88df****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The IDs of the custom prohibited software tags to associate. Duplicate values are not allowed.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The list of process configurations for the Windows operating system.
	WindowsProcesses []*UpdateProhibitedSoftwareRequestWindowsProcesses `json:"WindowsProcesses,omitempty" xml:"WindowsProcesses,omitempty" type:"Repeated"`
}

func (s UpdateProhibitedSoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProhibitedSoftwareRequest) GetLinuxProcesses() []*UpdateProhibitedSoftwareRequestLinuxProcesses {
	return s.LinuxProcesses
}

func (s *UpdateProhibitedSoftwareRequest) GetMacOSProcesses() []*UpdateProhibitedSoftwareRequestMacOSProcesses {
	return s.MacOSProcesses
}

func (s *UpdateProhibitedSoftwareRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProhibitedSoftwareRequest) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *UpdateProhibitedSoftwareRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdateProhibitedSoftwareRequest) GetWindowsProcesses() []*UpdateProhibitedSoftwareRequestWindowsProcesses {
	return s.WindowsProcesses
}

func (s *UpdateProhibitedSoftwareRequest) SetDescription(v string) *UpdateProhibitedSoftwareRequest {
	s.Description = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) SetLinuxProcesses(v []*UpdateProhibitedSoftwareRequestLinuxProcesses) *UpdateProhibitedSoftwareRequest {
	s.LinuxProcesses = v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) SetMacOSProcesses(v []*UpdateProhibitedSoftwareRequestMacOSProcesses) *UpdateProhibitedSoftwareRequest {
	s.MacOSProcesses = v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) SetName(v string) *UpdateProhibitedSoftwareRequest {
	s.Name = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) SetSoftwareId(v string) *UpdateProhibitedSoftwareRequest {
	s.SoftwareId = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) SetTagIds(v []*string) *UpdateProhibitedSoftwareRequest {
	s.TagIds = v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) SetWindowsProcesses(v []*UpdateProhibitedSoftwareRequestWindowsProcesses) *UpdateProhibitedSoftwareRequest {
	s.WindowsProcesses = v
	return s
}

func (s *UpdateProhibitedSoftwareRequest) Validate() error {
	if s.LinuxProcesses != nil {
		for _, item := range s.LinuxProcesses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MacOSProcesses != nil {
		for _, item := range s.MacOSProcesses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WindowsProcesses != nil {
		for _, item := range s.WindowsProcesses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProhibitedSoftwareRequestLinuxProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes. You must specify at least one of this parameter and Process. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// com.aliyun.security.sase
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process. If specified, only processes whose command line contains this content are matched. If left empty, the command line is not checked. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located. If specified, only processes with the same name in this directory are matched. If left empty, processes in any directory are matched. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// Everything.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s UpdateProhibitedSoftwareRequestLinuxProcesses) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareRequestLinuxProcesses) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) GetProcess() *string {
	return s.Process
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) SetBundleId(v string) *UpdateProhibitedSoftwareRequestLinuxProcesses {
	s.BundleId = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) SetCmdline(v string) *UpdateProhibitedSoftwareRequestLinuxProcesses {
	s.Cmdline = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) SetDirectory(v string) *UpdateProhibitedSoftwareRequestLinuxProcesses {
	s.Directory = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) SetProcess(v string) *UpdateProhibitedSoftwareRequestLinuxProcesses {
	s.Process = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestLinuxProcesses) Validate() error {
	return dara.Validate(s)
}

type UpdateProhibitedSoftwareRequestMacOSProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes. You must specify at least one of this parameter and Process. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// com.autotest.update
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process. If specified, only processes whose command line contains this content are matched. If left empty, the command line is not checked. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located. If specified, only processes with the same name in this directory are matched. If left empty, processes in any directory are matched. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// Everything.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s UpdateProhibitedSoftwareRequestMacOSProcesses) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareRequestMacOSProcesses) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) GetProcess() *string {
	return s.Process
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) SetBundleId(v string) *UpdateProhibitedSoftwareRequestMacOSProcesses {
	s.BundleId = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) SetCmdline(v string) *UpdateProhibitedSoftwareRequestMacOSProcesses {
	s.Cmdline = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) SetDirectory(v string) *UpdateProhibitedSoftwareRequestMacOSProcesses {
	s.Directory = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) SetProcess(v string) *UpdateProhibitedSoftwareRequestMacOSProcesses {
	s.Process = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestMacOSProcesses) Validate() error {
	return dara.Validate(s)
}

type UpdateProhibitedSoftwareRequestWindowsProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes. You must specify at least one of this parameter and Process. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// cn.apifox.app
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process. If specified, only processes whose command line contains this content are matched. If left empty, the command line is not checked. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located. If specified, only processes with the same name in this directory are matched. If left empty, processes in any directory are matched. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name. The value can be up to 1024 characters in length.
	//
	// example:
	//
	// Everything.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s UpdateProhibitedSoftwareRequestWindowsProcesses) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareRequestWindowsProcesses) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) GetProcess() *string {
	return s.Process
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) SetBundleId(v string) *UpdateProhibitedSoftwareRequestWindowsProcesses {
	s.BundleId = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) SetCmdline(v string) *UpdateProhibitedSoftwareRequestWindowsProcesses {
	s.Cmdline = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) SetDirectory(v string) *UpdateProhibitedSoftwareRequestWindowsProcesses {
	s.Directory = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) SetProcess(v string) *UpdateProhibitedSoftwareRequestWindowsProcesses {
	s.Process = &v
	return s
}

func (s *UpdateProhibitedSoftwareRequestWindowsProcesses) Validate() error {
	return dara.Validate(s)
}
