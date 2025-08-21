// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueInstancePublishTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTypeList(v []*string) *ContinueInstancePublishTaskResponseBody
	GetBizTypeList() []*string
	SetCreateTime(v string) *ContinueInstancePublishTaskResponseBody
	GetCreateTime() *string
	SetError(v string) *ContinueInstancePublishTaskResponseBody
	GetError() *string
	SetErrors(v map[string]interface{}) *ContinueInstancePublishTaskResponseBody
	GetErrors() map[string]interface{}
	SetId(v int64) *ContinueInstancePublishTaskResponseBody
	GetId() *int64
	SetModifyTime(v string) *ContinueInstancePublishTaskResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *ContinueInstancePublishTaskResponseBody
	GetRequestId() *string
	SetResponse(v string) *ContinueInstancePublishTaskResponseBody
	GetResponse() *string
	SetStatus(v string) *ContinueInstancePublishTaskResponseBody
	GetStatus() *string
	SetWarnings(v map[string]interface{}) *ContinueInstancePublishTaskResponseBody
	GetWarnings() map[string]interface{}
}

type ContinueInstancePublishTaskResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// {}
	Errors map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
	// example:
	//
	// 8522
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-04-12T06:30:33Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 5CBF0581-EAE7-1DC4-95C6-A089656A1E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8522
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {         "category_bind_faq": [             "以下类目没有发布到正式环境: 项目交付信息汇总"         ]     }
	Warnings map[string]interface{} `json:"Warnings,omitempty" xml:"Warnings,omitempty"`
}

func (s ContinueInstancePublishTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinueInstancePublishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueInstancePublishTaskResponseBody) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *ContinueInstancePublishTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ContinueInstancePublishTaskResponseBody) GetError() *string {
	return s.Error
}

func (s *ContinueInstancePublishTaskResponseBody) GetErrors() map[string]interface{} {
	return s.Errors
}

func (s *ContinueInstancePublishTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *ContinueInstancePublishTaskResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ContinueInstancePublishTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinueInstancePublishTaskResponseBody) GetResponse() *string {
	return s.Response
}

func (s *ContinueInstancePublishTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ContinueInstancePublishTaskResponseBody) GetWarnings() map[string]interface{} {
	return s.Warnings
}

func (s *ContinueInstancePublishTaskResponseBody) SetBizTypeList(v []*string) *ContinueInstancePublishTaskResponseBody {
	s.BizTypeList = v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetCreateTime(v string) *ContinueInstancePublishTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetError(v string) *ContinueInstancePublishTaskResponseBody {
	s.Error = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetErrors(v map[string]interface{}) *ContinueInstancePublishTaskResponseBody {
	s.Errors = v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetId(v int64) *ContinueInstancePublishTaskResponseBody {
	s.Id = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetModifyTime(v string) *ContinueInstancePublishTaskResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetRequestId(v string) *ContinueInstancePublishTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetResponse(v string) *ContinueInstancePublishTaskResponseBody {
	s.Response = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetStatus(v string) *ContinueInstancePublishTaskResponseBody {
	s.Status = &v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) SetWarnings(v map[string]interface{}) *ContinueInstancePublishTaskResponseBody {
	s.Warnings = v
	return s
}

func (s *ContinueInstancePublishTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
