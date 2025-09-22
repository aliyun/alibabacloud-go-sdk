// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceType(v string) *PushDeviceDataRequest
	GetDeviceType() *string
	SetDevices(v []*PushDeviceDataRequestDevices) *PushDeviceDataRequest
	GetDevices() []*PushDeviceDataRequestDevices
}

type PushDeviceDataRequest struct {
	// The type of the device. [View device type definitions](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/Deviceappendixes-en.pdf)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// List of devices to which data is pushed.
	//
	// This parameter is required.
	Devices []*PushDeviceDataRequestDevices `json:"devices,omitempty" xml:"devices,omitempty" type:"Repeated"`
}

func (s PushDeviceDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceDataRequest) GoString() string {
	return s.String()
}

func (s *PushDeviceDataRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *PushDeviceDataRequest) GetDevices() []*PushDeviceDataRequestDevices {
	return s.Devices
}

func (s *PushDeviceDataRequest) SetDeviceType(v string) *PushDeviceDataRequest {
	s.DeviceType = &v
	return s
}

func (s *PushDeviceDataRequest) SetDevices(v []*PushDeviceDataRequestDevices) *PushDeviceDataRequest {
	s.Devices = v
	return s
}

func (s *PushDeviceDataRequest) Validate() error {
	return dara.Validate(s)
}

type PushDeviceDataRequestDevices struct {
	// Measuring point information To avoid accuracy problems, the measurement point data is uniformly transmitted to the string. The function of missing required fields cannot be used normally. Some functions may be affected due to the lack of recommend fields. For details, please refer to the notes of equipment measuring points in the appendix. [Reference Point Definition](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/Deviceappendixes-en.pdf
	//
	// )
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 			"dp_imp": "329.0",
	//
	// 			"F": "148.0",
	//
	// 			"eq_imp": "363.0",
	//
	// 			"Ep_imp_1": "128.0",
	//
	// 			"Ep_imp_2": "157.0",
	//
	// 			"Ua": "226.0",
	//
	// 			"Ub": "285.0",
	//
	// 			"Ep_imp": "325.0",
	//
	// 			"Uc": "342.0",
	//
	// 			"Ep_imp_3": "109.0",
	//
	// 			"Ep_imp_4": "94.0",
	//
	// 			"P": "514.0",
	//
	// 			"Pa": "443.0",
	//
	// 			"Q": "265.0",
	//
	// 			"dp_exp": "261.0",
	//
	// 			"eq_exp": "399.0",
	//
	// 			"COSQ": "223.0",
	//
	// 			"Ia": "240.0",
	//
	// 			"Ib": "216.0",
	//
	// 			"Ic": "229.0",
	//
	// 			"Ep_exp": "115.0",
	//
	// 			"VdisPer": "120.0"
	//
	// 		}
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// If the deviceType parameter is set to 12, 13, or 17, you must set the system_id parameter. The field name is still device_id. If the deviceType parameter is set to 15 or 16, no Other situations will be transmitted.
	//
	// This parameter is required.
	//
	// example:
	//
	// device_code_xxx
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// Data generation time of measuring point.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-08 18:40:00
	RecordTime *string `json:"recordTime,omitempty" xml:"recordTime,omitempty"`
}

func (s PushDeviceDataRequestDevices) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceDataRequestDevices) GoString() string {
	return s.String()
}

func (s *PushDeviceDataRequestDevices) GetData() map[string]interface{} {
	return s.Data
}

func (s *PushDeviceDataRequestDevices) GetDeviceId() *string {
	return s.DeviceId
}

func (s *PushDeviceDataRequestDevices) GetRecordTime() *string {
	return s.RecordTime
}

func (s *PushDeviceDataRequestDevices) SetData(v map[string]interface{}) *PushDeviceDataRequestDevices {
	s.Data = v
	return s
}

func (s *PushDeviceDataRequestDevices) SetDeviceId(v string) *PushDeviceDataRequestDevices {
	s.DeviceId = &v
	return s
}

func (s *PushDeviceDataRequestDevices) SetRecordTime(v string) *PushDeviceDataRequestDevices {
	s.RecordTime = &v
	return s
}

func (s *PushDeviceDataRequestDevices) Validate() error {
	return dara.Validate(s)
}
