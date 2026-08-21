// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedSoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProhibitedSoftwareRequest
	GetDescription() *string
	SetLinuxProcesses(v []*CreateProhibitedSoftwareRequestLinuxProcesses) *CreateProhibitedSoftwareRequest
	GetLinuxProcesses() []*CreateProhibitedSoftwareRequestLinuxProcesses
	SetMacOSProcesses(v []*CreateProhibitedSoftwareRequestMacOSProcesses) *CreateProhibitedSoftwareRequest
	GetMacOSProcesses() []*CreateProhibitedSoftwareRequestMacOSProcesses
	SetName(v string) *CreateProhibitedSoftwareRequest
	GetName() *string
	SetTagIds(v []*string) *CreateProhibitedSoftwareRequest
	GetTagIds() []*string
	SetWindowsProcesses(v []*CreateProhibitedSoftwareRequestWindowsProcesses) *CreateProhibitedSoftwareRequest
	GetWindowsProcesses() []*CreateProhibitedSoftwareRequestWindowsProcesses
}

type CreateProhibitedSoftwareRequest struct {
	// The description of the disabled software.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of process configurations for the Linux operating system.
	LinuxProcesses []*CreateProhibitedSoftwareRequestLinuxProcesses `json:"LinuxProcesses,omitempty" xml:"LinuxProcesses,omitempty" type:"Repeated"`
	// The list of process configurations for the macOS operating system.
	MacOSProcesses []*CreateProhibitedSoftwareRequestMacOSProcesses `json:"MacOSProcesses,omitempty" xml:"MacOSProcesses,omitempty" type:"Repeated"`
	// The name of the disabled software.
	//
	// This parameter is required.
	//
	// example:
	//
	// Thunder
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of custom disabled software tags to associate. Duplicate values are not allowed.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The list of process configurations for the Windows operating system.
	WindowsProcesses []*CreateProhibitedSoftwareRequestWindowsProcesses `json:"WindowsProcesses,omitempty" xml:"WindowsProcesses,omitempty" type:"Repeated"`
}

func (s CreateProhibitedSoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProhibitedSoftwareRequest) GetLinuxProcesses() []*CreateProhibitedSoftwareRequestLinuxProcesses {
	return s.LinuxProcesses
}

func (s *CreateProhibitedSoftwareRequest) GetMacOSProcesses() []*CreateProhibitedSoftwareRequestMacOSProcesses {
	return s.MacOSProcesses
}

func (s *CreateProhibitedSoftwareRequest) GetName() *string {
	return s.Name
}

func (s *CreateProhibitedSoftwareRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreateProhibitedSoftwareRequest) GetWindowsProcesses() []*CreateProhibitedSoftwareRequestWindowsProcesses {
	return s.WindowsProcesses
}

func (s *CreateProhibitedSoftwareRequest) SetDescription(v string) *CreateProhibitedSoftwareRequest {
	s.Description = &v
	return s
}

func (s *CreateProhibitedSoftwareRequest) SetLinuxProcesses(v []*CreateProhibitedSoftwareRequestLinuxProcesses) *CreateProhibitedSoftwareRequest {
	s.LinuxProcesses = v
	return s
}

func (s *CreateProhibitedSoftwareRequest) SetMacOSProcesses(v []*CreateProhibitedSoftwareRequestMacOSProcesses) *CreateProhibitedSoftwareRequest {
	s.MacOSProcesses = v
	return s
}

func (s *CreateProhibitedSoftwareRequest) SetName(v string) *CreateProhibitedSoftwareRequest {
	s.Name = &v
	return s
}

func (s *CreateProhibitedSoftwareRequest) SetTagIds(v []*string) *CreateProhibitedSoftwareRequest {
	s.TagIds = v
	return s
}

func (s *CreateProhibitedSoftwareRequest) SetWindowsProcesses(v []*CreateProhibitedSoftwareRequestWindowsProcesses) *CreateProhibitedSoftwareRequest {
	s.WindowsProcesses = v
	return s
}

