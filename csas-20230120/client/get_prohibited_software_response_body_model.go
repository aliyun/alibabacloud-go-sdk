// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedSoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetProhibitedSoftwareResponseBody
	GetRequestId() *string
	SetSoftware(v *GetProhibitedSoftwareResponseBodySoftware) *GetProhibitedSoftwareResponseBody
	GetSoftware() *GetProhibitedSoftwareResponseBodySoftware
}

type GetProhibitedSoftwareResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 019F6DE3-3079-52DE-ABD1-39FB76B74FC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The prohibited software details.
	Software *GetProhibitedSoftwareResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Struct"`
}

func (s GetProhibitedSoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProhibitedSoftwareResponseBody) GetSoftware() *GetProhibitedSoftwareResponseBodySoftware {
	return s.Software
}

func (s *GetProhibitedSoftwareResponseBody) SetRequestId(v string) *GetProhibitedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBody) SetSoftware(v *GetProhibitedSoftwareResponseBodySoftware) *GetProhibitedSoftwareResponseBody {
	s.Software = v
	return s
}

func (s *GetProhibitedSoftwareResponseBody) Validate() error {
	if s.Software != nil {
		if err := s.Software.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProhibitedSoftwareResponseBodySoftware struct {
	// The creation time of the prohibited software, in the yyyy-MM-dd HH:mm:ss format. The time is displayed in UTC+8.
	//
	// example:
	//
	// 2025-09-05 10:20:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the prohibited software.
	//
	// example:
	//
	// P2P download tool
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the prohibited software is a system built-in prohibited software. Valid values:
	//
	// - **true**: A system built-in prohibited software that is shared across all Alibaba Cloud accounts and cannot be modified or deleted.
	//
	// - **false**: A custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The list of process configurations for the Linux operating system.
	LinuxProcesses []*GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses `json:"LinuxProcesses,omitempty" xml:"LinuxProcesses,omitempty" type:"Repeated"`
	// The list of process configurations for the macOS operating system.
	MacOSProcesses []*GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses `json:"MacOSProcesses,omitempty" xml:"MacOSProcesses,omitempty" type:"Repeated"`
	// The name of the prohibited software.
	//
	// example:
	//
	// Thunder
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the software prohibition policies that directly reference the prohibited software.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The prohibited software ID.
	//
	// example:
	//
	// swb-3e6a1f9c4b28****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The IDs of the prohibited software tags associated with the prohibited software.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The list of process configurations for the Windows operating system.
	WindowsProcesses []*GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses `json:"WindowsProcesses,omitempty" xml:"WindowsProcesses,omitempty" type:"Repeated"`
}

func (s GetProhibitedSoftwareResponseBodySoftware) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareResponseBodySoftware) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetDescription() *string {
	return s.Description
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetLinuxProcesses() []*GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	return s.LinuxProcesses
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetMacOSProcesses() []*GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	return s.MacOSProcesses
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetName() *string {
	return s.Name
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetTagIds() []*string {
	return s.TagIds
}

func (s *GetProhibitedSoftwareResponseBodySoftware) GetWindowsProcesses() []*GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	return s.WindowsProcesses
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetCreateTime(v string) *GetProhibitedSoftwareResponseBodySoftware {
	s.CreateTime = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetDescription(v string) *GetProhibitedSoftwareResponseBodySoftware {
	s.Description = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetIsDefault(v bool) *GetProhibitedSoftwareResponseBodySoftware {
	s.IsDefault = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetLinuxProcesses(v []*GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) *GetProhibitedSoftwareResponseBodySoftware {
	s.LinuxProcesses = v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetMacOSProcesses(v []*GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) *GetProhibitedSoftwareResponseBodySoftware {
	s.MacOSProcesses = v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetName(v string) *GetProhibitedSoftwareResponseBodySoftware {
	s.Name = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetPolicyIds(v []*string) *GetProhibitedSoftwareResponseBodySoftware {
	s.PolicyIds = v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetSoftwareId(v string) *GetProhibitedSoftwareResponseBodySoftware {
	s.SoftwareId = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetTagIds(v []*string) *GetProhibitedSoftwareResponseBodySoftware {
	s.TagIds = v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) SetWindowsProcesses(v []*GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) *GetProhibitedSoftwareResponseBodySoftware {
	s.WindowsProcesses = v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftware) Validate() error {
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

type GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses struct {
	// The application bundle identifier (Bundle ID). This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.xunlei.Thunder
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command-line arguments for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\Program Files\\Thunder Network\\Thunder
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// thuner.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetProcess() *string {
	return s.Process
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetBundleId(v string) *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.BundleId = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetCmdline(v string) *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Cmdline = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetDirectory(v string) *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Directory = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetProcess(v string) *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Process = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareLinuxProcesses) Validate() error {
	return dara.Validate(s)
}

type GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses struct {
	// The application bundle identifier (Bundle ID). This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.xunlei.Thunder
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command-line arguments for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\Program Files\\Thunder Network\\Thunder
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// thuner.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetProcess() *string {
	return s.Process
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetBundleId(v string) *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.BundleId = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetCmdline(v string) *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Cmdline = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetDirectory(v string) *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Directory = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetProcess(v string) *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Process = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareMacOSProcesses) Validate() error {
	return dara.Validate(s)
}

type GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses struct {
	// The application bundle identifier (Bundle ID). This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.xunlei.Thunder
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command-line arguments for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\Program Files\\Thunder Network\\Thunder
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// thuner.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetProcess() *string {
	return s.Process
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetBundleId(v string) *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.BundleId = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetCmdline(v string) *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Cmdline = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetDirectory(v string) *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Directory = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetProcess(v string) *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Process = &v
	return s
}

func (s *GetProhibitedSoftwareResponseBodySoftwareWindowsProcesses) Validate() error {
	return dara.Validate(s)
}
