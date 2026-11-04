// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDeviceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceNameList(v []*string) *TransferDeviceInstanceRequest
	GetDeviceNameList() []*string
	SetProductKey(v string) *TransferDeviceInstanceRequest
	GetProductKey() *string
	SetSourceInstanceId(v string) *TransferDeviceInstanceRequest
	GetSourceInstanceId() *string
	SetTargetInstanceId(v string) *TransferDeviceInstanceRequest
	GetTargetInstanceId() *string
}

type TransferDeviceInstanceRequest struct {
	// The list of names of the devices to be transferred. You can specify up to 100 device names.
	//
	// This parameter is required.
	//
	// example:
	//
	// RepeatList
	DeviceNameList []*string `json:"DeviceNameList,omitempty" xml:"DeviceNameList,omitempty" type:"Repeated"`
	// The ProductKey of the product to which the devices belong. IoT Platform generates this unique identifier when the product is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The ID of the source instance to which the devices belong:
	//
	// - Public instance ID: iotx-oxssharez200.
	//
	//     > Only the China (Shanghai) region is supported. Enterprise instance ID is iotx-oxssharez200.
	//
	// - Enterprise Edition instance ID: Log on to the IoT Platform console. On the **Instance Overview*	- page, click Enterprise instance name. On the **Instance Details*	- page, view the **Instance ID*	- in the **Basic Information*	- section.
	//
	// example:
	//
	// iotx-oxssharez200
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The ID of the target instance to which the devices are transferred. Only Enterprise Edition instances are supported as target instances.
	//
	// You can log on to the IoT Platform console. On the **Instance Overview*	- page, click the instance name. On the **Instance Details*	- page, view the **Instance ID*	- in the **Basic Information*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-060***
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
}

func (s TransferDeviceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferDeviceInstanceRequest) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceRequest) GetDeviceNameList() []*string {
	return s.DeviceNameList
}

func (s *TransferDeviceInstanceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *TransferDeviceInstanceRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *TransferDeviceInstanceRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *TransferDeviceInstanceRequest) SetDeviceNameList(v []*string) *TransferDeviceInstanceRequest {
	s.DeviceNameList = v
	return s
}

func (s *TransferDeviceInstanceRequest) SetProductKey(v string) *TransferDeviceInstanceRequest {
	s.ProductKey = &v
	return s
}

func (s *TransferDeviceInstanceRequest) SetSourceInstanceId(v string) *TransferDeviceInstanceRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *TransferDeviceInstanceRequest) SetTargetInstanceId(v string) *TransferDeviceInstanceRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *TransferDeviceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
