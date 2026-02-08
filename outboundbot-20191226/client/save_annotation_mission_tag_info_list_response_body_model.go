// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnnotationMissionTagInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveAnnotationMissionTagInfoListResponseBody
	GetCode() *string
	SetData(v *SaveAnnotationMissionTagInfoListResponseBodyData) *SaveAnnotationMissionTagInfoListResponseBody
	GetData() *SaveAnnotationMissionTagInfoListResponseBodyData
	SetHttpStatusCode(v int32) *SaveAnnotationMissionTagInfoListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveAnnotationMissionTagInfoListResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveAnnotationMissionTagInfoListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveAnnotationMissionTagInfoListResponseBody
	GetSuccess() *bool
}

type SaveAnnotationMissionTagInfoListResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	Data *SaveAnnotationMissionTagInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// 执行失败：查找元素或操作超时，outTaskId：trademark_regnew_public_cn-uax33hol2uu,flowId：null，nodeUUID：541b6c71-8cce-4ab0-af84-713ca9741821
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 1B356EDC-F69A-53B0-B4AF-2AC42200684E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. true: The call succeeded. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveAnnotationMissionTagInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionTagInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) GetData() *SaveAnnotationMissionTagInfoListResponseBodyData {
	return s.Data
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) SetCode(v string) *SaveAnnotationMissionTagInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) SetData(v *SaveAnnotationMissionTagInfoListResponseBodyData) *SaveAnnotationMissionTagInfoListResponseBody {
	s.Data = v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) SetHttpStatusCode(v int32) *SaveAnnotationMissionTagInfoListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) SetMessage(v string) *SaveAnnotationMissionTagInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) SetRequestId(v string) *SaveAnnotationMissionTagInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) SetSuccess(v bool) *SaveAnnotationMissionTagInfoListResponseBody {
	s.Success = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveAnnotationMissionTagInfoListResponseBodyData struct {
	// Number of successful executions
	//
	// example:
	//
	// 1
	ExecCount *int32 `json:"ExecCount,omitempty" xml:"ExecCount,omitempty"`
	// API message
	//
	// example:
	//
	// 执行失败：查找元素或操作超时，outTaskId：trademark_regnew_public_cn-uax33hol2uu，flowId：null，nodeUUID：541b6c71-8cce-4ab0-af84-713ca9741821
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveAnnotationMissionTagInfoListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionTagInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) GetExecCount() *int32 {
	return s.ExecCount
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) SetExecCount(v int32) *SaveAnnotationMissionTagInfoListResponseBodyData {
	s.ExecCount = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) SetMessage(v string) *SaveAnnotationMissionTagInfoListResponseBodyData {
	s.Message = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) SetSuccess(v bool) *SaveAnnotationMissionTagInfoListResponseBodyData {
	s.Success = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
