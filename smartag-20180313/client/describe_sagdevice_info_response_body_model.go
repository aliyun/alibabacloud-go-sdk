// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSAGDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControllerState(v string) *DescribeSAGDeviceInfoResponseBody
	GetControllerState() *string
	SetLastConnectedControllerTime(v string) *DescribeSAGDeviceInfoResponseBody
	GetLastConnectedControllerTime() *string
	SetRequestId(v string) *DescribeSAGDeviceInfoResponseBody
	GetRequestId() *string
	SetResettableStatus(v string) *DescribeSAGDeviceInfoResponseBody
	GetResettableStatus() *string
	SetServiceIP(v string) *DescribeSAGDeviceInfoResponseBody
	GetServiceIP() *string
	SetSmartAGType(v string) *DescribeSAGDeviceInfoResponseBody
	GetSmartAGType() *string
	SetStartupTime(v string) *DescribeSAGDeviceInfoResponseBody
	GetStartupTime() *string
	SetSynStatus(v string) *DescribeSAGDeviceInfoResponseBody
	GetSynStatus() *string
	SetVersion(v string) *DescribeSAGDeviceInfoResponseBody
	GetVersion() *string
	SetVpnState(v string) *DescribeSAGDeviceInfoResponseBody
	GetVpnState() *string
}

type DescribeSAGDeviceInfoResponseBody struct {
	// The control status of the Smart Access Gateway device. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **Abnormal**: Abnormal.
	//
	// example:
	//
	// Normal
	ControllerState *string `json:"ControllerState,omitempty" xml:"ControllerState,omitempty"`
	// The most recent time when the Smart Access Gateway device connected to the controller.
	//
	// example:
	//
	// 2021-07-14 00:27:48
	LastConnectedControllerTime *string `json:"LastConnectedControllerTime,omitempty" xml:"LastConnectedControllerTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6A991F4-F533-1627-8144-B64E01C5EE85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the reset button of the Smart Access Gateway device is enabled. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	ResettableStatus *string `json:"ResettableStatus,omitempty" xml:"ResettableStatus,omitempty"`
	// The service IP address of the Smart Access Gateway device.
	//
	// example:
	//
	// 42.XX.XX.151
	ServiceIP *string `json:"ServiceIP,omitempty" xml:"ServiceIP,omitempty"`
	// The type of the Smart Access Gateway device. Valid values:
	//
	// - **sag-100wm**.
	//
	// - **sag-1000**.
	//
	// example:
	//
	// sag-100wm
	SmartAGType *string `json:"SmartAGType,omitempty" xml:"SmartAGType,omitempty"`
	// The startup time of the Smart Access Gateway device.
	//
	// example:
	//
	// 2021-06-15 17:33:43
	StartupTime *string `json:"StartupTime,omitempty" xml:"StartupTime,omitempty"`
	// The synchronization status between the local Smart Access Gateway device and the cloud. Valid values:
	//
	// - **Synchronized**: Synchronization is complete.
	//
	// - **Unsynchronized**: Not synchronized.
	//
	// - **Synchronizing**: Synchronization is in progress.
	//
	// example:
	//
	// Unsynchronized
	SynStatus *string `json:"SynStatus,omitempty" xml:"SynStatus,omitempty"`
	// The software version that runs on the Smart Access Gateway device.
	//
	// example:
	//
	// 2.3.1.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The VPN connection status of the Smart Access Gateway device. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **Abnormal**: Abnormal.
	//
	// example:
	//
	// Abnormal
	VpnState *string `json:"VpnState,omitempty" xml:"VpnState,omitempty"`
}

func (s DescribeSAGDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSAGDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSAGDeviceInfoResponseBody) GetControllerState() *string {
	return s.ControllerState
}

func (s *DescribeSAGDeviceInfoResponseBody) GetLastConnectedControllerTime() *string {
	return s.LastConnectedControllerTime
}

func (s *DescribeSAGDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSAGDeviceInfoResponseBody) GetResettableStatus() *string {
	return s.ResettableStatus
}

func (s *DescribeSAGDeviceInfoResponseBody) GetServiceIP() *string {
	return s.ServiceIP
}

func (s *DescribeSAGDeviceInfoResponseBody) GetSmartAGType() *string {
	return s.SmartAGType
}

func (s *DescribeSAGDeviceInfoResponseBody) GetStartupTime() *string {
	return s.StartupTime
}

func (s *DescribeSAGDeviceInfoResponseBody) GetSynStatus() *string {
	return s.SynStatus
}

func (s *DescribeSAGDeviceInfoResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeSAGDeviceInfoResponseBody) GetVpnState() *string {
	return s.VpnState
}

func (s *DescribeSAGDeviceInfoResponseBody) SetControllerState(v string) *DescribeSAGDeviceInfoResponseBody {
	s.ControllerState = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetLastConnectedControllerTime(v string) *DescribeSAGDeviceInfoResponseBody {
	s.LastConnectedControllerTime = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetRequestId(v string) *DescribeSAGDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetResettableStatus(v string) *DescribeSAGDeviceInfoResponseBody {
	s.ResettableStatus = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetServiceIP(v string) *DescribeSAGDeviceInfoResponseBody {
	s.ServiceIP = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetSmartAGType(v string) *DescribeSAGDeviceInfoResponseBody {
	s.SmartAGType = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetStartupTime(v string) *DescribeSAGDeviceInfoResponseBody {
	s.StartupTime = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetSynStatus(v string) *DescribeSAGDeviceInfoResponseBody {
	s.SynStatus = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetVersion(v string) *DescribeSAGDeviceInfoResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) SetVpnState(v string) *DescribeSAGDeviceInfoResponseBody {
	s.VpnState = &v
	return s
}

func (s *DescribeSAGDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
