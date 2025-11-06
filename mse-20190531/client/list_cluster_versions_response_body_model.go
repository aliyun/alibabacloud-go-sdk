// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListClusterVersionsResponseBody
	GetCode() *int32
	SetData(v []*ListClusterVersionsResponseBodyData) *ListClusterVersionsResponseBody
	GetData() []*ListClusterVersionsResponseBodyData
	SetDynamicMessage(v string) *ListClusterVersionsResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListClusterVersionsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListClusterVersionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListClusterVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClusterVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClusterVersionsResponseBody
	GetSuccess() *bool
}

type ListClusterVersionsResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data []*ListClusterVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message.
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
	// 7717BE5B-C958-5F87-BF49-464AA276EDD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClusterVersionsResponseBody) GetData() []*ListClusterVersionsResponseBodyData {
	return s.Data
}

func (s *ListClusterVersionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListClusterVersionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClusterVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListClusterVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClusterVersionsResponseBody) SetCode(v int32) *ListClusterVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetData(v []*ListClusterVersionsResponseBodyData) *ListClusterVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterVersionsResponseBody) SetDynamicMessage(v string) *ListClusterVersionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetErrorCode(v string) *ListClusterVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetHttpStatusCode(v int32) *ListClusterVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetMessage(v string) *ListClusterVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetRequestId(v string) *ListClusterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetSuccess(v bool) *ListClusterVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListClusterVersionsResponseBody) Validate() error {
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

type ListClusterVersionsResponseBodyData struct {
	// The type of the instance.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The code of the instance type.
	//
	// example:
	//
	// NACOS_2_0_0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The version of the instance.
	//
	// example:
	//
	// 2.1.0
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListClusterVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterVersionsResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ListClusterVersionsResponseBodyData) GetShowName() *string {
	return s.ShowName
}

func (s *ListClusterVersionsResponseBodyData) SetClusterType(v string) *ListClusterVersionsResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListClusterVersionsResponseBodyData) SetCode(v string) *ListClusterVersionsResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListClusterVersionsResponseBodyData) SetShowName(v string) *ListClusterVersionsResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *ListClusterVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
