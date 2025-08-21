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
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error      *string                `json:"Error,omitempty" xml:"Error,omitempty"`
	Errors     map[string]interface{} `json:"Errors,omitempty" xml:"Errors,omitempty"`
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
	Status   *string                `json:"Status,omitempty" xml:"Status,omitempty"`
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
