// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSystemConfigsResponseBody
	GetCode() *string
	SetData(v []*ListSystemConfigsResponseBodyData) *ListSystemConfigsResponseBody
	GetData() []*ListSystemConfigsResponseBodyData
	SetHttpStatusCode(v int32) *ListSystemConfigsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSystemConfigsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListSystemConfigsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListSystemConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSystemConfigsResponseBody
	GetSuccess() *bool
}

type ListSystemConfigsResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*ListSystemConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// Instance does not exist. Instance=outb001
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSystemConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSystemConfigsResponseBody) GetData() []*ListSystemConfigsResponseBodyData {
	return s.Data
}

func (s *ListSystemConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSystemConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSystemConfigsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListSystemConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSystemConfigsResponseBody) SetCode(v string) *ListSystemConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSystemConfigsResponseBody) SetData(v []*ListSystemConfigsResponseBodyData) *ListSystemConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListSystemConfigsResponseBody) SetHttpStatusCode(v int32) *ListSystemConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSystemConfigsResponseBody) SetMessage(v string) *ListSystemConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSystemConfigsResponseBody) SetParams(v []*string) *ListSystemConfigsResponseBody {
	s.Params = v
	return s
}

func (s *ListSystemConfigsResponseBody) SetRequestId(v string) *ListSystemConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemConfigsResponseBody) SetSuccess(v bool) *ListSystemConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSystemConfigsResponseBody) Validate() error {
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

type ListSystemConfigsResponseBodyData struct {
	// The creation time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1786085104904
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The system configuration name.\\
	//
	// callableTime: the outbound job window.\\
	//
	// calleeDailyAttemptLimit: the maximum number of daily calls to a single callee number.
	//
	// example:
	//
	// callableTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration type ID.\\
	//
	// If ObjectType is set to INSTANCE, this parameter specifies the instance ID.\\
	//
	// If ObjectType is set to TENANT, this parameter specifies the tenant ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The configuration type.\\
	//
	// INSTANCE: instance-level.\\
	//
	// TENANT: tenant-level.
	//
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The update time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1786085104904
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The system configuration content.
	//
	// example:
	//
	// 5
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSystemConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSystemConfigsResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListSystemConfigsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSystemConfigsResponseBodyData) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListSystemConfigsResponseBodyData) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListSystemConfigsResponseBodyData) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListSystemConfigsResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *ListSystemConfigsResponseBodyData) SetCreatedTime(v string) *ListSystemConfigsResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListSystemConfigsResponseBodyData) SetName(v string) *ListSystemConfigsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSystemConfigsResponseBodyData) SetObjectId(v string) *ListSystemConfigsResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *ListSystemConfigsResponseBodyData) SetObjectType(v string) *ListSystemConfigsResponseBodyData {
	s.ObjectType = &v
	return s
}

func (s *ListSystemConfigsResponseBodyData) SetUpdatedTime(v string) *ListSystemConfigsResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *ListSystemConfigsResponseBodyData) SetValue(v string) *ListSystemConfigsResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListSystemConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
