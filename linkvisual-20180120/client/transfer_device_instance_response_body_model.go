// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDeviceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TransferDeviceInstanceResponseBody
	GetCode() *string
	SetData(v *TransferDeviceInstanceResponseBodyData) *TransferDeviceInstanceResponseBody
	GetData() *TransferDeviceInstanceResponseBodyData
	SetErrorMessage(v string) *TransferDeviceInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *TransferDeviceInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferDeviceInstanceResponseBody
	GetSuccess() *bool
}

type TransferDeviceInstanceResponseBody struct {
	// The response code:
	//
	// - **200**: The call was successful.
	//
	// - Other values: The call failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result.
	Data *TransferDeviceInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// The source instance and the target instance cannot be the same
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 27AF2E38-12ED-4F67-BD20-36E812CD7A42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferDeviceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferDeviceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferDeviceInstanceResponseBody) GetData() *TransferDeviceInstanceResponseBodyData {
	return s.Data
}

func (s *TransferDeviceInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TransferDeviceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferDeviceInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferDeviceInstanceResponseBody) SetCode(v string) *TransferDeviceInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetData(v *TransferDeviceInstanceResponseBodyData) *TransferDeviceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetErrorMessage(v string) *TransferDeviceInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetRequestId(v string) *TransferDeviceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) SetSuccess(v bool) *TransferDeviceInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *TransferDeviceInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TransferDeviceInstanceResponseBodyData struct {
	// The list of devices that failed to be transferred.
	FailedList []*TransferDeviceInstanceResponseBodyDataFailedList `json:"FailedList,omitempty" xml:"FailedList,omitempty" type:"Repeated"`
	// The list of devices that were successfully transferred.
	SuccessList []*TransferDeviceInstanceResponseBodyDataSuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s TransferDeviceInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TransferDeviceInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBodyData) GetFailedList() []*TransferDeviceInstanceResponseBodyDataFailedList {
	return s.FailedList
}

func (s *TransferDeviceInstanceResponseBodyData) GetSuccessList() []*TransferDeviceInstanceResponseBodyDataSuccessList {
	return s.SuccessList
}

func (s *TransferDeviceInstanceResponseBodyData) SetFailedList(v []*TransferDeviceInstanceResponseBodyDataFailedList) *TransferDeviceInstanceResponseBodyData {
	s.FailedList = v
	return s
}

func (s *TransferDeviceInstanceResponseBodyData) SetSuccessList(v []*TransferDeviceInstanceResponseBodyDataSuccessList) *TransferDeviceInstanceResponseBodyData {
	s.SuccessList = v
	return s
}

func (s *TransferDeviceInstanceResponseBodyData) Validate() error {
	if s.FailedList != nil {
		for _, item := range s.FailedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessList != nil {
		for _, item := range s.SuccessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TransferDeviceInstanceResponseBodyDataFailedList struct {
	// The name of the device.
	//
	// example:
	//
	// e5jdfalekj****
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The description of the failed transfer.
	//
	// example:
	//
	// The device has already been transferred or does not belong to the source instance
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TransferDeviceInstanceResponseBodyDataFailedList) String() string {
	return dara.Prettify(s)
}

func (s TransferDeviceInstanceResponseBodyDataFailedList) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) GetMessage() *string {
	return s.Message
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) SetDeviceName(v string) *TransferDeviceInstanceResponseBodyDataFailedList {
	s.DeviceName = &v
	return s
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) SetMessage(v string) *TransferDeviceInstanceResponseBodyDataFailedList {
	s.Message = &v
	return s
}

func (s *TransferDeviceInstanceResponseBodyDataFailedList) Validate() error {
	return dara.Validate(s)
}

type TransferDeviceInstanceResponseBodyDataSuccessList struct {
	// The name of the device.
	//
	// example:
	//
	// e598798j****
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The description of the successful transfer.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TransferDeviceInstanceResponseBodyDataSuccessList) String() string {
	return dara.Prettify(s)
}

func (s TransferDeviceInstanceResponseBodyDataSuccessList) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) GetMessage() *string {
	return s.Message
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) SetDeviceName(v string) *TransferDeviceInstanceResponseBodyDataSuccessList {
	s.DeviceName = &v
	return s
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) SetMessage(v string) *TransferDeviceInstanceResponseBodyDataSuccessList {
	s.Message = &v
	return s
}

func (s *TransferDeviceInstanceResponseBodyDataSuccessList) Validate() error {
	return dara.Validate(s)
}
