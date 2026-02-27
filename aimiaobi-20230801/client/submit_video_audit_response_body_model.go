// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitVideoAuditResponseBody
	GetCode() *string
	SetData(v *SubmitVideoAuditResponseBodyData) *SubmitVideoAuditResponseBody
	GetData() *SubmitVideoAuditResponseBodyData
	SetHttpStatusCode(v int32) *SubmitVideoAuditResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitVideoAuditResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitVideoAuditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitVideoAuditResponseBody
	GetSuccess() *bool
}

type SubmitVideoAuditResponseBody struct {
	// 业务处理结果状态码
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 视频审校任务提交后的返回数据
	Data *SubmitVideoAuditResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP响应状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 业务处理结果描述信息
	//
	// example:
	//
	// 任务提交成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 本次API请求的唯一标识
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否处理成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitVideoAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAuditResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoAuditResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitVideoAuditResponseBody) GetData() *SubmitVideoAuditResponseBodyData {
	return s.Data
}

func (s *SubmitVideoAuditResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitVideoAuditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitVideoAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoAuditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitVideoAuditResponseBody) SetCode(v string) *SubmitVideoAuditResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitVideoAuditResponseBody) SetData(v *SubmitVideoAuditResponseBodyData) *SubmitVideoAuditResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVideoAuditResponseBody) SetHttpStatusCode(v int32) *SubmitVideoAuditResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitVideoAuditResponseBody) SetMessage(v string) *SubmitVideoAuditResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitVideoAuditResponseBody) SetRequestId(v string) *SubmitVideoAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoAuditResponseBody) SetSuccess(v bool) *SubmitVideoAuditResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitVideoAuditResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoAuditResponseBodyData struct {
	// 视频审校任务的唯一标识，可用于后续查询任务状态和结果
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitVideoAuditResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAuditResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVideoAuditResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitVideoAuditResponseBodyData) SetTaskId(v string) *SubmitVideoAuditResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitVideoAuditResponseBodyData) Validate() error {
	return dara.Validate(s)
}
