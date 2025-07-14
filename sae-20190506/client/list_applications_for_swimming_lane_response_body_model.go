// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApplicationsForSwimmingLaneResponseBody
	GetCode() *string
	SetData(v []*ListApplicationsForSwimmingLaneResponseBodyData) *ListApplicationsForSwimmingLaneResponseBody
	GetData() []*ListApplicationsForSwimmingLaneResponseBodyData
	SetErrorCode(v string) *ListApplicationsForSwimmingLaneResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListApplicationsForSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationsForSwimmingLaneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApplicationsForSwimmingLaneResponseBody
	GetSuccess() *bool
}

type ListApplicationsForSwimmingLaneResponseBody struct {
	// example:
	//
	// 200
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListApplicationsForSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0a98a02315955564772843261e****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApplicationsForSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForSwimmingLaneResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApplicationsForSwimmingLaneResponseBody) GetData() []*ListApplicationsForSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *ListApplicationsForSwimmingLaneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListApplicationsForSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationsForSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForSwimmingLaneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApplicationsForSwimmingLaneResponseBody) SetCode(v string) *ListApplicationsForSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBody) SetData(v []*ListApplicationsForSwimmingLaneResponseBodyData) *ListApplicationsForSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBody) SetErrorCode(v string) *ListApplicationsForSwimmingLaneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBody) SetMessage(v string) *ListApplicationsForSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBody) SetRequestId(v string) *ListApplicationsForSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBody) SetSuccess(v bool) *ListApplicationsForSwimmingLaneResponseBody {
	s.Success = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsForSwimmingLaneResponseBodyData struct {
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 8c573618-8d72-4407-baf4-f7b64b******
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// example:
	//
	// demo
	BaseAppName *string `json:"BaseAppName,omitempty" xml:"BaseAppName,omitempty"`
	// example:
	//
	// mse-cn-hvm47******
	MseAppId *string `json:"MseAppId,omitempty" xml:"MseAppId,omitempty"`
	// example:
	//
	// test
	MseAppName *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	// example:
	//
	// sae-test
	MseNamespaceId *string            `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
	ServiceTags    map[string]*string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
}

func (s ListApplicationsForSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetBaseAppName() *string {
	return s.BaseAppName
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetMseAppId() *string {
	return s.MseAppId
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetMseAppName() *string {
	return s.MseAppName
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) GetServiceTags() map[string]*string {
	return s.ServiceTags
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetAppId(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetAppName(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetBaseAppId(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.BaseAppId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetBaseAppName(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.BaseAppName = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetMseAppId(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.MseAppId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetMseAppName(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.MseAppName = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetMseNamespaceId(v string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.MseNamespaceId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) SetServiceTags(v map[string]*string) *ListApplicationsForSwimmingLaneResponseBodyData {
	s.ServiceTags = v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
