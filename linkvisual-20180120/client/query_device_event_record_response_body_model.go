// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryDeviceEventRecordResponseBody
	GetCode() *int32
	SetData(v []*QueryDeviceEventRecordResponseBodyData) *QueryDeviceEventRecordResponseBody
	GetData() []*QueryDeviceEventRecordResponseBodyData
	SetErrorMessage(v string) *QueryDeviceEventRecordResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryDeviceEventRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDeviceEventRecordResponseBody
	GetSuccess() *bool
}

type QueryDeviceEventRecordResponseBody struct {
	// The response code:
	//
	// - **200**: The call was successful.
	//
	// - Other: An error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query results.
	Data []*QueryDeviceEventRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
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

func (s QueryDeviceEventRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryDeviceEventRecordResponseBody) GetData() []*QueryDeviceEventRecordResponseBodyData {
	return s.Data
}

func (s *QueryDeviceEventRecordResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryDeviceEventRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceEventRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDeviceEventRecordResponseBody) SetCode(v int32) *QueryDeviceEventRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetData(v []*QueryDeviceEventRecordResponseBodyData) *QueryDeviceEventRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetErrorMessage(v string) *QueryDeviceEventRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetRequestId(v string) *QueryDeviceEventRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) SetSuccess(v bool) *QueryDeviceEventRecordResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDeviceEventRecordResponseBodyData struct {
	// The start time of the recording.
	//
	// example:
	//
	// 2021-09-08 15:55:35
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time of the recording.
	//
	// example:
	//
	// 2021-09-08 15:56:32
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the recording file.
	//
	// example:
	//
	// file1
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The URL of the recording file.
	//
	// > The URL contains authentication information and is valid only for a single use. It cannot be reused.
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceEventRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordResponseBodyData) GetBeginTime() *string {
	return s.BeginTime
}

func (s *QueryDeviceEventRecordResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryDeviceEventRecordResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *QueryDeviceEventRecordResponseBodyData) GetVodUrl() *string {
	return s.VodUrl
}

func (s *QueryDeviceEventRecordResponseBodyData) SetBeginTime(v string) *QueryDeviceEventRecordResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetEndTime(v string) *QueryDeviceEventRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetFileName(v string) *QueryDeviceEventRecordResponseBodyData {
	s.FileName = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) SetVodUrl(v string) *QueryDeviceEventRecordResponseBodyData {
	s.VodUrl = &v
	return s
}

func (s *QueryDeviceEventRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
