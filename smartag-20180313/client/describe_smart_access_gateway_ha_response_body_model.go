// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDeviceId(v string) *DescribeSmartAccessGatewayHaResponseBody
	GetBackupDeviceId() *string
	SetDeviceLevelBackupState(v string) *DescribeSmartAccessGatewayHaResponseBody
	GetDeviceLevelBackupState() *string
	SetDeviceLevelBackupType(v string) *DescribeSmartAccessGatewayHaResponseBody
	GetDeviceLevelBackupType() *string
	SetLinkBackupInfoList(v *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) *DescribeSmartAccessGatewayHaResponseBody
	GetLinkBackupInfoList() *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList
	SetMainDeviceId(v string) *DescribeSmartAccessGatewayHaResponseBody
	GetMainDeviceId() *string
	SetRequestId(v string) *DescribeSmartAccessGatewayHaResponseBody
	GetRequestId() *string
	SetSmartAGId(v string) *DescribeSmartAccessGatewayHaResponseBody
	GetSmartAGId() *string
}

type DescribeSmartAccessGatewayHaResponseBody struct {
	// The serial number of the standby SAG device.
	//
	// example:
	//
	// sag11axxxx
	BackupDeviceId *string `json:"BackupDeviceId,omitempty" xml:"BackupDeviceId,omitempty"`
	// Indicates whether device-based HA is enabled. Valid values:
	//
	// 	- **ON**: enabled
	//
	// 	- **OFF**: disabled
	//
	// example:
	//
	// OFF
	DeviceLevelBackupState *string `json:"DeviceLevelBackupState,omitempty" xml:"DeviceLevelBackupState,omitempty"`
	// The deployment mode of the SAG devices that have HA enabled. Valid values:
	//
	// 	- **warm_backup**: active-active mode.
	//
	// 	- **cold_backup**: active-standby mode.
	//
	// 	- **no_backup**: HA is disabled.
	//
	// example:
	//
	// cold_backup
	DeviceLevelBackupType *string                                                     `json:"DeviceLevelBackupType,omitempty" xml:"DeviceLevelBackupType,omitempty"`
	LinkBackupInfoList    *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList `json:"LinkBackupInfoList,omitempty" xml:"LinkBackupInfoList,omitempty" type:"Struct"`
	// The serial number of the active SAG device.
	//
	// example:
	//
	// sag11axxxx
	MainDeviceId *string `json:"MainDeviceId,omitempty" xml:"MainDeviceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 582FE511-FEFE-42BC-BBF4-4F8ECF92Exxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-i8mogwi9kisigc3xxxx
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeSmartAccessGatewayHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayHaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetBackupDeviceId() *string {
	return s.BackupDeviceId
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetDeviceLevelBackupState() *string {
	return s.DeviceLevelBackupState
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetDeviceLevelBackupType() *string {
	return s.DeviceLevelBackupType
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetLinkBackupInfoList() *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList {
	return s.LinkBackupInfoList
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetMainDeviceId() *string {
	return s.MainDeviceId
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmartAccessGatewayHaResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetBackupDeviceId(v string) *DescribeSmartAccessGatewayHaResponseBody {
	s.BackupDeviceId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetDeviceLevelBackupState(v string) *DescribeSmartAccessGatewayHaResponseBody {
	s.DeviceLevelBackupState = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetDeviceLevelBackupType(v string) *DescribeSmartAccessGatewayHaResponseBody {
	s.DeviceLevelBackupType = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetLinkBackupInfoList(v *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) *DescribeSmartAccessGatewayHaResponseBody {
	s.LinkBackupInfoList = v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetMainDeviceId(v string) *DescribeSmartAccessGatewayHaResponseBody {
	s.MainDeviceId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetRequestId(v string) *DescribeSmartAccessGatewayHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) SetSmartAGId(v string) *DescribeSmartAccessGatewayHaResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBody) Validate() error {
	if s.LinkBackupInfoList != nil {
		if err := s.LinkBackupInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList struct {
	LinkBackupInfoList []*DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList `json:"LinkBackupInfoList,omitempty" xml:"LinkBackupInfoList,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) GetLinkBackupInfoList() []*DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	return s.LinkBackupInfoList
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) SetLinkBackupInfoList(v []*DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList {
	s.LinkBackupInfoList = v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoList) Validate() error {
	if s.LinkBackupInfoList != nil {
		for _, item := range s.LinkBackupInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList struct {
	BackupLinkId         *string `json:"BackupLinkId,omitempty" xml:"BackupLinkId,omitempty"`
	BackupLinkState      *string `json:"BackupLinkState,omitempty" xml:"BackupLinkState,omitempty"`
	LinkLevelBackupState *string `json:"LinkLevelBackupState,omitempty" xml:"LinkLevelBackupState,omitempty"`
	LinkLevelBackupType  *string `json:"LinkLevelBackupType,omitempty" xml:"LinkLevelBackupType,omitempty"`
	MainLinkId           *string `json:"MainLinkId,omitempty" xml:"MainLinkId,omitempty"`
	MainLinkState        *string `json:"MainLinkState,omitempty" xml:"MainLinkState,omitempty"`
}

func (s DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GetBackupLinkId() *string {
	return s.BackupLinkId
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GetBackupLinkState() *string {
	return s.BackupLinkState
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GetLinkLevelBackupState() *string {
	return s.LinkLevelBackupState
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GetLinkLevelBackupType() *string {
	return s.LinkLevelBackupType
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GetMainLinkId() *string {
	return s.MainLinkId
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) GetMainLinkState() *string {
	return s.MainLinkState
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) SetBackupLinkId(v string) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	s.BackupLinkId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) SetBackupLinkState(v string) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	s.BackupLinkState = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) SetLinkLevelBackupState(v string) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	s.LinkLevelBackupState = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) SetLinkLevelBackupType(v string) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	s.LinkLevelBackupType = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) SetMainLinkId(v string) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	s.MainLinkId = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) SetMainLinkState(v string) *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList {
	s.MainLinkState = &v
	return s
}

func (s *DescribeSmartAccessGatewayHaResponseBodyLinkBackupInfoListLinkBackupInfoList) Validate() error {
	return dara.Validate(s)
}
