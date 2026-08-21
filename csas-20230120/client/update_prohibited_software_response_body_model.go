// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedSoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProhibitedSoftwareResponseBody
	GetRequestId() *string
	SetSoftware(v *UpdateProhibitedSoftwareResponseBodySoftware) *UpdateProhibitedSoftwareResponseBody
	GetSoftware() *UpdateProhibitedSoftwareResponseBodySoftware
}

type UpdateProhibitedSoftwareResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2123E64A-FB25-561F-9988-B8781E430694
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the prohibited software.
	Software *UpdateProhibitedSoftwareResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Struct"`
}

func (s UpdateProhibitedSoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProhibitedSoftwareResponseBody) GetSoftware() *UpdateProhibitedSoftwareResponseBodySoftware {
	return s.Software
}

func (s *UpdateProhibitedSoftwareResponseBody) SetRequestId(v string) *UpdateProhibitedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBody) SetSoftware(v *UpdateProhibitedSoftwareResponseBodySoftware) *UpdateProhibitedSoftwareResponseBody {
	s.Software = v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBody) Validate() error {
	if s.Software != nil {
		if err := s.Software.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProhibitedSoftwareResponseBodySoftware struct {
	// The time when the prohibited software was created, in the yyyy-MM-dd HH:mm:ss format. The time is displayed in UTC+8.
	//
	// example:
	//
	// 2023-08-17 09:49:03
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the prohibited software.
	//
	// example:
	//
	// kxi3
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of process configurations for the Linux operating system.
	LinuxProcesses []*UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses `json:"LinuxProcesses,omitempty" xml:"LinuxProcesses,omitempty" type:"Repeated"`
	// The list of process configurations for the macOS operating system.
	MacOSProcesses []*UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses `json:"MacOSProcesses,omitempty" xml:"MacOSProcesses,omitempty" type:"Repeated"`
	// The name of the prohibited software.
	//
	// example:
	//
	// xshell
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the software prohibition policies that directly reference this prohibited software.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The ID of the prohibited software.
	//
	// example:
	//
	// swb-c64076fa7afd****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The IDs of the prohibited software tags associated with this prohibited software.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The list of process configurations for the Windows operating system.
	WindowsProcesses []*UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses `json:"WindowsProcesses,omitempty" xml:"WindowsProcesses,omitempty" type:"Repeated"`
}

func (s UpdateProhibitedSoftwareResponseBodySoftware) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareResponseBodySoftware) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetDescription() *string {
	return s.Description
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetLinuxProcesses() []*UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	return s.LinuxProcesses
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetMacOSProcesses() []*UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	return s.MacOSProcesses
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetName() *string {
	return s.Name
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) GetWindowsProcesses() []*UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	return s.WindowsProcesses
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetCreateTime(v string) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.CreateTime = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetDescription(v string) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.Description = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetLinuxProcesses(v []*UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.LinuxProcesses = v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetMacOSProcesses(v []*UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.MacOSProcesses = v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetName(v string) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.Name = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetPolicyIds(v []*string) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.PolicyIds = v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetSoftwareId(v string) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.SoftwareId = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetTagIds(v []*string) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.TagIds = v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) SetWindowsProcesses(v []*UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) *UpdateProhibitedSoftwareResponseBodySoftware {
	s.WindowsProcesses = v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftware) Validate() error {
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

type UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.aliyun.security.sase
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// kismain.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetProcess() *string {
	return s.Process
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetBundleId(v string) *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.BundleId = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetCmdline(v string) *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Cmdline = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetDirectory(v string) *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Directory = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetProcess(v string) *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Process = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) Validate() error {
	return dara.Validate(s)
}

type UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.autotest.app
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// QQPCTray.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetProcess() *string {
	return s.Process
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetBundleId(v string) *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.BundleId = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetCmdline(v string) *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Cmdline = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetDirectory(v string) *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Directory = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetProcess(v string) *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Process = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) Validate() error {
	return dara.Validate(s)
}

type UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses struct {
	// The bundle ID of the application. This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.aliyun.security.sase
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command line parameters for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\\\autotest
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// QQPCTray.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetProcess() *string {
	return s.Process
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetBundleId(v string) *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.BundleId = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetCmdline(v string) *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Cmdline = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetDirectory(v string) *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Directory = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetProcess(v string) *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Process = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) Validate() error {
	return dara.Validate(s)
}
