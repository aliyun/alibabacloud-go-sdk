// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*DescribeGuestApplicationsResponseBodyApplications) *DescribeGuestApplicationsResponseBody
	GetApplications() []*DescribeGuestApplicationsResponseBodyApplications
	SetRequestId(v string) *DescribeGuestApplicationsResponseBody
	GetRequestId() *string
}

type DescribeGuestApplicationsResponseBody struct {
	// The applications.
	Applications []*DescribeGuestApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 272CF39E-B5DE-5BE3-A09B-B43F1026****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestApplicationsResponseBody) GetApplications() []*DescribeGuestApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *DescribeGuestApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGuestApplicationsResponseBody) SetApplications(v []*DescribeGuestApplicationsResponseBodyApplications) *DescribeGuestApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *DescribeGuestApplicationsResponseBody) SetRequestId(v string) *DescribeGuestApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGuestApplicationsResponseBodyApplications struct {
	// The application name.
	//
	// example:
	//
	// Google Chrome
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application version.
	//
	// example:
	//
	// 115.0.5790.110
	ApplicationVersion *string `json:"ApplicationVersion,omitempty" xml:"ApplicationVersion,omitempty"`
	// The CPU utilization (%).
	//
	// example:
	//
	// 89
	CpuPercent *float64 `json:"CpuPercent,omitempty" xml:"CpuPercent,omitempty"`
	// The GPU utilization (%).
	//
	// example:
	//
	// 15
	GpuPercent *float64 `json:"GpuPercent,omitempty" xml:"GpuPercent,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://app-center-icon-prod-shanghai.oss-cn-shanghai.aliyuncs.com/market/preload/default****.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The I/O read and write performance. Unit: byte/s.
	//
	// example:
	//
	// 124906.0
	IoSpeed *float64 `json:"IoSpeed,omitempty" xml:"IoSpeed,omitempty"`
	// The memory utilization (%).
	//
	// example:
	//
	// 34
	MemPercent *float64 `json:"MemPercent,omitempty" xml:"MemPercent,omitempty"`
	// The process ID (PID).
	//
	// example:
	//
	// 1357
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The processes.
	ProcessData []*DescribeGuestApplicationsResponseBodyApplicationsProcessData `json:"ProcessData,omitempty" xml:"ProcessData,omitempty" type:"Repeated"`
	// The path to the process.
	//
	// example:
	//
	// C:\\\\Program Files\\\\Google\\\\Chrome\\\\Application\\\\ch****.exe
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// The status of the application.
	//
	// Valid value:
	//
	// 	- Idle: The application is installed in the cloud computer but is not running.
	//
	// 	- Running: The application has been installed in the cloud computer and is running.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGuestApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetApplicationVersion() *string {
	return s.ApplicationVersion
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetCpuPercent() *float64 {
	return s.CpuPercent
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetGpuPercent() *float64 {
	return s.GpuPercent
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetIconUrl() *string {
	return s.IconUrl
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetIoSpeed() *float64 {
	return s.IoSpeed
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetMemPercent() *float64 {
	return s.MemPercent
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetPid() *int32 {
	return s.Pid
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetProcessData() []*DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	return s.ProcessData
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *DescribeGuestApplicationsResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetApplicationName(v string) *DescribeGuestApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetApplicationVersion(v string) *DescribeGuestApplicationsResponseBodyApplications {
	s.ApplicationVersion = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetCpuPercent(v float64) *DescribeGuestApplicationsResponseBodyApplications {
	s.CpuPercent = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetGpuPercent(v float64) *DescribeGuestApplicationsResponseBodyApplications {
	s.GpuPercent = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetIconUrl(v string) *DescribeGuestApplicationsResponseBodyApplications {
	s.IconUrl = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetIoSpeed(v float64) *DescribeGuestApplicationsResponseBodyApplications {
	s.IoSpeed = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetMemPercent(v float64) *DescribeGuestApplicationsResponseBodyApplications {
	s.MemPercent = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetPid(v int32) *DescribeGuestApplicationsResponseBodyApplications {
	s.Pid = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetProcessData(v []*DescribeGuestApplicationsResponseBodyApplicationsProcessData) *DescribeGuestApplicationsResponseBodyApplications {
	s.ProcessData = v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetProcessPath(v string) *DescribeGuestApplicationsResponseBodyApplications {
	s.ProcessPath = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) SetStatus(v string) *DescribeGuestApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}

type DescribeGuestApplicationsResponseBodyApplicationsProcessData struct {
	// The application name.
	//
	// example:
	//
	// Google Chrome
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application version.
	//
	// example:
	//
	// 115.0.5790.110
	ApplicationVersion *string `json:"ApplicationVersion,omitempty" xml:"ApplicationVersion,omitempty"`
	// The CPU utilization (%).
	//
	// example:
	//
	// 89
	CpuPercent *float64 `json:"CpuPercent,omitempty" xml:"CpuPercent,omitempty"`
	// The GPU usage (%).
	//
	// example:
	//
	// 15
	GpuPercent *float64 `json:"GpuPercent,omitempty" xml:"GpuPercent,omitempty"`
	// The I/O read and write performance. Unit: byte/s.
	//
	// example:
	//
	// 124906.0
	Iospeed *float64 `json:"Iospeed,omitempty" xml:"Iospeed,omitempty"`
	// The memory usage (%).
	//
	// example:
	//
	// 34
	MemPercent *float64 `json:"MemPercent,omitempty" xml:"MemPercent,omitempty"`
	// The PID.
	//
	// example:
	//
	// 1357
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The path to the process.
	//
	// example:
	//
	// C:\\\\Program Files\\\\Google\\\\Chrome\\\\Application\\\\ch****.exe
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
}

func (s DescribeGuestApplicationsResponseBodyApplicationsProcessData) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestApplicationsResponseBodyApplicationsProcessData) GoString() string {
	return s.String()
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetApplicationVersion() *string {
	return s.ApplicationVersion
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetCpuPercent() *float64 {
	return s.CpuPercent
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetGpuPercent() *float64 {
	return s.GpuPercent
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetIospeed() *float64 {
	return s.Iospeed
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetMemPercent() *float64 {
	return s.MemPercent
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetPid() *int32 {
	return s.Pid
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetApplicationName(v string) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.ApplicationName = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetApplicationVersion(v string) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.ApplicationVersion = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetCpuPercent(v float64) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.CpuPercent = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetGpuPercent(v float64) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.GpuPercent = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetIospeed(v float64) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.Iospeed = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetMemPercent(v float64) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.MemPercent = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetPid(v int32) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.Pid = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) SetProcessPath(v string) *DescribeGuestApplicationsResponseBodyApplicationsProcessData {
	s.ProcessPath = &v
	return s
}

func (s *DescribeGuestApplicationsResponseBodyApplicationsProcessData) Validate() error {
	return dara.Validate(s)
}
