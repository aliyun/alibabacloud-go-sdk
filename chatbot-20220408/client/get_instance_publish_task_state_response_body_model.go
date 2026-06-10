// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancePublishTaskStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *GetInstancePublishTaskStateResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *GetInstancePublishTaskStateResponseBody
	GetCreateTime() *string
	SetError(v string) *GetInstancePublishTaskStateResponseBody
	GetError() *string
	SetErrors(v map[string]interface{}) *GetInstancePublishTaskStateResponseBody
	GetErrors() map[string]interface{}
	SetId(v int64) *GetInstancePublishTaskStateResponseBody
	GetId() *int64
	SetModifyTime(v string) *GetInstancePublishTaskStateResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *GetInstancePublishTaskStateResponseBody
	GetRequestId() *string
	SetResponse(v string) *GetInstancePublishTaskStateResponseBody
	GetResponse() *string
	SetStatus(v string) *GetInstancePublishTaskStateResponseBody
	GetStatus() *string
	SetWarnings(v map[string]interface{}) *GetInstancePublishTaskStateResponseBody
	GetWarnings() map[string]interface{}
}

type GetInstancePublishTaskStateResponseBody struct {
	// The list of business types.
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The UTC time when the task was created.
	//
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message if the task fails.
	//
	// example:
	//
	// 检查待发布模块是否空闲发生错误，faq
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// A map of error messages for each submodule, where the key is the submodule and the value is an array of errors.
	//
	// example:
	//
	// {
	//
	//     "robot_ds": [
	//
	//         "{\\"dialogName\\":\\"TEST\\",\\"errorMessage\\":[\\"medusa@invocation.error.service.offline@请发布对话流所引用服务后重试！\\"],\\"needRefresh\\":false,\\"success\\":false}"
	//
	//     ]
	//
	// }
	Errors map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// The task ID.
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
	// The task ID, returned as a string.
	//
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A map of warning messages for each submodule, where the key is the submodule and the value is an array of warnings.
	//
	// example:
	//
	// {
	//
	//     "category_bind_faq": [
	//
	//         "以下类目没有发布到正式环境：项目交付类目"
	//
	//     ]
	//
	// }
	Warnings map[string]interface{} `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s GetInstancePublishTaskStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePublishTaskStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancePublishTaskStateResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *GetInstancePublishTaskStateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetInstancePublishTaskStateResponseBody) GetError() *string {
	return s.Error
}

func (s *GetInstancePublishTaskStateResponseBody) GetErrors() map[string]interface{} {
	return s.Errors
}

func (s *GetInstancePublishTaskStateResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetInstancePublishTaskStateResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetInstancePublishTaskStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstancePublishTaskStateResponseBody) GetResponse() *string {
	return s.Response
}

func (s *GetInstancePublishTaskStateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstancePublishTaskStateResponseBody) GetWarnings() map[string]interface{} {
	return s.Warnings
}

func (s *GetInstancePublishTaskStateResponseBody) SetBizTypeList(v []*string) *GetInstancePublishTaskStateResponseBody {
	s.BizTypeList = v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetCreateTime(v string) *GetInstancePublishTaskStateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetError(v string) *GetInstancePublishTaskStateResponseBody {
	s.Error = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetErrors(v map[string]interface{}) *GetInstancePublishTaskStateResponseBody {
	s.Errors = v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetId(v int64) *GetInstancePublishTaskStateResponseBody {
	s.Id = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetModifyTime(v string) *GetInstancePublishTaskStateResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetRequestId(v string) *GetInstancePublishTaskStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetResponse(v string) *GetInstancePublishTaskStateResponseBody {
	s.Response = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetStatus(v string) *GetInstancePublishTaskStateResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) SetWarnings(v map[string]interface{}) *GetInstancePublishTaskStateResponseBody {
	s.Warnings = v
	return s
}

func (s *GetInstancePublishTaskStateResponseBody) Validate() error {
	return dara.Validate(s)
}
