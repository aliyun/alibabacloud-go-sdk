// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateRumAppResponseBody
	GetCode() *int32
	SetData(v *CreateRumAppResponseBodyData) *CreateRumAppResponseBody
	GetData() *CreateRumAppResponseBodyData
	SetHttpStatusCode(v int32) *CreateRumAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateRumAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRumAppResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateRumAppResponseBody
	GetResourceGroupId() *string
	SetSuccess(v bool) *CreateRumAppResponseBody
	GetSuccess() *bool
}

type CreateRumAppResponseBody struct {
	// The HTTP status code. 2XX indicates that the request was successful. 3XX indicates that the request was redirected. 4XX indicates that a request error occurred. 5XX indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application ID and domain names. This parameter is returned if the application is created. Multiple domain names are separated with commas (,).
	//
	// example:
	//
	// ggxxxnjuz@xxxx,xxxxxx-default-cn.rum.aliyuncs.com
	Data *CreateRumAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A474FF8-7861-4D00-81B5-5BC3DA4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRumAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRumAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRumAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateRumAppResponseBody) GetData() *CreateRumAppResponseBodyData {
	return s.Data
}

func (s *CreateRumAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRumAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRumAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRumAppResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRumAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRumAppResponseBody) SetCode(v int32) *CreateRumAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRumAppResponseBody) SetData(v *CreateRumAppResponseBodyData) *CreateRumAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateRumAppResponseBody) SetHttpStatusCode(v int32) *CreateRumAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRumAppResponseBody) SetMessage(v string) *CreateRumAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRumAppResponseBody) SetRequestId(v string) *CreateRumAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRumAppResponseBody) SetResourceGroupId(v string) *CreateRumAppResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRumAppResponseBody) SetSuccess(v bool) *CreateRumAppResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRumAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateRumAppResponseBodyData struct {
	// The domain name of the SDK.
	//
	// example:
	//
	// bxxxxxxx-sdk.rum.aliyuncs.com/v2/browser-sdk.js
	CdnDomain *string `json:"CdnDomain,omitempty" xml:"CdnDomain,omitempty"`
	// The endpoint that is used to report application data.
	//
	// example:
	//
	// xxxxxxxx-default-cn.rum.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The process ID (PID) of the application.
	//
	// example:
	//
	// avccccxxxx@24cxxxxbf384dc6
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s CreateRumAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRumAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRumAppResponseBodyData) GetCdnDomain() *string {
	return s.CdnDomain
}

func (s *CreateRumAppResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateRumAppResponseBodyData) GetPid() *string {
	return s.Pid
}

func (s *CreateRumAppResponseBodyData) SetCdnDomain(v string) *CreateRumAppResponseBodyData {
	s.CdnDomain = &v
	return s
}

func (s *CreateRumAppResponseBodyData) SetEndpoint(v string) *CreateRumAppResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *CreateRumAppResponseBodyData) SetPid(v string) *CreateRumAppResponseBodyData {
	s.Pid = &v
	return s
}

func (s *CreateRumAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
