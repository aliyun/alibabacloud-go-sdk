// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedSoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateProhibitedSoftwareResponseBody
	GetRequestId() *string
	SetSoftware(v *CreateProhibitedSoftwareResponseBodySoftware) *CreateProhibitedSoftwareResponseBody
	GetSoftware() *CreateProhibitedSoftwareResponseBodySoftware
}

type CreateProhibitedSoftwareResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B9AC7B08-80F5-5EDD-8E6B-033F2FE5D4E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the disabled software.
	Software *CreateProhibitedSoftwareResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Struct"`
}

func (s CreateProhibitedSoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProhibitedSoftwareResponseBody) GetSoftware() *CreateProhibitedSoftwareResponseBodySoftware {
	return s.Software
}

func (s *CreateProhibitedSoftwareResponseBody) SetRequestId(v string) *CreateProhibitedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBody) SetSoftware(v *CreateProhibitedSoftwareResponseBodySoftware) *CreateProhibitedSoftwareResponseBody {
	s.Software = v
	return s
}

func (s *CreateProhibitedSoftwareResponseBody) Validate() error {
	if s.Software != nil {
		if err := s.Software.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProhibitedSoftwareResponseBodySoftware struct {
	// The creation time of the disabled software, in the yyyy-MM-dd HH:mm:ss format. The time is in the UTC+8 time zone.
	//
	// example:
	//
	// 2025-09-05 10:20:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the disabled software.
	//
	// example:
	//
	// Endpoint group targeting github
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of process configurations for the Linux operating system.
	LinuxProcesses []*CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses `json:"LinuxProcesses,omitempty" xml:"LinuxProcesses,omitempty" type:"Repeated"`
	// The list of process configurations for the macOS operating system.
	MacOSProcesses []*CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses `json:"MacOSProcesses,omitempty" xml:"MacOSProcesses,omitempty" type:"Repeated"`
	// The software name.
	//
	// example:
	//
	// shell
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the disabled software.
	//
	// example:
	//
	// swb-83995ff2ae38****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The IDs of disabled software tags associated with this disabled software.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The list of process configurations for the Windows operating system.
	WindowsProcesses []*CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses `json:"WindowsProcesses,omitempty" xml:"WindowsProcesses,omitempty" type:"Repeated"`
}

func (s CreateProhibitedSoftwareResponseBodySoftware) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareResponseBodySoftware) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetDescription() *string {
	return s.Description
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetLinuxProcesses() []*CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	return s.LinuxProcesses
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetMacOSProcesses() []*CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	return s.MacOSProcesses
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetName() *string {
	return s.Name
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) GetWindowsProcesses() []*CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	return s.WindowsProcesses
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetCreateTime(v string) *CreateProhibitedSoftwareResponseBodySoftware {
	s.CreateTime = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetDescription(v string) *CreateProhibitedSoftwareResponseBodySoftware {
	s.Description = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetLinuxProcesses(v []*CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) *CreateProhibitedSoftwareResponseBodySoftware {
	s.LinuxProcesses = v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetMacOSProcesses(v []*CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) *CreateProhibitedSoftwareResponseBodySoftware {
	s.MacOSProcesses = v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetName(v string) *CreateProhibitedSoftwareResponseBodySoftware {
	s.Name = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetSoftwareId(v string) *CreateProhibitedSoftwareResponseBodySoftware {
	s.SoftwareId = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetTagIds(v []*string) *CreateProhibitedSoftwareResponseBodySoftware {
	s.TagIds = v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) SetWindowsProcesses(v []*CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) *CreateProhibitedSoftwareResponseBodySoftware {
	s.WindowsProcesses = v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftware) Validate() error {
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

type CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses struct {
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
	// C:\\Program Files\\Thunder Network\\Thunder
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// terraform
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) GetProcess() *string {
	return s.Process
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetBundleId(v string) *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.BundleId = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetCmdline(v string) *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Cmdline = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetDirectory(v string) *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Directory = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) SetProcess(v string) *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses {
	s.Process = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareLinuxProcesses) Validate() error {
	return dara.Validate(s)
}

type CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses struct {
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
	// C:\\Program Files\\Thunder Network\\Thunder
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// autotest_update.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) GetProcess() *string {
	return s.Process
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetBundleId(v string) *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.BundleId = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetCmdline(v string) *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Cmdline = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetDirectory(v string) *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Directory = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) SetProcess(v string) *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses {
	s.Process = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareMacOSProcesses) Validate() error {
	return dara.Validate(s)
}

type CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses struct {
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
	// C:\\Program Files\\Thunder Network\\Thunder
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The process name.
	//
	// example:
	//
	// anaconda3.exe
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
}

func (s CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetDirectory() *string {
	return s.Directory
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) GetProcess() *string {
	return s.Process
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetBundleId(v string) *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.BundleId = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetCmdline(v string) *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Cmdline = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetDirectory(v string) *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Directory = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) SetProcess(v string) *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses {
	s.Process = &v
	return s
}

func (s *CreateProhibitedSoftwareResponseBodySoftwareWindowsProcesses) Validate() error {
	return dara.Validate(s)
}
