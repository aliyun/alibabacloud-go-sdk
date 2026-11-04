// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocalFileUploadJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryLocalFileUploadJobResponseBody
	GetCode() *string
	SetData(v *QueryLocalFileUploadJobResponseBodyData) *QueryLocalFileUploadJobResponseBody
	GetData() *QueryLocalFileUploadJobResponseBodyData
	SetErrorMessage(v string) *QueryLocalFileUploadJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryLocalFileUploadJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryLocalFileUploadJobResponseBody
	GetSuccess() *bool
}

type QueryLocalFileUploadJobResponseBody struct {
	// The response code:
	//
	// - **200**: The call is successful.
	//
	// - Other values: An error occurred. For more information about error codes, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *QueryLocalFileUploadJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 449ECE75-A36A-4ACE-B253-147C78C440C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
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

func (s QueryLocalFileUploadJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryLocalFileUploadJobResponseBody) GetData() *QueryLocalFileUploadJobResponseBodyData {
	return s.Data
}

func (s *QueryLocalFileUploadJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryLocalFileUploadJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLocalFileUploadJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryLocalFileUploadJobResponseBody) SetCode(v string) *QueryLocalFileUploadJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetData(v *QueryLocalFileUploadJobResponseBodyData) *QueryLocalFileUploadJobResponseBody {
	s.Data = v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetErrorMessage(v string) *QueryLocalFileUploadJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetRequestId(v string) *QueryLocalFileUploadJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) SetSuccess(v bool) *QueryLocalFileUploadJobResponseBody {
	s.Success = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryLocalFileUploadJobResponseBodyData struct {
	// The task list.
	ResultList []*QueryLocalFileUploadJobResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The status of the video recording upload task. Valid values:
	//
	// - 0: The task is being initialized.
	//
	// - 1: The task is running.
	//
	// - 2: The task is completed.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBodyData) GetResultList() []*QueryLocalFileUploadJobResponseBodyDataResultList {
	return s.ResultList
}

func (s *QueryLocalFileUploadJobResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryLocalFileUploadJobResponseBodyData) SetResultList(v []*QueryLocalFileUploadJobResponseBodyDataResultList) *QueryLocalFileUploadJobResponseBodyData {
	s.ResultList = v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyData) SetStatus(v int32) *QueryLocalFileUploadJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyData) Validate() error {
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryLocalFileUploadJobResponseBodyDataResultList struct {
	// The status code of the video recording upload task. Valid values:
	//
	// - 200: Normal.
	//
	// - 9201: The device is offline. Bring the device online and try again.
	//
	// - 9409: No cloud storage quota is available. Check the cloud storage quota and try again.
	//
	// - 9417: A device, network, or other exception occurred. Try again, or check the device and network status.
	//
	// - 9435: The video recording file does not exist. Make sure that the video recording file is stored locally on the video device and try again.
	//
	// example:
	//
	// 9201
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The DeviceName of the video device.
	//
	// example:
	//
	// gb_hik_IPC
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The list of video recording files.
	FileList []*QueryLocalFileUploadJobResponseBodyDataResultListFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// The ID of the video device. The ID is issued by IoT Platform and uniquely identifies the device.
	//
	// example:
	//
	// IQ9mO4tvMnAj****QUdcg2n400
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The ProductKey of the product to which the video device belongs.
	//
	// example:
	//
	// g2****Q6ggK
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The end time of the video recording upload task.
	//
	// example:
	//
	// 1634746054
	SlotEndTime *int32 `json:"SlotEndTime,omitempty" xml:"SlotEndTime,omitempty"`
	// The start time of the video recording upload task.
	//
	// example:
	//
	// 1634745754
	SlotStartTime *int32 `json:"SlotStartTime,omitempty" xml:"SlotStartTime,omitempty"`
	// The execution status of the video recording upload task within the time period specified by **SlotStartTime*	- and **SlotEndTime**. Valid values:
	//
	// - 0: The task is being initialized.
	//
	// - 1: The task is running.
	//
	// - 2: The task is completed.
	//
	// - 3: An exception occurred during task execution.
	//
	// - 4: The task is manually terminated.
	//
	// example:
	//
	// 3
	SlotStatus *int32 `json:"SlotStatus,omitempty" xml:"SlotStatus,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetCode() *int32 {
	return s.Code
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetFileList() []*QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	return s.FileList
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetIotId() *string {
	return s.IotId
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetSlotEndTime() *int32 {
	return s.SlotEndTime
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetSlotStartTime() *int32 {
	return s.SlotStartTime
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) GetSlotStatus() *int32 {
	return s.SlotStatus
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetCode(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.Code = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetDeviceName(v string) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.DeviceName = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetFileList(v []*QueryLocalFileUploadJobResponseBodyDataResultListFileList) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.FileList = v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetIotId(v string) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.IotId = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetProductKey(v string) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.ProductKey = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetSlotEndTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.SlotEndTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetSlotStartTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.SlotStartTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) SetSlotStatus(v int32) *QueryLocalFileUploadJobResponseBodyDataResultList {
	s.SlotStatus = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultList) Validate() error {
	if s.FileList != nil {
		for _, item := range s.FileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryLocalFileUploadJobResponseBodyDataResultListFileList struct {
	// The end time of the video recording file.
	//
	// example:
	//
	// 1634746054
	FileEndTime *int32 `json:"FileEndTime,omitempty" xml:"FileEndTime,omitempty"`
	// The name of the video recording file.
	//
	// example:
	//
	// video1
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The start time of the video recording file.
	//
	// example:
	//
	// 1634747054
	FileStartTime *int32 `json:"FileStartTime,omitempty" xml:"FileStartTime,omitempty"`
}

func (s QueryLocalFileUploadJobResponseBodyDataResultListFileList) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalFileUploadJobResponseBodyDataResultListFileList) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) GetFileEndTime() *int32 {
	return s.FileEndTime
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) GetFileName() *string {
	return s.FileName
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) GetFileStartTime() *int32 {
	return s.FileStartTime
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) SetFileEndTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	s.FileEndTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) SetFileName(v string) *QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	s.FileName = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) SetFileStartTime(v int32) *QueryLocalFileUploadJobResponseBodyDataResultListFileList {
	s.FileStartTime = &v
	return s
}

func (s *QueryLocalFileUploadJobResponseBodyDataResultListFileList) Validate() error {
	return dara.Validate(s)
}
