// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMseSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetMseSourceResponseBody
	GetCode() *int32
	SetData(v []*GetMseSourceResponseBodyData) *GetMseSourceResponseBody
	GetData() []*GetMseSourceResponseBodyData
	SetHttpStatusCode(v int32) *GetMseSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMseSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMseSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMseSourceResponseBody
	GetSuccess() *bool
}

type GetMseSourceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	Data []*GetMseSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned, such as the "TaskId not found" message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5EB2D865-B43F-5526-8F92-857718CF73A2
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

func (s GetMseSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMseSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMseSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetMseSourceResponseBody) GetData() []*GetMseSourceResponseBodyData {
	return s.Data
}

func (s *GetMseSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMseSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMseSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMseSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMseSourceResponseBody) SetCode(v int32) *GetMseSourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetMseSourceResponseBody) SetData(v []*GetMseSourceResponseBodyData) *GetMseSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetMseSourceResponseBody) SetHttpStatusCode(v int32) *GetMseSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMseSourceResponseBody) SetMessage(v string) *GetMseSourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetMseSourceResponseBody) SetRequestId(v string) *GetMseSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMseSourceResponseBody) SetSuccess(v bool) *GetMseSourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetMseSourceResponseBody) Validate() error {
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

type GetMseSourceResponseBodyData struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// mse-af1****-nacos-ans.mse.aliyuncs.com:8848
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of cluster.
	//
	// example:
	//
	// mse_
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// mse_prepaid_public_cn-7pp2eec****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type.
	//
	// example:
	//
	// Nacos
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMseSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMseSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMseSourceResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *GetMseSourceResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMseSourceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMseSourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMseSourceResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMseSourceResponseBodyData) SetAddress(v string) *GetMseSourceResponseBodyData {
	s.Address = &v
	return s
}

func (s *GetMseSourceResponseBodyData) SetClusterId(v string) *GetMseSourceResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetMseSourceResponseBodyData) SetInstanceId(v string) *GetMseSourceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMseSourceResponseBodyData) SetName(v string) *GetMseSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMseSourceResponseBodyData) SetType(v string) *GetMseSourceResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMseSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
