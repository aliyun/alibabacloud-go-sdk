// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnnotationMissionSessionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveAnnotationMissionSessionListResponseBody
	GetCode() *string
	SetData(v *SaveAnnotationMissionSessionListResponseBodyData) *SaveAnnotationMissionSessionListResponseBody
	GetData() *SaveAnnotationMissionSessionListResponseBodyData
	SetHttpStatusCode(v int32) *SaveAnnotationMissionSessionListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveAnnotationMissionSessionListResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveAnnotationMissionSessionListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveAnnotationMissionSessionListResponseBody
	GetSuccess() *bool
}

type SaveAnnotationMissionSessionListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SaveAnnotationMissionSessionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request is successful, success is returned. If the request fails, an error code is returned.
	//
	// example:
	//
	// 执行失败：查找元素或操作超时,outTaskId:trademark_regnew_public_cn-zz4310rp1lb,flowId:null,nodeUUID:541b6c71-8cce-4ab0-af84-713ca9741821
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 76E93048-F90F-57B7-BD46-6097611A706D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveAnnotationMissionSessionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveAnnotationMissionSessionListResponseBody) GetData() *SaveAnnotationMissionSessionListResponseBodyData {
	return s.Data
}

func (s *SaveAnnotationMissionSessionListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveAnnotationMissionSessionListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveAnnotationMissionSessionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAnnotationMissionSessionListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveAnnotationMissionSessionListResponseBody) SetCode(v string) *SaveAnnotationMissionSessionListResponseBody {
	s.Code = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBody) SetData(v *SaveAnnotationMissionSessionListResponseBodyData) *SaveAnnotationMissionSessionListResponseBody {
	s.Data = v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBody) SetHttpStatusCode(v int32) *SaveAnnotationMissionSessionListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBody) SetMessage(v string) *SaveAnnotationMissionSessionListResponseBody {
	s.Message = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBody) SetRequestId(v string) *SaveAnnotationMissionSessionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBody) SetSuccess(v bool) *SaveAnnotationMissionSessionListResponseBody {
	s.Success = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveAnnotationMissionSessionListResponseBodyData struct {
	// The message returned for the operation.
	//
	// example:
	//
	// 执行失败：查找元素或操作超时,outTaskId:trademark_regnew_public_cn-zz4310rp1lb,flowId:null,nodeUUID:541b6c71-8cce-4ab0-af84-713ca9741821
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of messages.
	MessageList []*string `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	// The request result.
	SaveAnnotationMissionSessionListRequest *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest `json:"SaveAnnotationMissionSessionListRequest,omitempty" xml:"SaveAnnotationMissionSessionListRequest,omitempty" type:"Struct"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveAnnotationMissionSessionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) GetMessageList() []*string {
	return s.MessageList
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) GetSaveAnnotationMissionSessionListRequest() *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest {
	return s.SaveAnnotationMissionSessionListRequest
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) SetMessage(v string) *SaveAnnotationMissionSessionListResponseBodyData {
	s.Message = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) SetMessageList(v []*string) *SaveAnnotationMissionSessionListResponseBodyData {
	s.MessageList = v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) SetSaveAnnotationMissionSessionListRequest(v *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest) *SaveAnnotationMissionSessionListResponseBodyData {
	s.SaveAnnotationMissionSessionListRequest = v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) SetSuccess(v bool) *SaveAnnotationMissionSessionListResponseBodyData {
	s.Success = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBodyData) Validate() error {
	if s.SaveAnnotationMissionSessionListRequest != nil {
		if err := s.SaveAnnotationMissionSessionListRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest struct {
	// The content of the script for the approval requests
	//
	// example:
	//
	// []
	AnnotationMissionSessionListJsonString *string `json:"AnnotationMissionSessionListJsonString,omitempty" xml:"AnnotationMissionSessionListJsonString,omitempty"`
}

func (s SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest) GetAnnotationMissionSessionListJsonString() *string {
	return s.AnnotationMissionSessionListJsonString
}

func (s *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest) SetAnnotationMissionSessionListJsonString(v string) *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest {
	s.AnnotationMissionSessionListJsonString = &v
	return s
}

func (s *SaveAnnotationMissionSessionListResponseBodyDataSaveAnnotationMissionSessionListRequest) Validate() error {
	return dara.Validate(s)
}
