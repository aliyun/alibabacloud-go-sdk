// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearFaceDeviceDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *ClearFaceDeviceDBRequest
	GetDeviceName() *string
	SetIotInstanceId(v string) *ClearFaceDeviceDBRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *ClearFaceDeviceDBRequest
	GetIsolationId() *string
	SetProductKey(v string) *ClearFaceDeviceDBRequest
	GetProductKey() *string
}

type ClearFaceDeviceDBRequest struct {
	// The name of the device corresponding to the V-series edge all-in-one machine. On the **Basic Information*	- tab of the **Host Details*	- page for the V-series edge all-in-one machine in the Edge Computing console, find the **Video Gateway*	- information and obtain the **DeviceName*	- value.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDevice_1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If an ID value exists, you must pass this ID value. Otherwise, the call fails.
	//
	// > If a public instance does not have an ID value, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot_instc_pu****_c*-v64********
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The business isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The unique identifier of the product to which the device corresponding to the V-series edge all-in-one machine belongs. On the **Basic Information*	- tab of the **Host Details*	- page for the V-series edge all-in-one machine in the Edge Computing console, find the **Video Gateway*	- information and obtain the **ProductKey*	- value.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1Wmy******
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ClearFaceDeviceDBRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearFaceDeviceDBRequest) GoString() string {
	return s.String()
}

func (s *ClearFaceDeviceDBRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ClearFaceDeviceDBRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *ClearFaceDeviceDBRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *ClearFaceDeviceDBRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *ClearFaceDeviceDBRequest) SetDeviceName(v string) *ClearFaceDeviceDBRequest {
	s.DeviceName = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetIotInstanceId(v string) *ClearFaceDeviceDBRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetIsolationId(v string) *ClearFaceDeviceDBRequest {
	s.IsolationId = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) SetProductKey(v string) *ClearFaceDeviceDBRequest {
	s.ProductKey = &v
	return s
}

func (s *ClearFaceDeviceDBRequest) Validate() error {
	return dara.Validate(s)
}
