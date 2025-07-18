// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportUserDevicesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppStatuses(v []*string) *ExportUserDevicesRequest
  GetAppStatuses() []*string 
  SetDepartment(v string) *ExportUserDevicesRequest
  GetDepartment() *string 
  SetDeviceBelong(v string) *ExportUserDevicesRequest
  GetDeviceBelong() *string 
  SetDeviceStatuses(v []*string) *ExportUserDevicesRequest
  GetDeviceStatuses() []*string 
  SetDeviceTags(v []*string) *ExportUserDevicesRequest
  GetDeviceTags() []*string 
  SetDeviceTypes(v []*string) *ExportUserDevicesRequest
  GetDeviceTypes() []*string 
  SetDlpStatuses(v []*string) *ExportUserDevicesRequest
  GetDlpStatuses() []*string 
  SetHostname(v string) *ExportUserDevicesRequest
  GetHostname() *string 
  SetIaStatuses(v []*string) *ExportUserDevicesRequest
  GetIaStatuses() []*string 
  SetMac(v string) *ExportUserDevicesRequest
  GetMac() *string 
  SetNacStatuses(v []*string) *ExportUserDevicesRequest
  GetNacStatuses() []*string 
  SetPaStatuses(v []*string) *ExportUserDevicesRequest
  GetPaStatuses() []*string 
  SetSaseUserId(v string) *ExportUserDevicesRequest
  GetSaseUserId() *string 
  SetSharingStatus(v bool) *ExportUserDevicesRequest
  GetSharingStatus() *bool 
  SetUsername(v string) *ExportUserDevicesRequest
  GetUsername() *string 
}

type ExportUserDevicesRequest struct {
  AppStatuses []*string `json:"AppStatuses,omitempty" xml:"AppStatuses,omitempty" type:"Repeated"`
  Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
  // example:
  // 
  // Company
  DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
  DeviceStatuses []*string `json:"DeviceStatuses,omitempty" xml:"DeviceStatuses,omitempty" type:"Repeated"`
  DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
  DeviceTypes []*string `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
  DlpStatuses []*string `json:"DlpStatuses,omitempty" xml:"DlpStatuses,omitempty" type:"Repeated"`
  // example:
  // 
  // win10-64bit
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
  IaStatuses []*string `json:"IaStatuses,omitempty" xml:"IaStatuses,omitempty" type:"Repeated"`
  // example:
  // 
  // 00:16:XX:XX:7c:46
  Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
  NacStatuses []*string `json:"NacStatuses,omitempty" xml:"NacStatuses,omitempty" type:"Repeated"`
  PaStatuses []*string `json:"PaStatuses,omitempty" xml:"PaStatuses,omitempty" type:"Repeated"`
  // example:
  // 
  // su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
  SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
  // example:
  // 
  // true
  SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
  Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ExportUserDevicesRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportUserDevicesRequest) GoString() string {
  return s.String()
}

func (s *ExportUserDevicesRequest) GetAppStatuses() []*string  {
  return s.AppStatuses
}

func (s *ExportUserDevicesRequest) GetDepartment() *string  {
  return s.Department
}

func (s *ExportUserDevicesRequest) GetDeviceBelong() *string  {
  return s.DeviceBelong
}

func (s *ExportUserDevicesRequest) GetDeviceStatuses() []*string  {
  return s.DeviceStatuses
}

func (s *ExportUserDevicesRequest) GetDeviceTags() []*string  {
  return s.DeviceTags
}

func (s *ExportUserDevicesRequest) GetDeviceTypes() []*string  {
  return s.DeviceTypes
}

func (s *ExportUserDevicesRequest) GetDlpStatuses() []*string  {
  return s.DlpStatuses
}

func (s *ExportUserDevicesRequest) GetHostname() *string  {
  return s.Hostname
}

func (s *ExportUserDevicesRequest) GetIaStatuses() []*string  {
  return s.IaStatuses
}

func (s *ExportUserDevicesRequest) GetMac() *string  {
  return s.Mac
}

func (s *ExportUserDevicesRequest) GetNacStatuses() []*string  {
  return s.NacStatuses
}

func (s *ExportUserDevicesRequest) GetPaStatuses() []*string  {
  return s.PaStatuses
}

func (s *ExportUserDevicesRequest) GetSaseUserId() *string  {
  return s.SaseUserId
}

func (s *ExportUserDevicesRequest) GetSharingStatus() *bool  {
  return s.SharingStatus
}

func (s *ExportUserDevicesRequest) GetUsername() *string  {
  return s.Username
}

func (s *ExportUserDevicesRequest) SetAppStatuses(v []*string) *ExportUserDevicesRequest {
  s.AppStatuses = v
  return s
}

func (s *ExportUserDevicesRequest) SetDepartment(v string) *ExportUserDevicesRequest {
  s.Department = &v
  return s
}

func (s *ExportUserDevicesRequest) SetDeviceBelong(v string) *ExportUserDevicesRequest {
  s.DeviceBelong = &v
  return s
}

func (s *ExportUserDevicesRequest) SetDeviceStatuses(v []*string) *ExportUserDevicesRequest {
  s.DeviceStatuses = v
  return s
}

func (s *ExportUserDevicesRequest) SetDeviceTags(v []*string) *ExportUserDevicesRequest {
  s.DeviceTags = v
  return s
}

func (s *ExportUserDevicesRequest) SetDeviceTypes(v []*string) *ExportUserDevicesRequest {
  s.DeviceTypes = v
  return s
}

func (s *ExportUserDevicesRequest) SetDlpStatuses(v []*string) *ExportUserDevicesRequest {
  s.DlpStatuses = v
  return s
}

func (s *ExportUserDevicesRequest) SetHostname(v string) *ExportUserDevicesRequest {
  s.Hostname = &v
  return s
}

func (s *ExportUserDevicesRequest) SetIaStatuses(v []*string) *ExportUserDevicesRequest {
  s.IaStatuses = v
  return s
}

func (s *ExportUserDevicesRequest) SetMac(v string) *ExportUserDevicesRequest {
  s.Mac = &v
  return s
}

func (s *ExportUserDevicesRequest) SetNacStatuses(v []*string) *ExportUserDevicesRequest {
  s.NacStatuses = v
  return s
}

func (s *ExportUserDevicesRequest) SetPaStatuses(v []*string) *ExportUserDevicesRequest {
  s.PaStatuses = v
  return s
}

func (s *ExportUserDevicesRequest) SetSaseUserId(v string) *ExportUserDevicesRequest {
  s.SaseUserId = &v
  return s
}

func (s *ExportUserDevicesRequest) SetSharingStatus(v bool) *ExportUserDevicesRequest {
  s.SharingStatus = &v
  return s
}

func (s *ExportUserDevicesRequest) SetUsername(v string) *ExportUserDevicesRequest {
  s.Username = &v
  return s
}

func (s *ExportUserDevicesRequest) Validate() error {
  return dara.Validate(s)
}

