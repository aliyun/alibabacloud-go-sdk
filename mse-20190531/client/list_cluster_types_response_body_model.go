// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListClusterTypesResponseBody
	GetCode() *int32
	SetData(v []*ListClusterTypesResponseBodyData) *ListClusterTypesResponseBody
	GetData() []*ListClusterTypesResponseBodyData
	SetDynamicMessage(v string) *ListClusterTypesResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListClusterTypesResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListClusterTypesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListClusterTypesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClusterTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClusterTypesResponseBody
	GetSuccess() *bool
}

type ListClusterTypesResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*ListClusterTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
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
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 821B5B05-8919-5FBB-BA75-417BFC093EB8
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

func (s ListClusterTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTypesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClusterTypesResponseBody) GetData() []*ListClusterTypesResponseBodyData {
	return s.Data
}

func (s *ListClusterTypesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListClusterTypesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClusterTypesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListClusterTypesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClusterTypesResponseBody) SetCode(v int32) *ListClusterTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetData(v []*ListClusterTypesResponseBodyData) *ListClusterTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterTypesResponseBody) SetDynamicMessage(v string) *ListClusterTypesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetErrorCode(v string) *ListClusterTypesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetHttpStatusCode(v int32) *ListClusterTypesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetMessage(v string) *ListClusterTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetRequestId(v string) *ListClusterTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetSuccess(v bool) *ListClusterTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListClusterTypesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterTypesResponseBodyData struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The type of the MSE engine that can be activated.
	//
	// example:
	//
	// Zookeeper
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListClusterTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterTypesResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ListClusterTypesResponseBodyData) GetShowName() *string {
	return s.ShowName
}

func (s *ListClusterTypesResponseBodyData) SetCode(v string) *ListClusterTypesResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListClusterTypesResponseBodyData) SetShowName(v string) *ListClusterTypesResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *ListClusterTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
