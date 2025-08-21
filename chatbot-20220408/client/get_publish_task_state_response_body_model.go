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
