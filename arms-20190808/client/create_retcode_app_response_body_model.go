// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRetcodeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateRetcodeAppResponseBody
	GetCode() *int32
	SetData(v string) *CreateRetcodeAppResponseBody
	GetData() *string
	SetMessage(v string) *CreateRetcodeAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRetcodeAppResponseBody
	GetRequestId() *string
	SetRetcodeAppDataBean(v *CreateRetcodeAppResponseBodyRetcodeAppDataBean) *CreateRetcodeAppResponseBody
	GetRetcodeAppDataBean() *CreateRetcodeAppResponseBodyRetcodeAppDataBean
	SetSuccess(v bool) *CreateRetcodeAppResponseBody
	GetSuccess() *bool
}

type CreateRetcodeAppResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// bdidt76ehx@d5cf1cd3f7df411
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C647A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the Browser Monitoring task.
	RetcodeAppDataBean *CreateRetcodeAppResponseBodyRetcodeAppDataBean `json:"RetcodeAppDataBean,omitempty" xml:"RetcodeAppDataBean,omitempty" type:"Struct"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRetcodeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateRetcodeAppResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateRetcodeAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRetcodeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRetcodeAppResponseBody) GetRetcodeAppDataBean() *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	return s.RetcodeAppDataBean
}

func (s *CreateRetcodeAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRetcodeAppResponseBody) SetCode(v int32) *CreateRetcodeAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetData(v string) *CreateRetcodeAppResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetMessage(v string) *CreateRetcodeAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetRequestId(v string) *CreateRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetRetcodeAppDataBean(v *CreateRetcodeAppResponseBodyRetcodeAppDataBean) *CreateRetcodeAppResponseBody {
	s.RetcodeAppDataBean = v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetSuccess(v bool) *CreateRetcodeAppResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateRetcodeAppResponseBodyRetcodeAppDataBean struct {
	// The application ID.
	//
	// example:
	//
	// 135143
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The process identifier (PID) of the application.
	//
	// example:
	//
	// aokcdqn3ly@a195c6d6421****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the task.
	Tags *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) GetAppId() *int64 {
	return s.AppId
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) GetPid() *string {
	return s.Pid
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) GetTags() *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags {
	return s.Tags
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetAppId(v int64) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.AppId = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetPid(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.Pid = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetResourceGroupId(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetTags(v *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.Tags = v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) Validate() error {
	return dara.Validate(s)
}

type CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags struct {
	Tags []*CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags) GetTags() []*CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags {
	return s.Tags
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags) SetTags(v []*CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags {
	s.Tags = v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTags) Validate() error {
	return dara.Validate(s)
}

type CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) GetKey() *string {
	return s.Key
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) GetValue() *string {
	return s.Value
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) SetKey(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags {
	s.Key = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) SetValue(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags {
	s.Value = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBeanTagsTags) Validate() error {
	return dara.Validate(s)
}
