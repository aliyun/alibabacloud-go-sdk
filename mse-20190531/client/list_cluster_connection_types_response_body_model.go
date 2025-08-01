// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterConnectionTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListClusterConnectionTypesResponseBody
	GetCode() *int32
	SetData(v []*ListClusterConnectionTypesResponseBodyData) *ListClusterConnectionTypesResponseBody
	GetData() []*ListClusterConnectionTypesResponseBodyData
	SetDynamicMessage(v string) *ListClusterConnectionTypesResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListClusterConnectionTypesResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListClusterConnectionTypesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListClusterConnectionTypesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClusterConnectionTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClusterConnectionTypesResponseBody
	GetSuccess() *bool
}

type ListClusterConnectionTypesResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data []*ListClusterConnectionTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8625467C-27DD-5711-878E-6857E3074937
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterConnectionTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterConnectionTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClusterConnectionTypesResponseBody) GetData() []*ListClusterConnectionTypesResponseBodyData {
	return s.Data
}

func (s *ListClusterConnectionTypesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListClusterConnectionTypesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClusterConnectionTypesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListClusterConnectionTypesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterConnectionTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterConnectionTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClusterConnectionTypesResponseBody) SetCode(v int32) *ListClusterConnectionTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetData(v []*ListClusterConnectionTypesResponseBodyData) *ListClusterConnectionTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetDynamicMessage(v string) *ListClusterConnectionTypesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetErrorCode(v string) *ListClusterConnectionTypesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetHttpStatusCode(v int32) *ListClusterConnectionTypesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetMessage(v string) *ListClusterConnectionTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetRequestId(v string) *ListClusterConnectionTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetSuccess(v bool) *ListClusterConnectionTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterConnectionTypesResponseBodyData struct {
	// The connection type.
	//
	// example:
	//
	// slb
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListClusterConnectionTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterConnectionTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesResponseBodyData) GetShowName() *string {
	return s.ShowName
}

func (s *ListClusterConnectionTypesResponseBodyData) SetShowName(v string) *ListClusterConnectionTypesResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
