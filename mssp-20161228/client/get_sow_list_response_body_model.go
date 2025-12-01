// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSowListResponseBody
	GetCode() *string
	SetData(v []*GetSowListResponseBodyData) *GetSowListResponseBody
	GetData() []*GetSowListResponseBodyData
	SetHttpStatusCode(v int32) *GetSowListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSowListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSowListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSowListResponseBody
	GetSuccess() *bool
}

type GetSowListResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetSowListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt information for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// FA8883BC-CB18-5E28-A113-8249917CA05E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSowListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSowListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSowListResponseBody) GetData() []*GetSowListResponseBodyData {
	return s.Data
}

func (s *GetSowListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSowListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSowListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSowListResponseBody) SetCode(v string) *GetSowListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSowListResponseBody) SetData(v []*GetSowListResponseBodyData) *GetSowListResponseBody {
	s.Data = v
	return s
}

func (s *GetSowListResponseBody) SetHttpStatusCode(v int32) *GetSowListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSowListResponseBody) SetMessage(v string) *GetSowListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSowListResponseBody) SetRequestId(v string) *GetSowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSowListResponseBody) SetSuccess(v bool) *GetSowListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSowListResponseBody) Validate() error {
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

type GetSowListResponseBodyData struct {
	// Completion time.
	//
	// example:
	//
	// 2024-03-28 16:19:35
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// Operation remarks.
	//
	// example:
	//
	// 新建
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// Progress.
	//
	// example:
	//
	// IN_PREPARATION
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Record count.
	//
	// example:
	//
	// 173
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2023-03-24 16:51:26
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task type.
	//
	// example:
	//
	// 安全风险评估
	TaskTypeName *string `json:"TaskTypeName,omitempty" xml:"TaskTypeName,omitempty"`
	// Work order name.
	//
	// example:
	//
	// 安全产品配置问题与超量提醒
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
}

func (s GetSowListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSowListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSowListResponseBodyData) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetSowListResponseBodyData) GetOperateRemark() *string {
	return s.OperateRemark
}

func (s *GetSowListResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *GetSowListResponseBodyData) GetRecordCount() *int32 {
	return s.RecordCount
}

func (s *GetSowListResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSowListResponseBodyData) GetTaskTypeName() *string {
	return s.TaskTypeName
}

func (s *GetSowListResponseBodyData) GetWorkOrderName() *string {
	return s.WorkOrderName
}

func (s *GetSowListResponseBodyData) SetCompleteTime(v string) *GetSowListResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *GetSowListResponseBodyData) SetOperateRemark(v string) *GetSowListResponseBodyData {
	s.OperateRemark = &v
	return s
}

func (s *GetSowListResponseBodyData) SetProgress(v string) *GetSowListResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetSowListResponseBodyData) SetRecordCount(v int32) *GetSowListResponseBodyData {
	s.RecordCount = &v
	return s
}

func (s *GetSowListResponseBodyData) SetStartTime(v string) *GetSowListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetSowListResponseBodyData) SetTaskTypeName(v string) *GetSowListResponseBodyData {
	s.TaskTypeName = &v
	return s
}

func (s *GetSowListResponseBodyData) SetWorkOrderName(v string) *GetSowListResponseBodyData {
	s.WorkOrderName = &v
	return s
}

func (s *GetSowListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
