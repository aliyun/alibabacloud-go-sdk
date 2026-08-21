// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedSoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListProhibitedSoftwareResponseBody
	GetRequestId() *string
	SetSoftware(v []*ListProhibitedSoftwareResponseBodySoftware) *ListProhibitedSoftwareResponseBody
	GetSoftware() []*ListProhibitedSoftwareResponseBodySoftware
	SetTotalNum(v int64) *ListProhibitedSoftwareResponseBody
	GetTotalNum() *int64
}

type ListProhibitedSoftwareResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C7F49DCC-8EFE-59BE-8947-0529CC458C59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of prohibited software.
	Software []*ListProhibitedSoftwareResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Repeated"`
	// The total number of prohibited software entries.
	//
	// example:
	//
	// 28
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListProhibitedSoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProhibitedSoftwareResponseBody) GetSoftware() []*ListProhibitedSoftwareResponseBodySoftware {
	return s.Software
}

func (s *ListProhibitedSoftwareResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListProhibitedSoftwareResponseBody) SetRequestId(v string) *ListProhibitedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBody) SetSoftware(v []*ListProhibitedSoftwareResponseBodySoftware) *ListProhibitedSoftwareResponseBody {
	s.Software = v
	return s
}

func (s *ListProhibitedSoftwareResponseBody) SetTotalNum(v int64) *ListProhibitedSoftwareResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBody) Validate() error {
	if s.Software != nil {
		for _, item := range s.Software {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProhibitedSoftwareResponseBodySoftware struct {
	// The time when the prohibited software was created, in the yyyy-MM-dd HH:mm:ss format. The time is in the UTC+8 time zone.
	//
	// example:
	//
	// 2026-08-19 10:24:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the prohibited software.
	//
	// example:
	//
	// P2P download software
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The collection of dynamic policy IDs that reference the prohibited software as a disposal action.
	DynamicPolicyIds []*string `json:"DynamicPolicyIds,omitempty" xml:"DynamicPolicyIds,omitempty" type:"Repeated"`
	// Indicates whether the software is a system built-in prohibited software. Valid values:
	//
	// - **true**: A system built-in prohibited software that is shared across all Alibaba Cloud accounts and cannot be modified or deleted.
	//
	// - **false**: Custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The list of process configurations for the Linux operating system.
	LinuxProcesses []*ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses `json:"LinuxProcesses,omitempty" xml:"LinuxProcesses,omitempty" type:"Repeated"`
	// The list of process configurations for the macOS operating system.
	MacOSProcesses []*ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses `json:"MacOSProcesses,omitempty" xml:"MacOSProcesses,omitempty" type:"Repeated"`
	// The name of the prohibited software.
	//
	// example:
	//
	// Thunder
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of software prohibition policy IDs that directly reference the prohibited software.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The ID of the prohibited software.
	//
	// example:
	//
	// swb-238eee6903e8****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The collection of prohibited software tag IDs associated with the prohibited software.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The list of process configurations for the Windows operating system.
	WindowsProcesses []*ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses `json:"WindowsProcesses,omitempty" xml:"WindowsProcesses,omitempty" type:"Repeated"`
}

func (s ListProhibitedSoftwareResponseBodySoftware) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareResponseBodySoftware) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetDescription() *string {
	return s.Description
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetDynamicPolicyIds() []*string {
	return s.DynamicPolicyIds
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetLinuxProcesses() []*ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	return s.LinuxProcesses
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetMacOSProcesses() []*ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	return s.MacOSProcesses
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetName() *string {
	return s.Name
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListProhibitedSoftwareResponseBodySoftware) GetWindowsProcesses() []*ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	return s.WindowsProcesses
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetCreateTime(v string) *ListProhibitedSoftwareResponseBodySoftware {
	s.CreateTime = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetDescription(v string) *ListProhibitedSoftwareResponseBodySoftware {
	s.Description = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetDynamicPolicyIds(v []*string) *ListProhibitedSoftwareResponseBodySoftware {
	s.DynamicPolicyIds = v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetIsDefault(v bool) *ListProhibitedSoftwareResponseBodySoftware {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetLinuxProcesses(v []*ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) *ListProhibitedSoftwareResponseBodySoftware {
	s.LinuxProcesses = v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetMacOSProcesses(v []*ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) *ListProhibitedSoftwareResponseBodySoftware {
	s.MacOSProcesses = v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetName(v string) *ListProhibitedSoftwareResponseBodySoftware {
	s.Name = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetPolicyIds(v []*string) *ListProhibitedSoftwareResponseBodySoftware {
	s.PolicyIds = v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetSoftwareId(v string) *ListProhibitedSoftwareResponseBodySoftware {
	s.SoftwareId = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetTagIds(v []*string) *ListProhibitedSoftwareResponseBodySoftware {
	s.TagIds = v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) SetWindowsProcesses(v []*ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) *ListProhibitedSoftwareResponseBodySoftware {
	s.WindowsProcesses = v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftware) Validate() error {
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

type ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses struct {
	// The application bundle identifier (Bundle ID). This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.autotest.app
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command-line parameters for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// /User/sase/Applications
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// thunder
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetProcess() *string {
	return s.Process
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetBundleId(v string) *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.BundleId = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetCmdline(v string) *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Cmdline = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetDirectory(v string) *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Directory = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetProcess(v string) *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Process = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareLinuxProcesses) Validate() error {
	return dara.Validate(s)
}

type ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses struct {
	// The application bundle identifier (Bundle ID). This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.xunlei.Thunder
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command-line parameters for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// ~/Applications
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// autotest.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetProcess() *string {
	return s.Process
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetBundleId(v string) *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.BundleId = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetCmdline(v string) *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Cmdline = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetDirectory(v string) *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Directory = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetProcess(v string) *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Process = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareMacOSProcesses) Validate() error {
	return dara.Validate(s)
}

type ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses struct {
	// The application bundle identifier (Bundle ID). This parameter is required only for macOS processes.
	//
	// example:
	//
	// com.xunlei.Thunder
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The command-line parameters for starting the process.
	//
	// example:
	//
	// --start-minimized
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The directory where the process is located.
	//
	// example:
	//
	// C:\\Program Files\\Thunder Network
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// SASE.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetProcess() *string {
	return s.Process
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetBundleId(v string) *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.BundleId = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetCmdline(v string) *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Cmdline = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetDirectory(v string) *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Directory = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetProcess(v string) *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Process = &v
	return s
}

func (s *ListProhibitedSoftwareResponseBodySoftwareWindowsProcesses) Validate() error {
	return dara.Validate(s)
}
