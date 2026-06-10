// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishTaskStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *GetPublishTaskStateResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *GetPublishTaskStateResponseBody
	GetCreateTime() *string
	SetError(v string) *GetPublishTaskStateResponseBody
	GetError() *string
	SetErrors(v map[string]interface{}) *GetPublishTaskStateResponseBody
	GetErrors() map[string]interface{}
	SetId(v int64) *GetPublishTaskStateResponseBody
	GetId() *int64
	SetModifyTime(v string) *GetPublishTaskStateResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *GetPublishTaskStateResponseBody
	GetRequestId() *string
	SetResponse(v string) *GetPublishTaskStateResponseBody
	GetResponse() *string
	SetStatus(v string) *GetPublishTaskStateResponseBody
	GetStatus() *string
	SetWarnings(v map[string]interface{}) *GetPublishTaskStateResponseBody
	GetWarnings() map[string]interface{}
}

type GetPublishTaskStateResponseBody struct {
	// The list of business types.
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The UTC time when the task was created.
	//
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned if the task fails.
	//
	// example:
	//
	// 检查待发布模块是否空闲发生错误，faq
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// A map of error messages for each submodule, where the key is the submodule name and the value is the error message.
	//
	// example:
	//
	// {
	//
	//     "faq": [
	//
	//         "答案资源未发布, 资源类型: 全局服务,名称: 动态答案服务"
	//
	//     ]
	//
	// }
	Errors map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// The publish task ID.
	//
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UTC time when the task was last modified.
	//
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The publish task ID. This field is redundant. Use the `Id` field instead.
	//
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The task status. Valid values:
	//
	// - `FE_RUNNING`: Publishing
	//
	// - `FE_SUCCESS`: Success
	//
	// - `FE_FAILED`: Failed
	//
	// - `FE_ABORTED`: Aborted
	//
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A map of warning messages for each submodule, where the key is the submodule name and the value is the warning message.
	//
	// example:
	//
	// {
	//
	//     "faq": [
	//
	//         "答案资源未发布,类型:service名称:null,答案资源未发布,类型:service名称:null"
	//
	//     ]
	//
	// }
	Warnings map[string]interface{} `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s GetPublishTaskStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublishTaskStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublishTaskStateResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *GetPublishTaskStateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPublishTaskStateResponseBody) GetError() *string {
	return s.Error
}

func (s *GetPublishTaskStateResponseBody) GetErrors() map[string]interface{} {
	return s.Errors
}

func (s *GetPublishTaskStateResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetPublishTaskStateResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetPublishTaskStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublishTaskStateResponseBody) GetResponse() *string {
	return s.Response
}

func (s *GetPublishTaskStateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetPublishTaskStateResponseBody) GetWarnings() map[string]interface{} {
	return s.Warnings
}

func (s *GetPublishTaskStateResponseBody) SetBizTypeList(v []*string) *GetPublishTaskStateResponseBody {
	s.BizTypeList = v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetCreateTime(v string) *GetPublishTaskStateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetError(v string) *GetPublishTaskStateResponseBody {
	s.Error = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetErrors(v map[string]interface{}) *GetPublishTaskStateResponseBody {
	s.Errors = v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetId(v int64) *GetPublishTaskStateResponseBody {
	s.Id = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetModifyTime(v string) *GetPublishTaskStateResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetRequestId(v string) *GetPublishTaskStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetResponse(v string) *GetPublishTaskStateResponseBody {
	s.Response = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetStatus(v string) *GetPublishTaskStateResponseBody {
	s.Status = &v
	return s
}

func (s *GetPublishTaskStateResponseBody) SetWarnings(v map[string]interface{}) *GetPublishTaskStateResponseBody {
	s.Warnings = v
	return s
}

func (s *GetPublishTaskStateResponseBody) Validate() error {
	return dara.Validate(s)
}