func (s *CreateProhibitedSoftwareRequest) Validate() error {
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

type CreateProhibitedSoftwareRequestLinuxProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes. You must specify at least one of BundleId and Process. Maximum length: 1024 characters.
	//
	// example:
	//
	// com.autotest.app
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process. If specified, only processes whose command line contains this content are matched. If left empty, the command line is not checked. Maximum length: 1024 characters.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located. If specified, only processes with the same name in this directory are matched. If left empty, processes in any directory are matched. Maximum length: 1024 characters.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name. Maximum length: 1024 characters.
	//
	// example:
	//
	// autotest.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s CreateProhibitedSoftwareRequestLinuxProcesses) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareRequestLinuxProcesses) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) GetProcess() *string {
	return s.Process
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) SetBundleId(v string) *CreateProhibitedSoftwareRequestLinuxProcesses {
	s.BundleId = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) SetCmdline(v string) *CreateProhibitedSoftwareRequestLinuxProcesses {
	s.Cmdline = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) SetDirectory(v string) *CreateProhibitedSoftwareRequestLinuxProcesses {
	s.Directory = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) SetProcess(v string) *CreateProhibitedSoftwareRequestLinuxProcesses {
	s.Process = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestLinuxProcesses) Validate() error {
	return dara.Validate(s)
}

type CreateProhibitedSoftwareRequestMacOSProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes. You must specify at least one of BundleId and Process. Maximum length: 1024 characters.
	//
	// example:
	//
	// com.autotest.app
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process. If specified, only processes whose command line contains this content are matched. If left empty, the command line is not checked. Maximum length: 1024 characters.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located. If specified, only processes with the same name in this directory are matched. If left empty, processes in any directory are matched. Maximum length: 1024 characters.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name. Maximum length: 1024 characters.
	//
	// example:
	//
	// WeChat
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s CreateProhibitedSoftwareRequestMacOSProcesses) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareRequestMacOSProcesses) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) GetProcess() *string {
	return s.Process
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) SetBundleId(v string) *CreateProhibitedSoftwareRequestMacOSProcesses {
	s.BundleId = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) SetCmdline(v string) *CreateProhibitedSoftwareRequestMacOSProcesses {
	s.Cmdline = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) SetDirectory(v string) *CreateProhibitedSoftwareRequestMacOSProcesses {
	s.Directory = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) SetProcess(v string) *CreateProhibitedSoftwareRequestMacOSProcesses {
	s.Process = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestMacOSProcesses) Validate() error {
	return dara.Validate(s)
}

type CreateProhibitedSoftwareRequestWindowsProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes. You must specify at least one of BundleId and Process. Maximum length: 1024 characters.
	//
	// example:
	//
	// com.autotest.update
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process. If specified, only processes whose command line contains this content are matched. If left empty, the command line is not checked. Maximum length: 1024 characters.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located. If specified, only processes with the same name in this directory are matched. If left empty, processes in any directory are matched. Maximum length: 1024 characters.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name. Maximum length: 1024 characters.
	//
	// example:
	//
	// autotest.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s CreateProhibitedSoftwareRequestWindowsProcesses) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareRequestWindowsProcesses) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) GetProcess() *string {
	return s.Process
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) SetBundleId(v string) *CreateProhibitedSoftwareRequestWindowsProcesses {
	s.BundleId = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) SetCmdline(v string) *CreateProhibitedSoftwareRequestWindowsProcesses {
	s.Cmdline = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) SetDirectory(v string) *CreateProhibitedSoftwareRequestWindowsProcesses {
	s.Directory = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) SetProcess(v string) *CreateProhibitedSoftwareRequestWindowsProcesses {
	s.Process = &v
	return s
}

func (s *CreateProhibitedSoftwareRequestWindowsProcesses) Validate() error {
	return dara.Validate(s)
}
